package students

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name string `json:"name"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}

type CreateStudentDTO struct {
	Name string `json:"name"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}
