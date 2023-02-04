package students

import "github.com/gin-gonic/gin"

func FindAllController(c *gin.Context) {
	c.JSON(200, FindAllService())
}

func SayHelloController(c *gin.Context) {
	name := c.Param("name")
	c.JSON(200, gin.H{
		"message": "Hello " + name,
	})
}
