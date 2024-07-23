package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lookandhate/microservice-courese/auth/internal/config"
	"github.com/lookandhate/microservice-courese/auth/internal/grpc/auth"
	"github.com/lookandhate/microservice-courese/auth/internal/repository/user"
	service "github.com/lookandhate/microservice-courese/auth/internal/service/user"
	"github.com/lookandhate/microservice-courese/auth/pkg/auth_v1"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	cfg := config.MustLoad()
	serverHost := fmt.Sprintf("localhost:%d", cfg.GPRC.Port)
	log.Printf("Serving at %v", serverHost)

	connectionPool, err := pgxpool.New(ctx, cfg.DB.GetDSN())
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	userRepo := user.NewPostgresRepository(connectionPool)
	userService := service.NewUserService(userRepo)
	server := auth.NewAuthServer(userService)

	lis, err := net.Listen("tcp", serverHost)
	if err != nil {
		log.Fatalf("Failed to listen: %s", err)
	}

	s := grpc.NewServer()
	auth_v1.RegisterAuthServer(s, server)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve GRPC server %s", err)
	}
}
