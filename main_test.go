package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/LucasRMP/golang-studies-gin-rest-api/database"
	"github.com/LucasRMP/golang-studies-gin-rest-api/students"
	testUtils "github.com/LucasRMP/golang-studies-gin-rest-api/utils"
	"github.com/stretchr/testify/assert"
)

var studentId int

func SetupMock() {
	database.Connect()
	student := students.Student{
		Name: "Lucas Pessone",
		CPF:  "10987654321",
		RG:   "123456789",
	}
	database.DB.Create(&student)
	studentId = int(student.ID)
}

func TearDownMock() {
	mockStudent := students.Student{}
	database.DB.Delete(&mockStudent, studentId)
}

func TestFindAll(t *testing.T) {
	SetupMock()
	defer TearDownMock()
	engine := testUtils.GetTestEngine()
	engine.GET("/students", students.FindAllController)

	request, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()
	engine.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestFindByCPF(t *testing.T) {
	SetupMock()
	defer TearDownMock()
	engine := testUtils.GetTestEngine()
	engine.GET("/students/cpf/:cpf", students.FindByCPFController)

	request, _ := http.NewRequest("GET", "/students/cpf/10987654321", nil)
	response := httptest.NewRecorder()
	engine.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestFindByID(t *testing.T) {
	SetupMock()
	defer TearDownMock()
	engine := testUtils.GetTestEngine()
	engine.GET("/students/:id", students.FindByIdController)

	path := "/students/" + strconv.Itoa(studentId)
	request, _ := http.NewRequest("GET", path, nil)
	response := httptest.NewRecorder()
	engine.ServeHTTP(response, request)

	student := students.Student{}
	json.Unmarshal(response.Body.Bytes(), &student)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, int(student.ID), studentId)
	assert.Equal(t, student.Name, "Lucas Pessone")
	assert.Equal(t, student.CPF, "10987654321")
	assert.Equal(t, student.RG, "123456789")
}

func TestDelete(t *testing.T) {
	SetupMock()
	engine := testUtils.GetTestEngine()
	engine.DELETE("/students/:id", students.DeleteController)

	path := "/students/" + strconv.Itoa(studentId)
	request, _ := http.NewRequest("DELETE", path, nil)
	response := httptest.NewRecorder()
	engine.ServeHTTP(response, request)

	assert.Equal(t, http.StatusNoContent, response.Code)
}

func TestEdit(t *testing.T) {
	SetupMock()
	defer TearDownMock()
	engine := testUtils.GetTestEngine()
	engine.PUT("/students/:id", students.UpdateController)

	newName := "This is a updated Name"
	body, _ := json.Marshal(students.UpdateStudentDTO{
		Name: newName,
	})

	path := "/students/" + strconv.Itoa(studentId)
	request, _ := http.NewRequest("PUT", path, bytes.NewBuffer(body))
	response := httptest.NewRecorder()
	engine.ServeHTTP(response, request)

	student := students.Student{}
	json.Unmarshal(response.Body.Bytes(), &student)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, int(student.ID), studentId)
	assert.Equal(t, student.Name, newName)
	assert.Equal(t, student.CPF, "10987654321")
	assert.Equal(t, student.RG, "123456789")

}
