package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateController(c *gin.Context) {
	var studentDTO CreateStudentDTO

	if err := c.ShouldBindJSON(&studentDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ValidateInterface(studentDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student"})
		return
	}

	student := CreateService(studentDTO)

	c.JSON(http.StatusCreated, student)
}

func FindAllController(c *gin.Context) {
	c.JSON(http.StatusOK, FindAllService())
}

func FindByIdController(c *gin.Context) {
	id := c.Param("id")
	student, err := FindByIdService(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, student)
}

func FindByCPFController(c *gin.Context) {
	cpf := c.Param("cpf")
	student, err := FindByCPFService(cpf)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, student)
}

func UpdateController(c *gin.Context) {
	id := c.Param("id")
	var studentDTO UpdateStudentDTO
	if err := c.ShouldBindJSON(&studentDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student, err := UpdateService(id, studentDTO)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := ValidateInterface(studentDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student"})
		return
	}

	c.JSON(http.StatusOK, student)
}

func DeleteController(c *gin.Context) {
	id := c.Param("id")
	err := DeleteService(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
