package http

import (
	"github.com/DevLucca/rinha/application/controller"
	"github.com/gin-gonic/gin"
)

type Router struct {
	personCtrl *controller.PersonController
}

func NewRouter(ctrl *controller.PersonController) *Router {
	return &Router{
		personCtrl: ctrl,
	}
}

func (r *Router) Load(g *gin.Engine) {
	g.POST("/pessoas/", r.personCtrl.Create)
	g.GET("/pessoas/:id", r.personCtrl.Retrieve)
	g.GET("/pessoas/", r.personCtrl.Search)
	g.GET("/contagem-pessoas/", r.personCtrl.Count)
}
