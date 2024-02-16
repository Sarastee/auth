package config

import "github.com/joho/godotenv"

// Load loads read env files and load them into ENV
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return err
}

// GRPCConfig interface GRPCConfig
type GRPCConfig interface {
	Address() string
}

// PGConfig interface PGConfig
type PGConfig interface {
	DSN() string
}
