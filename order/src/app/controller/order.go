package controller

import (
	"fmt"
	"net/http"

	"github.com/AmitSuresh/orderapi/src/app/model"
	"github.com/AmitSuresh/orderapi/src/infra/proto/order_service"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CreateOrder(g *gin.Engine, client order_service.OrderServiceClient) {
	g.POST("/order", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		var o model.CreateOrder

		if err := c.BindJSON(&o); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_message": "Invalid input: " + err.Error()})
			return
		}

		req := &order_service.OrderReq{
			Customer: o.CustomerName,
			Quantity: o.Quantity,
			Sku:      o.Sku,
		}
		resp, err := client.CreateOrder(c.Request.Context(), req)
		if err != nil {
			if s, ok := status.FromError(err); ok {
				switch s.Code() {
				case codes.AlreadyExists:
					c.JSON(http.StatusBadRequest, gin.H{"error_message": s.Message()})
					return
				case codes.InvalidArgument:
					c.JSON(http.StatusBadRequest, gin.H{"error_message": s.Message()})
					return
				default:
					c.JSON(http.StatusInternalServerError, gin.H{"error_message": s.Message()})
					return
				}
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error_message": err})
			return
		}

		c.JSON(http.StatusAccepted, gin.H{"success_message": fmt.Sprintf("Customer of id %s is created", resp.Id)})
	})
}
