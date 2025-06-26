package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

const dbFileName = "dnd_characters.db"

// DB holds the database connection
var DB *sql.DB

// InitDB initializes the SQLite database
func InitDB() error {
	// Get the directory where the executable is located
	execPath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get executable path: %w", err)
	}

	dbPath := filepath.Join(filepath.Dir(execPath), dbFileName)

	// Open database connection
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	DB = db

	// Test the connection
	if err := DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	// Create tables if they don't exist
	if err := createTables(); err != nil {
		return fmt.Errorf("failed to create tables: %w", err)
	}

	return nil
}

// createTables creates the necessary database tables
func createTables() error {
	createCharactersTable := `
	CREATE TABLE IF NOT EXISTS characters (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		race TEXT NOT NULL,
		class TEXT NOT NULL,
		background TEXT NOT NULL,
		level INTEGER DEFAULT 1,
		strength INTEGER NOT NULL,
		dexterity INTEGER NOT NULL,
		constitution INTEGER NOT NULL,
		intelligence INTEGER NOT NULL,
		wisdom INTEGER NOT NULL,
		charisma INTEGER NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := DB.Exec(createCharactersTable)
	if err != nil {
		return fmt.Errorf("failed to create characters table: %w", err)
	}

	return nil
}

// CloseDB closes the database connection
func CloseDB() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}

// GetDBPath returns the path to the database file
func GetDBPath() (string, error) {
	execPath, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Join(filepath.Dir(execPath), dbFileName), nil
}
