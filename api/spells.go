package handler

import (
	"encoding/json"
	"net/http"

	"github.com/kevinbarbour/dungeon-and-dragon/internal/character"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get class name from query parameter
	className := r.URL.Query().Get("class")
	if className == "" {
		http.Error(w, "Class parameter is required", http.StatusBadRequest)
		return
	}

	// Get spells for the class
	spellList := character.GetSpellsForClass(className)

	// Return as JSON
	w.Header().Set("Content-Type", "application/json")
	response := struct {
		Cantrips    []character.Spell `json:"cantrips"`
		FirstLevel  []character.Spell `json:"firstLevel"`
		BonusSpells []character.Spell `json:"bonusSpells"`
	}{
		Cantrips:    spellList.Cantrips,
		FirstLevel:  spellList.FirstLevel,
		BonusSpells: spellList.BonusSpells,
	}
	json.NewEncoder(w).Encode(response)
}
