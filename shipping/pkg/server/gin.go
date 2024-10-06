package server

import (
	"github.com/AmitSuresh/shipping/pkg/config"
	"github.com/gin-gonic/gin"
)

func NewGin(cfg *config.Config) *gin.Engine {
	g := gin.New()

	g.Use(gin.Recovery())
	g.Use(gin.Logger())

	return g
}
