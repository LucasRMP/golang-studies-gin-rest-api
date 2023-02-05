package students

import (
	"github.com/LucasRMP/golang-studies-gin-rest-api/database"
)

func CreateService(input CreateStudentDTO) Student {
	studentToCreate := Student{
		Name: input.Name,
		CPF:  input.CPF,
		RG:   input.RG,
	}

	database.DB.Create(&studentToCreate)
	return studentToCreate
}

func FindAllService() []Student {
	var students []Student
	database.DB.Find(&students)
	return students
}

func FindByIdService(id string) (Student, error) {
	var student Student
	result := database.DB.First(&student, id)
	if result.Error != nil {
		return student, result.Error
	}
	return student, nil
}

func FindByCPFService(cpf string) (Student, error) {
	var student Student
	result := database.DB.Where("cpf = ?", cpf).First(&student)
	if result.Error != nil {
		return student, result.Error
	}
	return student, nil
}

func DeleteService(id string) error {
	var student Student
	result := database.DB.Delete(&student, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateService(id string, input UpdateStudentDTO) (Student, error) {
	var student Student
	result := database.DB.First(&student, id)
	if result.Error != nil {
		return student, result.Error
	}

	if input.Name != "" {
		student.Name = input.Name
	}
	database.DB.Save(&student)
	return student, nil
}
