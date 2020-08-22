package calculation

import (
	"github.com/gin-gonic/gin"
)

type SquareNumber struct {
	Number int `json:"number"`
}

func Square(c *gin.Context) {
	var k SquareNumber
	err := c.BindJSON(&k)
	if err != nil {
		return
	}
	k.Number = k.Number * k.Number
	c.JSON(200, gin.H{
		"SquareNumber" : k.Number,
	})
}
