package ui

import (
	"fmt"
	"strconv"

	"github.com/kevinbarbour/dungeon-and-dragon/internal/character"
	"github.com/kevinbarbour/dungeon-and-dragon/internal/database"
	"github.com/manifoldco/promptui"
)

// SelectAbilityGenerationMethod prompts the user to select an ability generation method
func SelectAbilityGenerationMethod() (character.AbilityGenerationMethod, error) {
	methods := character.GetAllMethods()

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F3B2 {{ . | cyan }}",
		Inactive: "  {{ . | white }}",
		Selected: "\U0001F3B2 {{ . | red | cyan }}",
	}

	prompt := promptui.Select{
		Label:     "Choose ability score generation method",
		Items:     methods,
		Templates: templates,
		Size:      len(methods),
		HideHelp:  true,
	}

	prompt.Templates.FuncMap = promptui.FuncMap
	prompt.Templates.FuncMap["methodName"] = character.GetMethodName

	// Convert methods to display strings
	displayItems := make([]string, len(methods))
	for i, method := range methods {
		displayItems[i] = character.GetMethodName(method)
	}

	displayPrompt := promptui.Select{
		Label:     "Choose ability score generation method",
		Items:     displayItems,
		Templates: templates,
		Size:      len(displayItems),
		HideHelp:  true,
	}

	index, _, err := displayPrompt.Run()
	if err != nil {
		return 0, err
	}

	return methods[index], nil
}

// SelectRace prompts the user to select a race
func SelectRace() (character.Race, error) {
	races := character.GetAllRaces()

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F9DD {{ .Name | cyan }} - {{ .Description | faint }}",
		Inactive: "  {{ .Name | white }} - {{ .Description | faint }}",
		Selected: "\U0001F9DD {{ .Name | red | cyan }}",
	}

	prompt := promptui.Select{
		Label:     "Choose your character's race",
		Items:     races,
		Templates: templates,
		Size:      10,
		HideHelp:  true,
	}

	index, _, err := prompt.Run()
	if err != nil {
		return character.Race{}, err
	}

	return races[index], nil
}

// SelectClass prompts the user to select a class
func SelectClass() (character.Class, error) {
	classes := character.GetAllClasses()

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F5E1 {{ .Name | cyan }} - {{ .Description | faint }}",
		Inactive: "  {{ .Name | white }} - {{ .Description | faint }}",
		Selected: "\U0001F5E1 {{ .Name | red | cyan }}",
	}

	prompt := promptui.Select{
		Label:     "Choose your character's class",
		Items:     classes,
		Templates: templates,
		Size:      10,
		HideHelp:  true,
	}

	index, _, err := prompt.Run()
	if err != nil {
		return character.Class{}, err
	}

	return classes[index], nil
}

// SelectBackground prompts the user to select a background
func SelectBackground() (character.Background, error) {
	backgrounds := character.GetAllBackgrounds()

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F4DC {{ .Name | cyan }} - {{ .Description | faint }}",
		Inactive: "  {{ .Name | white }} - {{ .Description | faint }}",
		Selected: "\U0001F4DC {{ .Name | red | cyan }}",
	}

	prompt := promptui.Select{
		Label:     "Choose your character's background",
		Items:     backgrounds,
		Templates: templates,
		Size:      len(backgrounds),
		HideHelp:  true,
	}

	index, _, err := prompt.Run()
	if err != nil {
		return character.Background{}, err
	}

	return backgrounds[index], nil
}

