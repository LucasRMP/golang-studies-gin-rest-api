package students

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name string `json:"name"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}

type CreateStudentDTO struct {
	Name string `json:"name" validate:"nonzero"`
	CPF  string `json:"cpf" validate:"len=11,regexp=^[0-9]+$"`
	RG   string `json:"rg" validate:"len=9,regexp=^[0-9]+$"`
}

type UpdateStudentDTO struct {
	Name string `json:"name" validate:"nonzero"`
}

func ValidateInterface(input interface{}) error {
	return validator.Validate(input)
}
