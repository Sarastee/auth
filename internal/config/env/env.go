package env

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/caarlos0/env/v10"
)

type Postgres struct {
	Host               string `env:"POSTGRES_HOST" envDefault:"localhost"`
	Port               string `env:"POSTGRES_PORT" envDefault:"5432"`
	User               string `env:"POSTGRES_USER" envDefault:"root"`
	Password           string `env:"POSTGRES_PASSWORD" envDefault:"password"`
	DB                 string `env:"POSTGRES_DB" envDefault:"postgres"`
	SSLMode            string `env:"POSTGRES_SSL_MODE" envDefault:"disable"`
	DSN                string `env:"POSTGRES_DSN"`
	MaxOpenConnections int    `env:"POSTGRES_MAX_OPEN_CONNS" envDefault:"100"`
}

type GRPC struct {
	Host     string `env:"GRPC_HOST" envDefault:"localhost"`
	Port     string `env:"GRPC_PORT" envDefault:"50051"`
	Protocol string `env:"GRPC_PROTOCOL" envDefault:"tcp"`
	Address  string
}

type Config struct {
	Env      string `env:"ENV" envDefault:"local"`
	Postgres Postgres
	GRPC     GRPC
}

func New() (*Config, error) {

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed load config from enviroment: %v", err)
	}

	DSN(&cfg.Postgres)
	cfg.GRPC.Address = net.JoinHostPort(cfg.GRPC.Host, cfg.GRPC.Port)

	return cfg, nil
}

func DSN(p *Postgres) {
	p.DSN = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		p.User,
		p.Password,
		p.Host,
		p.Port,
		p.DB,
		p.SSLMode,
	)
}

func configPath() string {
	cfgPath := flag.String("config-path", "", "sets the config path")
	if cfgPath != nil {
		if *cfgPath != "" {
			return *cfgPath
		}
	}

	envCfgPath := os.Getenv("CONFIG_PATH")

	return envCfgPath
}
