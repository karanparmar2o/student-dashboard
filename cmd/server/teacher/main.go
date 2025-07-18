package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	teacherpb "github.com/karanparmar2o/student-dashboard/api/teacher"
	"github.com/karanparmar2o/student-dashboard/internal/teacher/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Load everything from config
	cfg := config.LoadTeacherConfig()

	// Channel for graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start gRPC server
	go func() {
		lis, err := net.Listen("tcp", cfg.GRPCPort)
		if err != nil {
			log.Fatalf("❌ failed to listen on %s: %v", cfg.GRPCPort, err)
		}
		grpcServer := grpc.NewServer()
		teacherpb.RegisterTeacherServiceServer(grpcServer, cfg.Handler)
		log.Printf("✅ gRPC server running at %s\n", cfg.GRPCPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("❌ failed to serve gRPC: %v", err)
		}
	}()

	// Start HTTP gateway
	go func() {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		}

		err := teacherpb.RegisterTeacherServiceHandlerFromEndpoint(ctx, mux, "localhost"+cfg.GRPCPort, opts)
		if err != nil {
			log.Fatalf("❌ failed to start HTTP gateway: %v", err)
		}

		log.Printf("🌐 HTTP gateway running at %s\n", cfg.HTTPPort)
		if err := http.ListenAndServe(cfg.HTTPPort, mux); err != nil {
			log.Fatalf("❌ failed to serve HTTP: %v", err)
		}
	}()

	// Wait for stop signal
	<-stop
	log.Println("🛑 Shutting down teacher service gracefully...")
}
