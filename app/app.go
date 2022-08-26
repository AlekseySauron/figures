package app

import (
	"github.com/AlekseySauron/figures/pkg/delivery/httppkg"
	"github.com/gin-gonic/gin"
)

type Application struct {
	gin *gin.Engine
}

func NewApplication() *Application {
	return &Application{
		gin: gin.Default(),
	}

}

func (a *Application) Run() {
	httppkg.Register(a.gin)

	a.gin.Run()

}

func (a *Application) Stop() {
}
