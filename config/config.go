package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Postgres
	HTTPServer
	Pool
}

type Postgres struct {
	User       string        `env:"POSTGRES_USER" env-default:"postgres"`
	Pass       string        `env:"POSTGRES_PASSWORD" env-default:"postgres"`
	Host       string        `env:"POSTGRES_HOST" env-default:"postgres"`
	Port       string        `env:"POSTGRES_PORT" env-default:"5432"`
	DB         string        `env:"POSTGRES_DB" env-default:"postgres"`
	Timeout    time.Duration `env:"POSTGRES_TIMEOUT" env-default:"5s"`
	Migrations string        `env:"POSTGRES_MIGRATIONS" env-default:"./migrations"`
	SSL        string        `env:"POSTGRES_SSL" env-default:"disable"`
}

type HTTPServer struct {
	BindAddress     string        `env:"BIND_ADDRESS" env-default:"localhost"`
	BindPort        string        `env:"BIND_PORT" env-default:"8080"`
	ShutdownTimeout time.Duration `env:"SHUTDOWN_TIMEOUT" env-default:"5s"`
	ReadTimeout     time.Duration `env:"READ_TIMEOUT" env-default:"5s"`
	WriteTimeout    time.Duration `env:"WRITE_TIMEOUT" env-default:"5s"`
}

type Pool struct {
	MaxConns    int32         `env:"POOL_MAX_CONNS" env-default:"20"`
	MinConns    int32         `env:"POOL_MIN_CONNS" env-default:"5"`
	MaxLifetime time.Duration `env:"POOL_MAX_LIFETIME" env-default:"1h"`
	HealthCheck time.Duration `env:"POOL_HEALTH_CHECK" env-default:"1m"`
}

func New(env string) (*Config, error) {
	conf := &Config{}

	if err := godotenv.Overload(env); err != nil {
		return nil, fmt.Errorf("godotenv.Overload: %v", err)
	}

	if err := cleanenv.ReadEnv(conf); err != nil {
		return nil, fmt.Errorf("cleanenv.Readenv: %v", err)
	}

	return conf, nil
}
