package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	myqueue "github.com/AmitSuresh/shipping/internal/queue"
	"github.com/AmitSuresh/shipping/internal/repo"
	myserver "github.com/AmitSuresh/shipping/internal/server"
	"github.com/AmitSuresh/shipping/pkg/config"
	queue "github.com/AmitSuresh/shipping/pkg/kafka"
	logs "github.com/AmitSuresh/shipping/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	l := logs.NewLogger()
	defer l.Sync()

	cfg := config.NewConfig(l)
	l.Info(cfg.DbDSN)
	db := repo.NewRepo(l, cfg)
	//repo.MigrateAll(db)
	gs := myserver.NewServer(l, cfg)

	k, kerr := queue.NewProducer(cfg, l)
	if kerr != nil {
		l.Error("Error getting producer", zap.Error(kerr))
	}

	t := myqueue.NewPartition(&myqueue.MsgStr, 10)
	myserver.CreateShipping(gs, db, k, t)

	go func() {
		err := gs.Server.ListenAndServe()
		if err != nil {
			l.Info(err.Error())
			//os.Exit(1)
		}
	}()

	ctxWithTime, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	sig := <-sigChan
	switch sig {
	case os.Interrupt:
		l.Info("received shutdown signal1")
	case syscall.SIGTERM:
		l.Info("received shutdown signal2")
	default:
		l.Info("unknown signal received.")
	}

	err := gs.Server.Shutdown(ctxWithTime)
	if err != nil {
		l.Fatal("error shutting down gracefully.")
		//os.Exit(1)
	}
}
