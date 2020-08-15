package calculation

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type SquareNumber struct {
	Number int `json:"number"`
}

func Square(c *gin.Context) {
	var k SquareNumber
	// k.Number = k.Number * k.Number
	err := c.BindJSON(&k)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"message" : "死ね",
		})
		return
	}
	c.JSON(200, gin.H{
		"SquareNumber" : k.Number,
	})
	c.JSON(400, gin.H{
		"message" : "死ね!!消えろ!!",
	})
}
