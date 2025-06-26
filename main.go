package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/kevinbarbour/dungeon-and-dragon/internal/character"
	"github.com/kevinbarbour/dungeon-and-dragon/internal/database"
	"github.com/kevinbarbour/dungeon-and-dragon/internal/ui"
	"github.com/kevinbarbour/dungeon-and-dragon/web/handlers"
)

func main() {
	// Parse command line flags
	webMode := flag.Bool("web", false, "Run in web server mode")
	port := flag.String("port", "8080", "Port to run web server on")
	flag.Parse()

	// Initialize database
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.CloseDB()

	if *webMode {
		startWebServer(*port)
	} else {
		startCLI()
	}
}

func startWebServer(port string) {
	fmt.Printf("🌐 Starting D&D Character Creator Web Server on port %s...\n", port)

	// Initialize web handler
	webHandler, err := handlers.NewWebHandler()
	if err != nil {
		log.Fatalf("Failed to initialize web handler: %v", err)
	}

	// Setup routes
	http.HandleFunc("/", webHandler.HomeHandler)
	http.HandleFunc("/characters", webHandler.CharacterListHandler)
	http.HandleFunc("/characters/new", webHandler.CharacterCreateHandler)
	http.HandleFunc("/characters/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/characters/" {
			webHandler.CharacterListHandler(w, r)
			return
		}

		// Handle character view and delete
		if r.Method == "POST" && len(r.URL.Path) > 12 && r.URL.Path[len(r.URL.Path)-7:] == "/delete" {
			webHandler.CharacterDeleteHandler(w, r)
		} else {
			webHandler.CharacterViewHandler(w, r)
		}
	})
	http.HandleFunc("/api/generate-abilities", webHandler.GenerateAbilitiesHandler)
	http.HandleFunc("/static/", handlers.StaticHandler)

	fmt.Printf("🚀 Server running at http://localhost:%s\n", port)
	fmt.Println("📝 Create characters at http://localhost:" + port + "/characters/new")
	fmt.Println("👥 View characters at http://localhost:" + port + "/characters")
	fmt.Println("\nPress Ctrl+C to stop the server")

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start web server: %v", err)
	}
}

