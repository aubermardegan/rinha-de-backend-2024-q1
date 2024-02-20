package repository

import (
	"database/sql"
	"fmt"

	"github.com/amardegan/rinha-de-backend-2024-q1/config"
	_ "github.com/lib/pq"
)

func InitPostgreSQL() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DATABASE_HOST, config.DATABASE_PORT, config.DATABASE_USER, config.DATABASE_PASSWORD, config.DATABASE_NAME)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	return db, err
}
