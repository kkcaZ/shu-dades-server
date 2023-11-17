package models

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     Role   `json:"role"`
}

type Role string

const (
	Supplier Role = "supplier"
	Customer Role = "customer"
)
