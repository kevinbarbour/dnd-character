package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/kevinbarbour/dungeon-and-dragon/internal/database"
)

// WebHandler holds the template and any dependencies
type WebHandler struct {
	templates *template.Template
}

// NewWebHandler creates a new web handler with parsed templates
func NewWebHandler() (*WebHandler, error) {
	// Create function map for templates
	funcMap := template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
	}

	// Parse base template first, then all others
	templates := template.New("").Funcs(funcMap)

	// Parse base template
	_, err := templates.ParseFiles(filepath.Join("web", "templates", "base.html"))
	if err != nil {
		return nil, err
	}

	// Parse all other templates
	_, err = templates.ParseGlob(filepath.Join("web", "templates", "*.html"))
	if err != nil {
		return nil, err
	}

	return &WebHandler{
		templates: templates,
	}, nil
}

// HomeHandler serves the home page
func (h *WebHandler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Get character count for display
	count, err := database.GetCharacterCount()
	if err != nil {
		count = 0 // Default to 0 if error
	}

	data := struct {
		Title          string
		CharacterCount int
	}{
		Title:          "D&D 5e Character Creator",
		CharacterCount: count,
	}

	if err := h.templates.ExecuteTemplate(w, "home.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// StaticHandler serves static files
func StaticHandler(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/static/", http.FileServer(http.Dir("web/static/"))).ServeHTTP(w, r)
}
