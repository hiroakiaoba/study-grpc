package main

import (
	"log"
	"net"
	"os"
	"os/signal"

	srv "api/gen/service"
	"api/handler"
	"api/middleware"
	"api/repository"
	"api/util"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
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

	// create repositories
	userRepository := repository.NewInMemoryUserRepository()

	// create utilities
	authUtil := util.NewJWTAuth()

	// create middlewares
	authMid := middleware.NewAuthMiddleware(userRepository, authUtil)

	authInterceptor := grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(authMid.Authenticate))
	server := grpc.NewServer(authInterceptor)

	srv.RegisterTodoServiceServer(
		server,
		handler.NewTodoHandler(),
	)
	srv.RegisterUserServiceServer(
		server,
		handler.NewUserHandler(userRepository, authUtil),
	)
	srv.RegisterProjectServiceServer(
		server,
		handler.NewProjectHandler(),
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
