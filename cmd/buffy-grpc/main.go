package main

import (
	episodes_v1 "github.com/ElladanTasartir/buffy-grpc/gen/go/episodes/v1"
	greeting_v1 "github.com/ElladanTasartir/buffy-grpc/gen/go/greeting/v1"
	service "github.com/ElladanTasartir/buffy-grpc/internal/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
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
		panic(err)
	}

	defer logger.Sync()

	initListener()

	server = grpc.NewServer()

	greeting_v1.RegisterGreeterServiceServer(server, service.GreetingService{})
	episodes_v1.RegisterEpisodesServiceServer(server, service.EpisodesService{})
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
