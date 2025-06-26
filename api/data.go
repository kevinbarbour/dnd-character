package handler

import (
	"encoding/json"
	"net/http"

	"github.com/kevinbarbour/dungeon-and-dragon/pkg/character"
	"github.com/kevinbarbour/dungeon-and-dragon/pkg/database"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Initialize database connection
	if err := database.InitDB(); err != nil {
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}
	defer database.CloseDB()

	// Get character count for display
	count, err := database.GetCharacterCount()
	if err != nil {
		count = 0 // Default to 0 if error
	}

	data := struct {
		Title          string                 `json:"title"`
		CharacterCount int                    `json:"characterCount"`
		Races          []character.Race       `json:"races"`
		Classes        []character.Class      `json:"classes"`
		Backgrounds    []character.Background `json:"backgrounds"`
	}{
		Title:          "D&D 5e Character Creator",
		CharacterCount: count,
		Races:          character.GetAllRaces(),
		Classes:        character.GetAllClasses(),
		Backgrounds:    character.GetAllBackgrounds(),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "JSON encoding failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
