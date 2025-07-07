package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/narantyomaulana/go-grpc-ercommerce-be/internal/handler"
	"github.com/narantyomaulana/go-grpc-ercommerce-be/pb/service"
	"github.com/narantyomaulana/go-grpc-ercommerce-be/pkg/database"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx := context.Background()

	godotenv.Load()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Panicf("Error when listening: %v", err)
	}

	database.ConnectDB(ctx, os.Getenv("DB_URI"))
	log.Println("Database connection established")

	serviceHandler := handler.NewServiceHandler()

	serv := grpc.NewServer()

	service.RegisterHelloWorldServiceServer(serv, serviceHandler)

	if os.Getenv("ENVIRONMENT") == "dev" {
		reflection.Register(serv)
		log.Println("Reflection registered for gRPC server")
	}

	log.Println("Server is starting on port 50051...")
	if err := serv.Serve(lis); err != nil {
		log.Panicf("Server is error when serving: %v", err)
	}
}
