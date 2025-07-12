package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	studentpb "github.com/karanparmar2o/student-dashboard/api/student"
	"github.com/karanparmar2o/student-dashboard/internal/student/handler"
	repository "github.com/karanparmar2o/student-dashboard/internal/student/repo"
	"github.com/karanparmar2o/student-dashboard/internal/student/service"

	"google.golang.org/grpc"
)

func main() {
	// Initialize repo, service, handler
	repo := repository.NewStudentRepo()
	svc := service.NewStudentService(repo)
	h := handler.NewStudentHandler(svc)

	// Start listener for gRPC
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		grpcServer := grpc.NewServer()
		studentpb.RegisterStudentServiceServer(grpcServer, h)
		log.Println("gRPC server running at :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Start HTTP gateway
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := studentpb.RegisterStudentServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	if err != nil {
		log.Fatalf("failed to start HTTP gateway: %v", err)
	}

	log.Println("HTTP gateway running at :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("failed to serve HTTP: %v", err)
	}
}
