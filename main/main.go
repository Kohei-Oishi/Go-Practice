package main

import (
	"Practice/calculation"
	"Practice/ping"
	"Practice/returnname"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	//ping, pongの実装
	server.GET("/ping", ping.Ping)

	//名前を入れたら名前を返すものの実装
	server.POST("/returnname", returnname.Retuenname)

	//足し算の実装
	server.POST("/sum", calculation.Sum)

	//2乗計算の実装
	server.POST("/square", calculation.Square)

	//実行
	server.Run()
}
