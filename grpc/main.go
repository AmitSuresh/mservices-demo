package main

import (
	"net"

	"github.com/AmitSuresh/grpc/infra/config"
	"github.com/AmitSuresh/grpc/internal"
	"github.com/AmitSuresh/grpc/internal/domain"
	"github.com/AmitSuresh/grpc/proto/order_service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	listener net.Listener
	gs       *grpc.Server
	l        *zap.Logger
	cfg      *config.Config
)

func init() {
	l, _ = zap.NewProduction()
	cfg = config.LoadConfig(l)

	gs = grpc.NewServer()
	db, err := gorm.Open(postgres.Open(cfg.Dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		l.Fatal("error connecting to db", zap.Error(err))
	}

	li, err := net.Listen("tcp", cfg.GrpcAddr)
	if err != nil {
		l.Fatal("error starting listener", zap.Error(err))
	}
	listener = li

	reflection.Register(gs)
	db.AutoMigrate(&domain.Order{})
	repo := domain.NewOrderRepo(l, db)
	svr := internal.NewOrderService(repo, l)
	order_service.RegisterOrderServiceServer(gs, svr)

}

func main() {
	defer l.Sync()
	if err := gs.Serve(listener); err != nil {
		zap.L().Fatal("failed to start grpc", zap.Error(err))
	}
}
