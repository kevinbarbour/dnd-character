package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/kevinbarbour/dungeon-and-dragon/internal/character"
	"github.com/kevinbarbour/dungeon-and-dragon/internal/database"
)

// CharacterListHandler shows all saved characters
func (h *WebHandler) CharacterListHandler(w http.ResponseWriter, r *http.Request) {
	characters, err := database.GetAllCharacters()
	if err != nil {
		http.Error(w, "Failed to load characters: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title      string
		Characters []database.SavedCharacter
	}{
		Title:      "Your Characters",
		Characters: characters,
	}

	if err := h.templates.ExecuteTemplate(w, "list.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// CharacterViewHandler shows a specific character
func (h *WebHandler) CharacterViewHandler(w http.ResponseWriter, r *http.Request) {
	// Extract character ID from URL path
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		http.Error(w, "Invalid character ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(pathParts[2])
	if err != nil {
		http.Error(w, "Invalid character ID", http.StatusBadRequest)
		return
	}

	savedChar, err := database.GetCharacterByID(id)
	if err != nil {
		http.Error(w, "Character not found: "+err.Error(), http.StatusNotFound)
		return
	}

	data := struct {
		Title     string
		Character *database.SavedCharacter
	}{
		Title:     savedChar.Character.Name + " - Character Sheet",
		Character: savedChar,
	}

	if err := h.templates.ExecuteTemplate(w, "view.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// CharacterCreateHandler shows the character creation form
func (h *WebHandler) CharacterCreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Show the creation form
		data := struct {
			Title       string
			Races       []character.Race
			Classes     []character.Class
			Backgrounds []character.Background
			Methods     []character.AbilityGenerationMethod
		}{
			Title:       "Create New Character",
			Races:       character.GetAllRaces(),
			Classes:     character.GetAllClasses(),
			Backgrounds: character.GetAllBackgrounds(),
			Methods:     character.GetAllMethods(),
		}

		if err := h.templates.ExecuteTemplate(w, "create_standalone.html", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	if r.Method == "POST" {
		// Process the form submission
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		// Create character from form data
		char, err := h.createCharacterFromForm(r)
		if err != nil {
			http.Error(w, "Failed to create character: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Save to database
		id, err := database.SaveCharacter(char)
		if err != nil {
			http.Error(w, "Failed to save character: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Redirect to the new character's page
		http.Redirect(w, r, fmt.Sprintf("/characters/%d", id), http.StatusSeeOther)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

// CharacterDeleteHandler deletes a character
func (h *WebHandler) CharacterDeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract character ID from URL path
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		http.Error(w, "Invalid character ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(pathParts[2])
	if err != nil {
		http.Error(w, "Invalid character ID", http.StatusBadRequest)
		return
	}

	// Delete the character
	if err := database.DeleteCharacter(id); err != nil {
		http.Error(w, "Failed to delete character: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect to character list
	http.Redirect(w, r, "/characters", http.StatusSeeOther)
}

// GenerateAbilitiesHandler generates ability scores via AJAX
func (h *WebHandler) GenerateAbilitiesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the method from request
	var request struct {
		Method string `json:"method"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Convert method string to enum
	var method character.AbilityGenerationMethod
	switch request.Method {
	case "dice":
		method = character.DiceRoll
	case "pointbuy":
		method = character.PointBuy
	case "standard":
		method = character.StandardArray
	default:
		http.Error(w, "Invalid method", http.StatusBadRequest)
		return
	}

	// Generate abilities
	abilities := character.GenerateAbilityScores(method)

	// Return as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(abilities)
}

// createCharacterFromForm creates a character from form data
func (h *WebHandler) createCharacterFromForm(r *http.Request) (character.Character, error) {
	var char character.Character
	char.Level = 1

	// Get name
	char.Name = strings.TrimSpace(r.FormValue("name"))
	if char.Name == "" {
		return char, fmt.Errorf("character name is required")
	}

	// Get race
	raceName := r.FormValue("race")
	race, found := character.GetRaceByName(raceName)
	if !found {
		return char, fmt.Errorf("invalid race: %s", raceName)
	}
	char.Race = race

	// Get class
	className := r.FormValue("class")
	class, found := character.GetClassByName(className)
	if !found {
		return char, fmt.Errorf("invalid class: %s", className)
	}
	char.Class = class

	// Get background
	backgroundName := r.FormValue("background")
	background, found := character.GetBackgroundByName(backgroundName)
	if !found {
		return char, fmt.Errorf("invalid background: %s", backgroundName)
	}
	char.Background = background

	// Get ability scores
	abilities, err := h.parseAbilities(r)
	if err != nil {
		return char, fmt.Errorf("invalid abilities: %w", err)
	}

	// Apply racial bonuses
	abilities.ApplyRacialBonuses(race)
	char.Abilities = abilities

	return char, nil
}

// parseAbilities parses ability scores from form
func (h *WebHandler) parseAbilities(r *http.Request) (character.AbilityScores, error) {
	var abilities character.AbilityScores

	str, err := strconv.Atoi(r.FormValue("strength"))
	if err != nil || str < 3 || str > 18 {
		return abilities, fmt.Errorf("invalid strength score")
	}
	abilities.Strength = str

	dex, err := strconv.Atoi(r.FormValue("dexterity"))
	if err != nil || dex < 3 || dex > 18 {
		return abilities, fmt.Errorf("invalid dexterity score")
	}
	abilities.Dexterity = dex

	con, err := strconv.Atoi(r.FormValue("constitution"))
	if err != nil || con < 3 || con > 18 {
		return abilities, fmt.Errorf("invalid constitution score")
	}
	abilities.Constitution = con

	intel, err := strconv.Atoi(r.FormValue("intelligence"))
	if err != nil || intel < 3 || intel > 18 {
		return abilities, fmt.Errorf("invalid intelligence score")
	}
	abilities.Intelligence = intel

	wis, err := strconv.Atoi(r.FormValue("wisdom"))
	if err != nil || wis < 3 || wis > 18 {
		return abilities, fmt.Errorf("invalid wisdom score")
	}
	abilities.Wisdom = wis

	cha, err := strconv.Atoi(r.FormValue("charisma"))
	if err != nil || cha < 3 || cha > 18 {
		return abilities, fmt.Errorf("invalid charisma score")
	}
	abilities.Charisma = cha

	return abilities, nil
}
