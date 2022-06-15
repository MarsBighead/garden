package router

import (
	"garden/config"

	"github.com/gin-gonic/gin"
)

type Act interface {
	Post(c *gin.Context)
	Get(c *gin.Context)
}

type Router func(env *config.Environment) Act

var Routers = map[string]Router{}

func Add(name string, r Router) {
	Routers[name] = r
}
