package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kevinbarbour/dungeon-and-dragon/internal/character"
	"github.com/kevinbarbour/dungeon-and-dragon/internal/ui"
)

func main() {
	// Show welcome message
	ui.ShowWelcome()

	// Create a new character
	char, err := createCharacter()
	if err != nil {
		log.Fatalf("Error creating character: %v", err)
	}

	// Show final character sheet
	fmt.Println("\n🎉 Character created successfully!")
	fmt.Println(char.String())

	fmt.Println("\nThank you for using the D&D 5e Character Creator!")
	fmt.Println("May your adventures be legendary! 🗡️✨")
}

func createCharacter() (character.Character, error) {
	var char character.Character
	char.Level = 1

	// Step 1: Get character name
	name, err := ui.GetCharacterName()
	if err != nil {
		return char, fmt.Errorf("failed to get character name: %w", err)
	}
	char.Name = name

	// Step 2: Choose ability score generation method
	method, err := ui.SelectAbilityGenerationMethod()
	if err != nil {
		return char, fmt.Errorf("failed to select ability generation method: %w", err)
	}

	// Step 3: Generate ability scores
	abilities := character.GenerateAbilityScores(method)
	ui.ShowAbilityScores(abilities, method)

	err = ui.AskToContinue("Continue with these ability scores?")
	if err != nil {
		return char, fmt.Errorf("user cancelled: %w", err)
	}

	// Step 4: Choose race
	race, err := ui.SelectRace()
	if err != nil {
		return char, fmt.Errorf("failed to select race: %w", err)
	}
	char.Race = race

	// Apply racial bonuses to ability scores
	abilities.ApplyRacialBonuses(race)
	char.Abilities = abilities

	fmt.Printf("\n✨ Applied %s racial bonuses to ability scores!\n", race.Name)
	character.PrintAbilityScores(char.Abilities)

	err = ui.AskToContinue("Continue?")
	if err != nil {
		return char, fmt.Errorf("user cancelled: %w", err)
	}

	// Step 5: Choose class
	class, err := ui.SelectClass()
	if err != nil {
		return char, fmt.Errorf("failed to select class: %w", err)
	}
	char.Class = class

	// Step 6: Choose background
	background, err := ui.SelectBackground()
	if err != nil {
		return char, fmt.Errorf("failed to select background: %w", err)
	}
	char.Background = background

	// Step 7: Confirm character creation
	confirmed, err := ui.ConfirmCharacter(char)
	if err != nil {
		return char, fmt.Errorf("failed to confirm character: %w", err)
	}

	if !confirmed {
		fmt.Println("\nCharacter creation cancelled. Goodbye!")
		os.Exit(0)
	}

	return char, nil
}