// GetCharacterName prompts the user to enter a character name
func GetCharacterName() (string, error) {
	validate := func(input string) error {
		if len(input) < 1 {
			return fmt.Errorf("character name cannot be empty")
		}
		if len(input) > 50 {
			return fmt.Errorf("character name cannot be longer than 50 characters")
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     "Enter your character's name",
		Templates: templates,
		Validate:  validate,
	}

	return prompt.Run()
}

// ConfirmCharacter asks the user to confirm their character creation
func ConfirmCharacter(char character.Character) (bool, error) {
	fmt.Println("\n" + char.String())

	prompt := promptui.Prompt{
		Label:     "Create this character? (y/N)",
		Default:   "N",
		AllowEdit: true,
	}

	result, err := prompt.Run()
	if err != nil {
		return false, err
	}

	return result == "y" || result == "Y" || result == "yes" || result == "Yes", nil
}

// ShowWelcome displays the welcome message
func ShowWelcome() {
	fmt.Println(`
╔══════════════════════════════════════════════════════════════╗
║                                                              ║
║        🐉 D&D 5th Edition Character Creator 🐉               ║
║                                                              ║
║   Welcome, adventurer! Let's create your new character.     ║
║   Use arrow keys to navigate and Enter to select.           ║
║                                                              ║
╚══════════════════════════════════════════════════════════════╝
`)
}

// ShowAbilityScores displays the generated ability scores
func ShowAbilityScores(abilities character.AbilityScores, method character.AbilityGenerationMethod) {
	fmt.Printf("\n🎲 Generated ability scores using %s:\n\n", character.GetMethodName(method))
	character.PrintAbilityScores(abilities)
	fmt.Println()
}

// AskToContinue prompts the user to continue
func AskToContinue(message string) error {
	prompt := promptui.Prompt{
		Label: message + " (Press Enter to continue)",
	}
	_, err := prompt.Run()
	return err
}

// MainMenuOption represents a main menu option
type MainMenuOption string

const (
	CreateNewCharacter    MainMenuOption = "Create New Character"
	ManageSavedCharacters MainMenuOption = "Manage Saved Characters"
	Exit                  MainMenuOption = "Exit"
)

// ShowMainMenu displays the main menu and returns the selected option
func ShowMainMenu() (MainMenuOption, error) {
	options := []MainMenuOption{
		CreateNewCharacter,
		ManageSavedCharacters,
		Exit,
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "▶ {{ . | cyan }}",
		Inactive: "  {{ . | white }}",
		Selected: "▶ {{ . | red | cyan }}",
	}

	prompt := promptui.Select{
		Label:     "What would you like to do",
		Items:     options,
		Templates: templates,
		Size:      len(options),
		HideHelp:  true,
	}

	index, _, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return options[index], nil
}

// CharacterManagementOption represents character management options
type CharacterManagementOption string

const (
	ListCharacters  CharacterManagementOption = "List All Characters"
	ViewCharacter   CharacterManagementOption = "View Character Details"
	EditCharacter   CharacterManagementOption = "Edit Character"
	DeleteCharacter CharacterManagementOption = "Delete Character"
	BackToMainMenu  CharacterManagementOption = "Back to Main Menu"
)

// ShowCharacterManagementMenu displays the character management menu
func ShowCharacterManagementMenu() (CharacterManagementOption, error) {
	options := []CharacterManagementOption{
		ListCharacters,
		ViewCharacter,
		EditCharacter,
		DeleteCharacter,
		BackToMainMenu,
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "📋 {{ . | cyan }}",
		Inactive: "  {{ . | white }}",
		Selected: "📋 {{ . | red | cyan }}",
	}

	prompt := promptui.Select{
		Label:     "Character Management",
		Items:     options,
		Templates: templates,
		Size:      len(options),
		HideHelp:  true,
	}

	index, _, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return options[index], nil
}

// SelectSavedCharacter allows the user to select from saved characters
func SelectSavedCharacter(characters []database.SavedCharacter, action string) (*database.SavedCharacter, error) {
	if len(characters) == 0 {
		fmt.Println("\n📭 No saved characters found.")
		return nil, fmt.Errorf("no characters available")
	}

	// Create display items with character info
	displayItems := make([]string, len(characters))
	for i, char := range characters {
		displayItems[i] = fmt.Sprintf("%s (Level %d %s %s) - Created: %s",
			char.Character.Name,
			char.Character.Level,
			char.Character.Race.Name,
			char.Character.Class.Name,
			char.CreatedAt.Format("2006-01-02"),
		)
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "🧙 {{ . | cyan }}",
		Inactive: "  {{ . | white }}",
		Selected: "🧙 {{ . | red | cyan }}",
	}

	prompt := promptui.Select{
		Label:     fmt.Sprintf("Select a character to %s", action),
		Items:     displayItems,
		Templates: templates,
		Size:      10,
		HideHelp:  true,
	}

	index, _, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	return &characters[index], nil
}

// ConfirmSaveCharacter asks if the user wants to save the character
func ConfirmSaveCharacter() (bool, error) {
	prompt := promptui.Prompt{
		Label:     "Save this character to database? (Y/n)",
		Default:   "Y",
		AllowEdit: true,
	}

	result, err := prompt.Run()
	if err != nil {
		return false, err
	}

	return result == "Y" || result == "y" || result == "yes" || result == "Yes" || result == "", nil
}

// ConfirmDelete asks for confirmation before deleting a character
func ConfirmDelete(characterName string) (bool, error) {
	prompt := promptui.Prompt{
		Label:     fmt.Sprintf("Are you sure you want to delete '%s'? This cannot be undone (y/N)", characterName),
		Default:   "N",
		AllowEdit: true,
	}

	result, err := prompt.Run()
	if err != nil {
		return false, err
	}

	return result == "y" || result == "Y" || result == "yes" || result == "Yes", nil
}

// ShowCharacterList displays a formatted list of all characters
func ShowCharacterList(characters []database.SavedCharacter) {
	if len(characters) == 0 {
		fmt.Println("\n📭 No saved characters found.")
		return
	}

	fmt.Printf("\n📚 Saved Characters (%d total):\n", len(characters))
	fmt.Println("═══════════════════════════════════════════════════════════════")

	for _, char := range characters {
		fmt.Printf("🧙 %s (ID: %d)\n", char.Character.Name, char.ID)
		fmt.Printf("   Level %d %s %s\n", char.Character.Level, char.Character.Race.Name, char.Character.Class.Name)
		fmt.Printf("   Background: %s\n", char.Character.Background.Name)
		fmt.Printf("   Created: %s\n", char.CreatedAt.Format("2006-01-02 15:04"))
		if !char.UpdatedAt.Equal(char.CreatedAt) {
			fmt.Printf("   Updated: %s\n", char.UpdatedAt.Format("2006-01-02 15:04"))
		}
		fmt.Println("───────────────────────────────────────────────────────────────")
	}
}

// ShowCharacterDetails displays detailed information about a character
func ShowCharacterDetails(saved *database.SavedCharacter) {
	fmt.Printf("\n📋 Character Details (ID: %d)\n", saved.ID)
	fmt.Println("═══════════════════════════════════════════════════════════════")
	fmt.Println(saved.Character.String())
	fmt.Printf("Created: %s\n", saved.CreatedAt.Format("2006-01-02 15:04:05"))
	if !saved.UpdatedAt.Equal(saved.CreatedAt) {
		fmt.Printf("Updated: %s\n", saved.UpdatedAt.Format("2006-01-02 15:04:05"))
	}
}

// GetCharacterID prompts for a character ID
func GetCharacterID() (int, error) {
	validate := func(input string) error {
		if len(input) < 1 {
			return fmt.Errorf("character ID cannot be empty")
		}
		if _, err := strconv.Atoi(input); err != nil {
			return fmt.Errorf("character ID must be a number")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Enter character ID",
		Validate: validate,
	}

	result, err := prompt.Run()
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(result)
}
