package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite3", "./cinematicket.db")
	if err != nil {
		return err
	}

	// Чтение и выполнение миграций
	migrationSQL, err := os.ReadFile("database/migrations.sql")
	if err != nil {
		return err
	}

	_, err = DB.Exec(string(migrationSQL))
	if err != nil {
		return err
	}

	log.Println("Database initialized successfully")
	return nil
}
