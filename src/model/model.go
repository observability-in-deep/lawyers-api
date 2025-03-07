package models

import "time"

type Lawyer struct {
	ID       string     `json:"id" db:"id"`
	Name     string     `json:"name" db:"name"`
	Email    string     `json:"email" db:"email"`
	OAB      string     `json:"oab" db:"oab"`
	Phone    string     `json:"phone" db:"phone"`
	CreateAt *time.Time `json:"create_at" db:"create_at"`
	UpdateAt *time.Time `json:"update_at" db:"update_at"`
	DeleteAt *time.Time `json:"delete_at" db:"delete_at"`
}

type Customers struct {
	ID         string     `json:"id" db:"id"`
	Name       string     `json:"name" db:"name"`
	Phone      string     `json:"phone" db:"phone"`
	Email      string     `json:"email" db:"email"`
	Folder     string     `json:"folder" db:"folder"`
	Cpf        string     `json:"cpf" db:"cpf"`
	UpdateAt   *time.Time `json:"update_at" db:"update_at"`
	CreateAt   *time.Time `json:"create_at" db:"create_at"`
	LawyerName string     `json:"lawyer_name" db:"lawyer_name"`
}
