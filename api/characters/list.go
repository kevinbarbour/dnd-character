package handler

import (
	"encoding/json"
	"net/http"

	"github.com/kevinbarbour/dungeon-and-dragon/pkg/database"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Initialize database connection
	if err := database.InitDB(); err != nil {
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}
	defer database.CloseDB()

	// Get all characters
	characters, err := database.GetAllCharacters()
	if err != nil {
		http.Error(w, "Failed to load characters: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title      string                    `json:"title"`
		Characters []database.SavedCharacter `json:"characters"`
	}{
		Title:      "Your Characters",
		Characters: characters,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "JSON encoding failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
