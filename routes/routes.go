package routes

import (
	"github.com/LucasRMP/golang-studies-gin-rest-api/students"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine) {
	engine.GET("/students", students.FindAllController)
	engine.GET("/students/:name", students.SayHelloController)
}
