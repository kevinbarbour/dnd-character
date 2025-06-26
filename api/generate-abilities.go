package handler

import (
	"encoding/json"
	"net/http"

	"github.com/kevinbarbour/dungeon-and-dragon/internal/character"
)

func Handler(w http.ResponseWriter, r *http.Request) {
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
