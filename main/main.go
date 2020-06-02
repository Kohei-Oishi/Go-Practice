package main

import (
	"Practice/ping"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	//ping, pongの実装
	server.GET("ping", ping.Ping)

	//userの実装
	//server.POST("user/create", user.Createuser)

	//実行
	server.Run()
}
