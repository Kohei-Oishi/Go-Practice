package calculation

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type NumberRequest struct {
	Number1 int `json:"number_1"`
	Number2 int `json:"number_2"`
}

type SumNumberResponce struct {
	Number int `json:"number"`
}

func Sum(c *gin.Context) {
	var number NumberRequest
	var sumnum SumNumberResponce
	err := c.ShouldBindJSON(&number)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : "数字がなぁ、ちんこ",
		})
		return
	}
	sumnum.Number = number.Number1 + number.Number2
	c.JSON(http.StatusOK, gin.H{
		"SumNumber" : sumnum.Number,
	})
}
