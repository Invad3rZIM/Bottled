package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DatabaseConnection struct {
	db *sql.DB
}

func NewDatabaseConnection(db *sql.DB) *DatabaseConnection {
	return &DatabaseConnection{
		db: db,
	}
}
