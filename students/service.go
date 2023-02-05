package students

import (
	"github.com/LucasRMP/golang-studies-gin-rest-api/database"
)

func FindAllService() []Student {
	return []Student{
		{
			Name: "John Doe",
			CPF:  "123.456.789-00",
			RG:   "12.345.678-9",
		},
	}
}

func CreateService(input CreateStudentDTO) Student {
	studentToCreate := Student{
		Name: input.Name,
		CPF:  input.CPF,
		RG:   input.RG,
	}

	database.DB.Create(&studentToCreate)
	return studentToCreate
}
