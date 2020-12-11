package main

import (
	"log"
	"net"
	"os"
	"os/signal"

	srv "api/gen/service"
	"api/handler"
	"api/repository"
	"api/util"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "50051"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	// create repositories
	userRepository := repository.NewInMemoryUserRepository()

	// create utilities
	auth := util.NewJWTAuth()

	srv.RegisterTodoServiceServer(
		server,
		handler.NewTodoHandler(),
	)
	srv.RegisterUserServiceServer(
		server,
		handler.NewUserHandler(userRepository, auth),
	)
	reflection.Register(server)

	go func() {
		log.Printf("start gRPC server port: %v", port)
		server.Serve(lis)
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	server.GracefulStop()
}
