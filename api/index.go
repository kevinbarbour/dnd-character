package handler

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/kevinbarbour/dungeon-and-dragon/internal/character"
	"github.com/kevinbarbour/dungeon-and-dragon/internal/database"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Initialize database connection
	if err := database.InitDB(); err != nil {
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}
	defer database.CloseDB()

	// Load template
	tmplPath := filepath.Join("public", "templates", "home.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Template loading failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Get character count for display
	count, err := database.GetCharacterCount()
	if err != nil {
		count = 0 // Default to 0 if error
	}

	data := struct {
		Title          string
		CharacterCount int
		Races          []character.Race
		Classes        []character.Class
		Backgrounds    []character.Background
	}{
		Title:          "D&D 5e Character Creator",
		CharacterCount: count,
		Races:          character.GetAllRaces(),
		Classes:        character.GetAllClasses(),
		Backgrounds:    character.GetAllBackgrounds(),
	}

	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Template execution failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
