package database

import (
	"aviatoV3/internal/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func DB() *sql.DB {
	dbConfig := config.GetConfig()
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Password, dbConfig.Database)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil
	}
	return db
}
