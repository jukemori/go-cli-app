package pkg

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Database connection parameters (update with your PostgreSQL connection details)
const (
	host     = "localhost"
	port     = 5432
	user     = "your_username"
	password = "your_password"
	dbname   = "your_dbname"
)

// OpenDatabase opens a connection to the PostgreSQL database and returns the connection object.
func OpenDatabase() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

// CreateTable creates the necessary table in the PostgreSQL database to store bread data.
func CreateTable(db *sql.DB) error {
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS breads (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			createdAt TIMESTAMP NOT NULL
		)
	`
	_, err := db.Exec(createTableQuery)
	return err
}

// SaveData saves the fetched bread data to the PostgreSQL database.
func SaveData(db *sql.DB, id, name, createdAt string) error {
	insertQuery := "INSERT INTO breads (id, name, createdAt) VALUES ($1, $2, $3)"
	_, err := db.Exec(insertQuery, id, name, createdAt)
	return err
}
