package main

import (
	"context"

	"github.com/DevLucca/rinha/application/controller"
	applicationService "github.com/DevLucca/rinha/application/service"
	"github.com/DevLucca/rinha/domain/repository"
	"github.com/DevLucca/rinha/domain/service"
	"github.com/DevLucca/rinha/infra/config"
	"github.com/DevLucca/rinha/infra/http"
	"github.com/DevLucca/rinha/infra/persistence/cache"
	"github.com/DevLucca/rinha/infra/persistence/cache/redis"
	"github.com/DevLucca/rinha/infra/persistence/database/mysql"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			// Infra Deps
			context.Background,
			// *config.Config
			config.Read,

			// cache.Cache
			func(cfg *config.Config) cache.Cache {
				client := redis.NewClient(redis.ConfigOptions{
					Server:   cfg.Cache.Server,
					DB:       cfg.Cache.DB,
					Password: cfg.Cache.Password,
					Port:     cfg.Cache.Port,
					Prefix:   cfg.Cache.Prefix,
				})
				return client
			},

			// Persistence Database
			mysql.NewMySQLClient,

			// Repository
			fx.Annotate(
				// inmemory.NewInMemoryPersonRepository,
				mysql.NewMySQLPeopleRepository,
				fx.As(new(repository.Person)),
			),
		),

		// Domain Deps
		fx.Provide(
			service.NewPersonService,
		),

		// Application Deps
		fx.Provide(
			controller.NewPersonController,
			applicationService.NewPersonService,
		),

		// Server Deps
		fx.Provide(
			http.NewRouter,
			http.NewServer,
		),

		fx.Invoke(func(*http.Server) {}),
	).Run()
}
