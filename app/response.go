// 统一的返参调用方法
package app

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	Success = 200
	Error = 500
)

type Response struct {
	Code        int         `json:"code"`
	Message     string      `json:"message"`
	CurrentTime int64       `json:"current_time"`
	Data        interface{} `json:"data"`
}

// RespSuccess 成功返回
func RespSuccess(c *gin.Context, data interface{}) {
	if data == nil {
		data = gin.H{}
	}

	resp := &Response{
		Code:        Success,
		Message:     "success",
		CurrentTime: time.Now().Unix(),
		Data:        data,
	}
	c.JSON(http.StatusOK, resp)
}
