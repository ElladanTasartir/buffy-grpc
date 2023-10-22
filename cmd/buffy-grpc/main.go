package main

import (
	episodesv1 "github.com/ElladanTasartir/buffy-grpc/gen/go/episodes/v1"
	greetingv1 "github.com/ElladanTasartir/buffy-grpc/gen/go/greeting/v1"
	service "github.com/ElladanTasartir/buffy-grpc/internal/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var (
	listener net.Listener
	server   *grpc.Server
	logger   *zap.Logger
)

func main() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		logger.Panic("Failed to create logger", zap.Error(err))
	}

	defer logger.Sync()

	initListener()

	server = grpc.NewServer()
	reflection.Register(server)

	episodesService, err := service.NewEpisodesService()
	if err != nil {
		logger.Panic("Failed to create EpisodesService", zap.Error(err))
	}

	greetingv1.RegisterGreeterServiceServer(server, service.GreetingService{})
	episodesv1.RegisterEpisodesServiceServer(server, episodesService)
	logger.Info("Handlers registered")

	go signalsListener(server)

	logger.Info("Starting gRPC server...")
	if err = server.Serve(listener); err != nil {
		logger.Panic("Failed to start gRPC server", zap.Error(err))
	}
}

func initListener() {
	var err error
	addr := "localhost:50051"

	listener, err = net.Listen("tcp", addr)
	if err != nil {
		logger.Panic("Failed to listen",
			zap.String("address", addr),
			zap.Error(err),
		)
	}

	logger.Info("Started Listening...", zap.String("address", addr))
	return
}

func signalsListener(server *grpc.Server) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	_ = <-sigs

	logger.Info("Gracefully stopping server...")
	server.GracefulStop()
}
