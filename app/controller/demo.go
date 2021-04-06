package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"rere/app"
)

func Foo(c *gin.Context) {
	//foo := c.Query("foo")
	foo := c.DefaultQuery("foo", "foo")
	app.RespSuccess(c, gin.H{
		"message": fmt.Sprintf("foo: %s", foo),
	})
}

func Hello(c *gin.Context) {
	name := c.Param("name")
	app.RespSuccess(c, gin.H{
		"message": fmt.Sprintf("Hello %s", name),
	})
}

func Form(c *gin.Context) {
	foo := c.PostForm("foo")
	app.RespSuccess(c, gin.H{
		"message": fmt.Sprintf("foo is %s", foo),
	})
}
