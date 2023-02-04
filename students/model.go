package students

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}

var Students []Student
