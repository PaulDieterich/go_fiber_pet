package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"pauldieterich/go_fiber_pet/pkg/models"
	"strconv"
)

var Db *gorm.DB

func init() {
	port, err := strconv.Atoi("5432")
	if err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		"localhost", port,
		"root", "root",
		"pets", "disable",
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	err = db.AutoMigrate(&models.Pet{})
	if err != nil {
		panic(fmt.Sprintf("Unabled to auto migrate database:", err))
	}

	Db = db

}
