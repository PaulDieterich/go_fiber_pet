package models

type Pet struct {
	Id   int    `gorm:"primary key; autoIncrement" json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
