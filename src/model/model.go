package models

type Lawyer struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Client struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Folder string `json:"folder"`
	Cpf    string `json:"cpf"`
	Number string `json:"number"`
}
