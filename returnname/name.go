package returnname

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Human struct {
	Name string `json:"name"`
}

func Retuenname(c *gin.Context)  {
	var name Human
	err := c.ShouldBindJSON(&name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : "そんな奴存在しねぇよバーカ",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Name" : name.Name,
	})
}
