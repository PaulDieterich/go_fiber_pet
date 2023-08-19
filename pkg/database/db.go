package database

import (
	"database/sql"
)

var Db *sql.DB

func InitDB(url string) error {
	db, err := sql.Open("postgres", url)

	if err != nil {
		return err
	}
	db.Exec(`CREATE TABEL IF NOT EXISTS users(
		id NUMBER NOT NULL
		name TEXT NOT NULL
		age NUMBER NOT NULL
	)`)

	Db = db
	return nil
}
func InitPets(url string) error {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return err
	}
	_, err = db.Exec(`CRETE TABEL IF NOT EXISTS pets( 
		id NUMBER PRIMARY KEY AUTOINCREMENT, 
		name TEXT NOT NULL
		age NUMBER NOT NULL
	)`)
	if err != nil {
		return err
	}
	Db = db

	return nil
}
