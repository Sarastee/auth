package config

import (
	"fmt"
	"net"

	"github.com/joho/godotenv"
)

// GRPCConfigSearcher interface for search grpc config
type GRPCConfigSearcher interface {
	Get() (*GRPCConfig, error)
}

// PgConfigSearcher interface for search PG config.
type PgConfigSearcher interface {
	Get() (*PgConfig, error)
}

// HTTPConfigSearcher interface for search HTTP config.
type HTTPConfigSearcher interface {
	Get() (*HTTPConfig, error)
}

// SwaggerConfigSearcher interface for search Swagger config.
type SwaggerConfigSearcher interface {
	Get() (*SwaggerConfig, error)
}

// Load dotenv from path to env
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}
	return nil
}

// GRPCConfig grpc config.
type GRPCConfig struct {
	Host string
	Port string
}

// Address get address for grpc server.
func (cfg *GRPCConfig) Address() string {
	return net.JoinHostPort(cfg.Host, cfg.Port)
}

// PgConfig config for postgresql.
type PgConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

// DSN ..
func (cfg *PgConfig) DSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName,
	)
}

// HTTPConfig is config for HTTP
type HTTPConfig struct {
	Host string
	Port string
}

// Address get address from config
func (cfg *HTTPConfig) Address() string {
	return net.JoinHostPort(cfg.Host, cfg.Port)
}

// SwaggerConfig is config for Swagger
type SwaggerConfig struct {
	Host string
	Port string
}

// Address get address from config
func (cfg *SwaggerConfig) Address() string {
	return net.JoinHostPort(cfg.Host, cfg.Port)
}
