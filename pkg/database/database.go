package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// DB holds the database connection
var DB *sql.DB

// InitDB initializes the PostgreSQL database
func InitDB() error {
	// Get database URL from environment variable
	databaseURL := os.Getenv("POSTGRES_URL")
	if databaseURL == "" {
		// Fallback for local development
		databaseURL = os.Getenv("DATABASE_URL")
		if databaseURL == "" {
			return fmt.Errorf("POSTGRES_URL or DATABASE_URL environment variable is required")
		}
	}

	// Open database connection
	db, err := sql.Open("postgres", databaseURL)
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
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		race VARCHAR(100) NOT NULL,
		class VARCHAR(100) NOT NULL,
		background VARCHAR(100) NOT NULL,
		level INTEGER DEFAULT 1,
		strength INTEGER NOT NULL,
		dexterity INTEGER NOT NULL,
		constitution INTEGER NOT NULL,
		intelligence INTEGER NOT NULL,
		wisdom INTEGER NOT NULL,
		charisma INTEGER NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := DB.Exec(createCharactersTable)
	if err != nil {
		return fmt.Errorf("failed to create characters table: %w", err)
	}

	createCharacterSpellsTable := `
	CREATE TABLE IF NOT EXISTS character_spells (
		id SERIAL PRIMARY KEY,
		character_id INTEGER NOT NULL REFERENCES characters(id) ON DELETE CASCADE,
		spell_name VARCHAR(255) NOT NULL,
		UNIQUE(character_id, spell_name)
	);`

	_, err = DB.Exec(createCharacterSpellsTable)
	if err != nil {
		return fmt.Errorf("failed to create character_spells table: %w", err)
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
