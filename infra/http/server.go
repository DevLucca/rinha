package http

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Server struct{}

func NewServer(lc fx.Lifecycle, router *Router) *Server {
	srv := gin.Default()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			router.Load(srv)
			go srv.Run()
			return nil
		},
	})
	return &Server{}
}
