package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/kevinbarbour/dungeon-and-dragon/internal/character"
)

// SavedCharacter represents a character with database metadata
type SavedCharacter struct {
	ID        int
	Character character.Character
	CreatedAt time.Time
	UpdatedAt time.Time
}

// SaveCharacter saves a character to the database
func SaveCharacter(char character.Character) (int, error) {
	query := `
		INSERT INTO characters (
			name, race, class, background, level,
			strength, dexterity, constitution, intelligence, wisdom, charisma
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	result, err := DB.Exec(query,
		char.Name,
		char.Race.Name,
		char.Class.Name,
		char.Background.Name,
		char.Level,
		char.Abilities.Strength,
		char.Abilities.Dexterity,
		char.Abilities.Constitution,
		char.Abilities.Intelligence,
		char.Abilities.Wisdom,
		char.Abilities.Charisma,
	)

	if err != nil {
		return 0, fmt.Errorf("failed to save character: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get character ID: %w", err)
	}

	return int(id), nil
}

// GetAllCharacters retrieves all characters from the database
func GetAllCharacters() ([]SavedCharacter, error) {
	query := `
		SELECT id, name, race, class, background, level,
			   strength, dexterity, constitution, intelligence, wisdom, charisma,
			   created_at, updated_at
		FROM characters
		ORDER BY created_at DESC
	`

	rows, err := DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query characters: %w", err)
	}
	defer rows.Close()

	var characters []SavedCharacter
	for rows.Next() {
		var saved SavedCharacter
		var raceName, className, backgroundName string

		err := rows.Scan(
			&saved.ID,
			&saved.Character.Name,
			&raceName,
			&className,
			&backgroundName,
			&saved.Character.Level,
			&saved.Character.Abilities.Strength,
			&saved.Character.Abilities.Dexterity,
			&saved.Character.Abilities.Constitution,
			&saved.Character.Abilities.Intelligence,
			&saved.Character.Abilities.Wisdom,
			&saved.Character.Abilities.Charisma,
			&saved.CreatedAt,
			&saved.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan character: %w", err)
		}

		// Reconstruct the race, class, and background objects
		if race, found := character.GetRaceByName(raceName); found {
			saved.Character.Race = race
		}
		if class, found := character.GetClassByName(className); found {
			saved.Character.Class = class
		}
		if background, found := character.GetBackgroundByName(backgroundName); found {
			saved.Character.Background = background
		}

		characters = append(characters, saved)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating characters: %w", err)
	}

	return characters, nil
}

// GetCharacterByID retrieves a specific character by ID
func GetCharacterByID(id int) (*SavedCharacter, error) {
	query := `
		SELECT id, name, race, class, background, level,
			   strength, dexterity, constitution, intelligence, wisdom, charisma,
			   created_at, updated_at
		FROM characters
		WHERE id = ?
	`

	var saved SavedCharacter
	var raceName, className, backgroundName string

	err := DB.QueryRow(query, id).Scan(
		&saved.ID,
		&saved.Character.Name,
		&raceName,
		&className,
		&backgroundName,
		&saved.Character.Level,
		&saved.Character.Abilities.Strength,
		&saved.Character.Abilities.Dexterity,
		&saved.Character.Abilities.Constitution,
		&saved.Character.Abilities.Intelligence,
		&saved.Character.Abilities.Wisdom,
		&saved.Character.Abilities.Charisma,
		&saved.CreatedAt,
		&saved.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("character with ID %d not found", id)
		}
		return nil, fmt.Errorf("failed to get character: %w", err)
	}

	// Reconstruct the race, class, and background objects
	if race, found := character.GetRaceByName(raceName); found {
		saved.Character.Race = race
	}
	if class, found := character.GetClassByName(className); found {
		saved.Character.Class = class
	}
	if background, found := character.GetBackgroundByName(backgroundName); found {
		saved.Character.Background = background
	}

	return &saved, nil
}

// UpdateCharacter updates an existing character in the database
func UpdateCharacter(id int, char character.Character) error {
	query := `
		UPDATE characters SET
			name = ?, race = ?, class = ?, background = ?, level = ?,
			strength = ?, dexterity = ?, constitution = ?, intelligence = ?, wisdom = ?, charisma = ?,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = ?
	`

	result, err := DB.Exec(query,
		char.Name,
		char.Race.Name,
		char.Class.Name,
		char.Background.Name,
		char.Level,
		char.Abilities.Strength,
		char.Abilities.Dexterity,
		char.Abilities.Constitution,
		char.Abilities.Intelligence,
		char.Abilities.Wisdom,
		char.Abilities.Charisma,
		id,
	)

	if err != nil {
		return fmt.Errorf("failed to update character: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("character with ID %d not found", id)
	}

	return nil
}

// DeleteCharacter deletes a character from the database
func DeleteCharacter(id int) error {
	query := `DELETE FROM characters WHERE id = ?`

	result, err := DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete character: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("character with ID %d not found", id)
	}

	return nil
}

// GetCharacterCount returns the total number of saved characters
func GetCharacterCount() (int, error) {
	query := `SELECT COUNT(*) FROM characters`

	var count int
	err := DB.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to get character count: %w", err)
	}

	return count, nil
}
