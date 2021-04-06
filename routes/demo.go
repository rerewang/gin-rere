package routes

import (
	"github.com/gin-gonic/gin"
	"rere/app/controller"
)

// initDemo .
func initDemo(g *gin.Engine) {
	route := g.Group("/demo")

	route.GET("/foo", controller.Foo)
	route.GET("/hello/:name", controller.Hello)
	route.POST("/form", controller.Form)
}