package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AmitSuresh/shipping/internal/queue"
	"github.com/AmitSuresh/shipping/internal/repo"
	"github.com/AmitSuresh/shipping/pkg/config"
	"github.com/AmitSuresh/shipping/pkg/server"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ginServer struct {
	GinEng *gin.Engine
	Server *http.Server
	L      *zap.Logger
}

func NewServer(l *zap.Logger, cfg *config.Config) *ginServer {
	g := server.NewGin(cfg)

	return &ginServer{
		GinEng: g,
		Server: server.NewServer(cfg, g),
		L:      l,
	}
}

func CreateShipping(gs *ginServer, db repo.RepoHandler, k *kafka.Producer, t *kafka.TopicPartition) {
	shGroup := gs.GinEng.Group("/api/orders")
	shGroup.GET("/", func(c *gin.Context) {
		GetOrderInfo(c, db, k, gs.L, t)
	})
	shGroup.GET("/all", func(c *gin.Context) {
		nOrders := c.DefaultQuery("num_orders", "500")
		pval, perr := strconv.Atoi(nOrders)
		if perr != nil {
			c.JSON(http.StatusBadRequest, perr)
			return
		}
		orders, err := db.QueryAllOrders(pval)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusAccepted, orders)
	})
}

func GetOrderInfo(c *gin.Context, db repo.RepoHandler, k *kafka.Producer, l *zap.Logger, t *kafka.TopicPartition) {
	nOrders := c.DefaultQuery("num_orders", "500")
	pval, perr := strconv.Atoi(nOrders)
	if perr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error_message": "Please enter correct number", "parse_error": perr})
		return
	}

	shippedOrders, err := db.QueryOrders(pval)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error_message": err})
		return
	}

	//var msg *kafka.Message
	for _, v := range shippedOrders {
		l.Info("", zap.Any("v is", v))
		b, err := json.Marshal(v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_message": err})
			return
		}
		msg := queue.NewMsg(b, t)
		l.Info("Message Info",
			zap.ByteString("Value", msg.Value),
			zap.Int32("Partition", msg.TopicPartition.Partition),
			zap.String("Topic", *msg.TopicPartition.Topic),
			zap.Any("Offset", msg.TopicPartition.Offset),
		)

		err = queue.ProduceMsg(k, msg, nil, l)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_message": fmt.Sprintf(err.Error())})
			return
		}
	}
	c.JSON(http.StatusAccepted, gin.H{"success_message": shippedOrders})
}
