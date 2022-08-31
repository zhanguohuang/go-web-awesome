package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhanguohuang/go-web-awesome/handler"
)

func main() {
	route := gin.Default()
	route.GET("/testing", handler.StartPage)
	route.Run(":8085")
}
