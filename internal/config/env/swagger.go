package env

import (
	"errors"
	"os"

	"github.com/sarastee/auth/internal/config"
)

const (
	swaggerHostEnvName = "SWAGGER_HOST"
	swaggerPortEnvName = "SWAGGER_PORT"
)

// SwaggerConfigSearcher searcher for Swagger config.
type SwaggerConfigSearcher struct{}

// NewSwaggerConfigSearcher get instance for Swagger config searcher.
func NewSwaggerConfigSearcher() *SwaggerConfigSearcher {
	return &SwaggerConfigSearcher{}
}

// Get searcher for Swagger config.
func (s *SwaggerConfigSearcher) Get() (*config.SwaggerConfig, error) {
	host := os.Getenv(swaggerHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("swagger host not found")
	}

	port := os.Getenv(swaggerPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("swagger port not found")
	}

	return &config.SwaggerConfig{
		Host: host,
		Port: port,
	}, nil
}
