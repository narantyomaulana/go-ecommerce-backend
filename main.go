package main

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/narantyomaulana/go-grpc-ercommerce-be/internal/grpcmiddleware"
	"github.com/narantyomaulana/go-grpc-ercommerce-be/internal/handler"
	"github.com/narantyomaulana/go-grpc-ercommerce-be/internal/repository"
	"github.com/narantyomaulana/go-grpc-ercommerce-be/internal/service"
	"github.com/narantyomaulana/go-grpc-ercommerce-be/pb/auth"
	"github.com/narantyomaulana/go-grpc-ercommerce-be/pkg/database"
	gocache "github.com/patrickmn/go-cache"
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

	db := database.ConnectDB(ctx, os.Getenv("DB_URI"))
	log.Println("Database connection established")

	cacheService := gocache.New(time.Hour*24, time.Hour)

	authMiddleware := grpcmiddleware.NewAuthMiddleware(cacheService)

	authRepository := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepository, cacheService)
	authHandler := handler.NewAuthHandler(authService)

	serv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpcmiddleware.ErrorMiddleware,
			authMiddleware.Middleware,
		),
	)

	auth.RegisterAuthServiceServer(serv, authHandler)

	if os.Getenv("ENVIRONMENT") == "dev" {
		reflection.Register(serv)
		log.Println("Reflection registered for gRPC server")
	}

	log.Println("Server is starting on port 50051...")
	if err := serv.Serve(lis); err != nil {
		log.Panicf("Server is error when serving: %v", err)
	}
}
