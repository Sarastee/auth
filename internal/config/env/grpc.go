package env

import (
	"errors"
	"os"

	"github.com/sarastee/auth/internal/config"
)

const (
	grpcHostEnvName = "GRPC_HOST"
	grpcPortEnvName = "GRPC_PORT"
)

type GRPCCfgSearcher struct{}

func NewGRPCCfgSearcher() *GRPCCfgSearcher {
	return &GRPCCfgSearcher{}
}

func (s *GRPCCfgSearcher) Get() (*config.GRPCConfig, error) {
	hostGRPC := os.Getenv(grpcHostEnvName)
	if len(hostGRPC) == 0 {
		return nil, errors.New("grpc host not found")
	}

	portGRPC := os.Getenv(grpcPortEnvName)
	if len(portGRPC) == 0 {
		return nil, errors.New("grpc port not found")
	}

	return &config.GRPCConfig{
		Host: hostGRPC,
		Port: portGRPC,
	}, nil
}
