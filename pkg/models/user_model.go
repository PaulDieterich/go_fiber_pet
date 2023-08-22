package models

type User struct {
	Id       int    `gorm:"primary key; autoIncrement" json:"uuid"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
