package main

import (
	"gpt-wework/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", Ping)
	r.GET("/wechat/check", service.CheckWeixinSign)
	r.POST("/wechat/check", service.TalkWeixin)
	r.Run(":80")
}

func Ping(c *gin.Context) {
	c.Data(200, "text/plain;charset=utf-8", []byte("pong"))
}
