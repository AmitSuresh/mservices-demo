package server

import (
	"net/http"
	"time"

	"github.com/AmitSuresh/shipping/pkg/config"
	"github.com/gin-gonic/gin"
)

func NewServer(cfg *config.Config, g *gin.Engine) *http.Server {
	return &http.Server{
		Addr:         cfg.ServerAddr,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  150 * time.Second,
		Handler:      g,
	}
}
