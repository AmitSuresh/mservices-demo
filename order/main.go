package main

import (
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/AmitSuresh/orderapi/src/app/controller"
	"github.com/AmitSuresh/orderapi/src/app/model"
	"github.com/AmitSuresh/orderapi/src/infra/config"
	queue "github.com/AmitSuresh/orderapi/src/infra/kafka"
	"github.com/AmitSuresh/orderapi/src/infra/proto/order_service"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	cfg           *config.Config
	l             *zap.Logger
	server        *http.Server
	g             *gin.Engine
	shippedOrders model.ShipmentDetails
	TopicStr      string = "Shipping_Info"
)

func main() {

	l, _ = zap.NewProduction()
	cfg = config.LoadConfig(l)
	g = gin.New()

	g.Use(gin.Recovery())
	g.Use(gin.Logger())

	cc, err := grpc.NewClient(cfg.GrpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		l.Fatal("error starting new client", zap.Error(err))
	}

	grpcClient := order_service.NewOrderServiceClient(cc)
	controller.CreateOrder(g, grpcClient)
	server = &http.Server{
		Addr:    cfg.ServerAddr,
		Handler: g,
	}

	k, err := queue.NewConsumer(cfg, l)
	if err != nil {
		l.Fatal("error", zap.Error(err))
	}
	err = k.SubscribeTopics([]string{TopicStr}, nil)
	if err != nil {
		l.Fatal("Failed to subscribe to topics", zap.Error(err))
	}

	msg_count := 0
	run := true

	go func() {
		if err := server.ListenAndServe(); err != nil {
			zap.L().Fatal("failed to start server")
		}
	}()

	go func() {

	}()
	for run {
		ev := k.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			msg_count += 1
			if msg_count%2 == 0 {
				k.Commit()
			}
			l.Info("Message on", zap.Any("", e.Value))
			switch *e.TopicPartition.Topic {
			case TopicStr:
				order := &model.OrderShipping{}
				err := json.Unmarshal(e.Value, &order)
				if err != nil {
					l.Error("error unmarshalling msg", zap.Error(err))
					continue
				}
				model.ProcessOrderShipping(order, &shippedOrders, l)
			}

		case kafka.PartitionEOF:
			l.Info("Reached", zap.Any("", e))
		case kafka.Error:
			l.Error("Error", zap.Error(e))
			run = false
		default:
			//l.Info("Default", zap.Any("", e))
		}
	}
	k.Close()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

}
