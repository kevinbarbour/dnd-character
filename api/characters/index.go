package handler

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/kevinbarbour/dungeon-and-dragon/internal/database"
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

	// Load template
	tmplPath := filepath.Join("public", "templates", "list_standalone.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Template loading failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title      string
		Characters []database.SavedCharacter
	}{
		Title:      "Your Characters",
		Characters: characters,
	}

	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Template execution failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
