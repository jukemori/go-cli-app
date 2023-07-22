package pkg

import (
	"database/sql"
	"fmt"
	"os"
	"time"
	_ "github.com/lib/pq"
)

// OpenDatabase opens a connection to the PostgreSQL database and returns the connection object.
func OpenDatabase() (*sql.DB, error) {
	// Read the database credentials from environment variables
	host := "localhost"
	port := 5432
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

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

func CreateTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS breads (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL
		);
	`)
	return err
}

// SaveData saves the fetched bread data to the PostgreSQL database.
func SaveData(db *sql.DB, id, name, createdAt string) error {
	fmt.Println("Data to be inserted:")
	fmt.Println("ID:", id)
	fmt.Println("Name:", name)
	fmt.Println("CreatedAt:", createdAt)

	// Convert createdAt string to time.Time
	createdAtTime, err := time.Parse(time.RFC3339, createdAt)
	if err != nil {
		return fmt.Errorf("failed to parse createdAt: %v", err)
	}

	insertQuery := `
		INSERT INTO contentful_entries (id, name, created_at)
		VALUES ($1, $2, $3)
		ON CONFLICT (id) DO UPDATE
		SET name = excluded.name, created_at = excluded.created_at
	`

	_, err = db.Exec(insertQuery, id, name, createdAtTime)
	return err
}


// Rest of the functions remain the same...
