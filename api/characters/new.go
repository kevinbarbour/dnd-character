package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

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

	if r.Method == "GET" {
		// Show the creation form
		tmplPath := filepath.Join("public", "templates", "create_standalone.html")
		tmpl, err := template.ParseFiles(tmplPath)
		if err != nil {
			http.Error(w, "Template loading failed: "+err.Error(), http.StatusInternalServerError)
			return
		}

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

		w.Header().Set("Content-Type", "text/html")
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, "Template execution failed: "+err.Error(), http.StatusInternalServerError)
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
		char, err := createCharacterFromForm(r)
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

// createCharacterFromForm creates a character from form data
func createCharacterFromForm(r *http.Request) (character.Character, error) {
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
	abilities, err := parseAbilities(r)
	if err != nil {
		return char, fmt.Errorf("invalid abilities: %w", err)
	}

	// Apply racial bonuses
	abilities.ApplyRacialBonuses(race)
	char.Abilities = abilities

	// Get spells if the class is a spellcaster
	if character.IsSpellcaster(className) {
		spells := []string{}

		// Get cantrips
		cantrips := r.Form["cantrips"]
		spells = append(spells, cantrips...)

		// Get first level spells
		firstLevelSpells := r.Form["spells"]
		spells = append(spells, firstLevelSpells...)

		char.Spells = spells
	}

	return char, nil
}

// parseAbilities parses ability scores from form
func parseAbilities(r *http.Request) (character.AbilityScores, error) {
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

// GenerateAbilitiesHandler generates ability scores via AJAX
func GenerateAbilitiesHandler(w http.ResponseWriter, r *http.Request) {
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
