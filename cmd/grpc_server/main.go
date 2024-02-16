package main

import (
	"flag"
	"log"
	"net"

	"github.com/sarastee/auth/internal/config"
	envcfg "github.com/sarastee/auth/internal/config/env"
	desc "github.com/sarastee/auth/pkg/user_api_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "./config/.env", "path to config file")
}

type server struct {
	desc.UnimplementedUserAPIV1Server
}

func main() {
	flag.Parse()
	//ctx := context.Background()

	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	gRPCConfig, err := envcfg.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	//pgConfig, err := envcfg.NewPGConfig()
	//if err != nil {
	//	log.Fatalf("failed to get pg config: %v", err)
	//}

	lis, err := net.Listen("tcp", gRPCConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	//if err != nil {
	//	log.Panicf("failed to connect to database: %v", err)
	//}
	//defer pool.Close()

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserAPIV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
