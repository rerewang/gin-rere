package api

import (
	"github.com/gin-gonic/gin"
	"rere/internal/api/route"
)

func Init(g *gin.Engine) error {
	route.Init(g)
	return nil
}