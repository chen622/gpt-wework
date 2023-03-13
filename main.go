package main

import (
	"gpt-wework/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/health", Ping)
	r.GET("/wework/gpt-api", service.CheckWeixinSign)
	r.POST("/wework/gpt-api", service.TalkWeixin)
	r.Run()
}

func Ping(c *gin.Context) {
	c.Data(200, "text/plain;charset=utf-8", []byte("pong"))
}
