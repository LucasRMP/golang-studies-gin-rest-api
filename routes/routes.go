package routes

import (
	"github.com/LucasRMP/golang-studies-gin-rest-api/students"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine) {
	engine.POST("/students", students.CreateController)
	engine.GET("/students", students.FindAllController)
	engine.GET("/students/:id", students.FindByIdController)
	engine.GET("/students/cpf/:cpf", students.FindByCPFController)
	engine.PUT("/students/:id", students.UpdateController)
	engine.DELETE("/students/:id", students.DeleteController)
}
