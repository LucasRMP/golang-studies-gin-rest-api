package main

import (
	"github.com/LucasRMP/golang-studies-gin-rest-api/routes"
	"github.com/LucasRMP/golang-studies-gin-rest-api/students"
	"github.com/gin-gonic/gin"
)

func main() {
	students.Students = []students.Student{
		{ID: 1, Name: "Lucas", CPF: "123456789", RG: "987654321"},
		{ID: 2, Name: "João", CPF: "123456789", RG: "987654321"},
		{ID: 3, Name: "Maria", CPF: "123456789", RG: "987654321"},
	}
	startServer()
}

func startServer() {
	engine := gin.Default()

	routes.RegisterRoutes(engine)

	engine.Run(":3333")
}
