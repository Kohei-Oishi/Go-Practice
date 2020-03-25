package main

import (
	"Practice/ping"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	//ping, pongの実装
	server.GET("ping", ping.Ping)

	//実行
	server.Run()
}
