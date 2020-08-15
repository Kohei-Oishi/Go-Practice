package returnname

import (
	"github.com/gin-gonic/gin"
)

type Human struct {
	Name string `json:"name"`
}

func Retuenname(c *gin.Context)  {
	var name Human
	err := c.BindJSON(&name)
	if err != nil {
		c.JSON(400, gin.H{
			"message" : "そんな奴存在しねぇよバーカ",
		})
		return
	}
	c.JSON(200, gin.H{
		"Name" : name.Name,
	})
	c.JSON(400, gin.H{
		"message" : "ごめんな、お前には用はないんや",
	})
}
