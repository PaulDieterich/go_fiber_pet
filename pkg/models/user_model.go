package models

type Role struct {
	Id   int    `gorm:"primary key; autoIncrement" json:"id"`
	Name string `gorm:"name"`
}

type User struct {
	Id       int    `gorm:"primary key; autoIncrement" json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     Role   `gorm:"foreignKey:Id;references:Id" json:"role"`
}
