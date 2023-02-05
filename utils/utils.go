package utils

import (
	"github.com/gin-gonic/gin"
)

func GetTestEngine() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	return engine
}
