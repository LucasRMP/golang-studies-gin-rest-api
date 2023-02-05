package main

import (
	"github.com/LucasRMP/golang-studies-gin-rest-api/database"
	"github.com/LucasRMP/golang-studies-gin-rest-api/routes"
	"github.com/LucasRMP/golang-studies-gin-rest-api/students"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&students.Student{})

	startServer()
}

func startServer() {
	engine := gin.Default()

	routes.RegisterRoutes(engine)

	engine.Run(":3333")
}
