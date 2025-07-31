package app

import (
	"context"
	"fmt"

	"github.com/gfdmit/subscription-service/config"
	v1 "github.com/gfdmit/subscription-service/internal/handlers/http/v1"
	"github.com/gfdmit/subscription-service/internal/httpserver"
	"github.com/gfdmit/subscription-service/internal/repository/postgres"
	"github.com/gfdmit/subscription-service/internal/service/subscription"
)

func Run(conf config.Config) error {
	ctx := context.Background()

	repo, err := postgres.New(conf, ctx)
	if err != nil {
		return fmt.Errorf("error when setting up repository: %v", err)
	}

	service := subscription.New(repo)

	handler := v1.New(service)

	httpserver := httpserver.New(conf.HTTPServer, handler)

	return httpserver.Run(ctx)
}