func startCLI() {
	// Show welcome message
	ui.ShowWelcome()

	// Main application loop
	for {
		option, err := ui.ShowMainMenu()
		if err != nil {
			log.Fatalf("Error showing main menu: %v", err)
		}

		switch option {
		case ui.CreateNewCharacter:
			if err := handleCreateCharacter(); err != nil {
				fmt.Printf("Error creating character: %v\n", err)
				ui.AskToContinue("Press Enter to continue")
			}
		case ui.ManageSavedCharacters:
			if err := handleCharacterManagement(); err != nil {
				fmt.Printf("Error in character management: %v\n", err)
				ui.AskToContinue("Press Enter to continue")
			}
		case ui.Exit:
			fmt.Println("\nThank you for using the D&D 5e Character Creator!")
			fmt.Println("May your adventures be legendary! 🗡️✨")
			return
		}
	}
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

// handleCreateCharacter handles the character creation flow
func handleCreateCharacter() error {
	char, err := createCharacter()
	if err != nil {
		return err
	}

	// Ask if user wants to save the character
	save, err := ui.ConfirmSaveCharacter()
	if err != nil {
		return fmt.Errorf("failed to confirm save: %w", err)
	}

	if save {
		id, err := database.SaveCharacter(char)
		if err != nil {
			return fmt.Errorf("failed to save character: %w", err)
		}
		fmt.Printf("\n💾 Character saved successfully with ID: %d\n", id)
	}

	// Show final character sheet
	fmt.Println("\n🎉 Character created successfully!")
	fmt.Println(char.String())

	return ui.AskToContinue("Press Enter to continue")
}

// handleCharacterManagement handles the character management menu
func handleCharacterManagement() error {
	for {
		option, err := ui.ShowCharacterManagementMenu()
		if err != nil {
			return err
		}

		switch option {
		case ui.ListCharacters:
			if err := handleListCharacters(); err != nil {
				fmt.Printf("Error listing characters: %v\n", err)
				ui.AskToContinue("Press Enter to continue")
			}
		case ui.ViewCharacter:
			if err := handleViewCharacter(); err != nil {
				fmt.Printf("Error viewing character: %v\n", err)
				ui.AskToContinue("Press Enter to continue")
			}
		case ui.EditCharacter:
			if err := handleEditCharacter(); err != nil {
				fmt.Printf("Error editing character: %v\n", err)
				ui.AskToContinue("Press Enter to continue")
			}
		case ui.DeleteCharacter:
			if err := handleDeleteCharacter(); err != nil {
				fmt.Printf("Error deleting character: %v\n", err)
				ui.AskToContinue("Press Enter to continue")
			}
		case ui.BackToMainMenu:
			return nil
		}
	}
}

// handleListCharacters lists all saved characters
func handleListCharacters() error {
	characters, err := database.GetAllCharacters()
	if err != nil {
		return fmt.Errorf("failed to get characters: %w", err)
	}

	ui.ShowCharacterList(characters)
	return ui.AskToContinue("Press Enter to continue")
}

// handleViewCharacter shows details of a specific character
func handleViewCharacter() error {
	characters, err := database.GetAllCharacters()
	if err != nil {
		return fmt.Errorf("failed to get characters: %w", err)
	}

	selected, err := ui.SelectSavedCharacter(characters, "view")
	if err != nil {
		return err
	}

	ui.ShowCharacterDetails(selected)
	return ui.AskToContinue("Press Enter to continue")
}

// handleEditCharacter allows editing of a character
func handleEditCharacter() error {
	characters, err := database.GetAllCharacters()
	if err != nil {
		return fmt.Errorf("failed to get characters: %w", err)
	}

	selected, err := ui.SelectSavedCharacter(characters, "edit")
	if err != nil {
		return err
	}

	fmt.Printf("\n📝 Editing character: %s\n", selected.Character.Name)
	fmt.Println("Note: This will create a new character with the same base stats.")
	fmt.Println("You can modify the race, class, and background.")

	// Create a copy of the character for editing
	editChar := selected.Character

	// Allow user to change race
	fmt.Println("\n🔄 Select new race (current: " + editChar.Race.Name + ")")
	newRace, err := ui.SelectRace()
	if err != nil {
		return fmt.Errorf("failed to select new race: %w", err)
	}

	// Reset abilities to base values (remove old racial bonuses)
	baseAbilities := editChar.Abilities
	baseAbilities.Strength -= editChar.Race.StrengthBonus
	baseAbilities.Dexterity -= editChar.Race.DexterityBonus
	baseAbilities.Constitution -= editChar.Race.ConstitutionBonus
	baseAbilities.Intelligence -= editChar.Race.IntelligenceBonus
	baseAbilities.Wisdom -= editChar.Race.WisdomBonus
	baseAbilities.Charisma -= editChar.Race.CharismaBonus

	// Apply new racial bonuses
	editChar.Race = newRace
	baseAbilities.ApplyRacialBonuses(newRace)
	editChar.Abilities = baseAbilities

	// Allow user to change class
	fmt.Println("\n🔄 Select new class (current: " + editChar.Class.Name + ")")
	newClass, err := ui.SelectClass()
	if err != nil {
		return fmt.Errorf("failed to select new class: %w", err)
	}
	editChar.Class = newClass

	// Allow user to change background
	fmt.Println("\n🔄 Select new background (current: " + editChar.Background.Name + ")")
	newBackground, err := ui.SelectBackground()
	if err != nil {
		return fmt.Errorf("failed to select new background: %w", err)
	}
	editChar.Background = newBackground

	// Confirm changes
	confirmed, err := ui.ConfirmCharacter(editChar)
	if err != nil {
		return fmt.Errorf("failed to confirm changes: %w", err)
	}

	if confirmed {
		err = database.UpdateCharacter(selected.ID, editChar)
		if err != nil {
			return fmt.Errorf("failed to update character: %w", err)
		}
		fmt.Printf("\n✅ Character '%s' updated successfully!\n", editChar.Name)
	} else {
		fmt.Println("\n❌ Character edit cancelled.")
	}

	return ui.AskToContinue("Press Enter to continue")
}

// handleDeleteCharacter deletes a character
func handleDeleteCharacter() error {
	characters, err := database.GetAllCharacters()
	if err != nil {
		return fmt.Errorf("failed to get characters: %w", err)
	}

	selected, err := ui.SelectSavedCharacter(characters, "delete")
	if err != nil {
		return err
	}

	confirmed, err := ui.ConfirmDelete(selected.Character.Name)
	if err != nil {
		return fmt.Errorf("failed to confirm deletion: %w", err)
	}

	if confirmed {
		err = database.DeleteCharacter(selected.ID)
		if err != nil {
			return fmt.Errorf("failed to delete character: %w", err)
		}
		fmt.Printf("\n🗑️ Character '%s' deleted successfully.\n", selected.Character.Name)
	} else {
		fmt.Println("\n❌ Character deletion cancelled.")
	}

	return ui.AskToContinue("Press Enter to continue")
}
