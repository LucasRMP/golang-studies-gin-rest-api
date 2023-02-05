package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindAllController(c *gin.Context) {
	c.JSON(http.StatusOK, FindAllService())
}

func CreateController(c *gin.Context) {
	var studentDTO CreateStudentDTO
	if err := c.ShouldBindJSON(&studentDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student := CreateService(studentDTO)

	c.JSON(http.StatusCreated, student)
}
