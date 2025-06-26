package ui

import (
	"fmt"

	"github.com/kevinbarbour/dungeon-and-dragon/internal/character"
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
