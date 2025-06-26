package character

import (
	"fmt"
	"math/rand"
	"time"
)

// AbilityGenerationMethod represents different ways to generate ability scores
type AbilityGenerationMethod int

const (
	DiceRoll AbilityGenerationMethod = iota
	PointBuy
	StandardArray
)

// GenerateAbilityScores generates ability scores based on the chosen method
func GenerateAbilityScores(method AbilityGenerationMethod) AbilityScores {
	switch method {
	case DiceRoll:
		return generateDiceRollAbilities()
	case PointBuy:
		return generatePointBuyAbilities()
	case StandardArray:
		return generateStandardArrayAbilities()
	default:
		return generateStandardArrayAbilities()
	}
}

// generateDiceRollAbilities generates abilities using 4d6 drop lowest
func generateDiceRollAbilities() AbilityScores {
	rand.Seed(time.Now().UnixNano())

	rollStat := func() int {
		rolls := make([]int, 4)
		for i := 0; i < 4; i++ {
			rolls[i] = rand.Intn(6) + 1
		}

		// Find and remove the lowest roll
		lowest := 0
		for i := 1; i < 4; i++ {
			if rolls[i] < rolls[lowest] {
				lowest = i
			}
		}

		total := 0
		for i := 0; i < 4; i++ {
			if i != lowest {
				total += rolls[i]
			}
		}
		return total
	}

	return AbilityScores{
		Strength:     rollStat(),
		Dexterity:    rollStat(),
		Constitution: rollStat(),
		Intelligence: rollStat(),
		Wisdom:       rollStat(),
		Charisma:     rollStat(),
	}
}

// generatePointBuyAbilities generates abilities using point buy system (simplified)
func generatePointBuyAbilities() AbilityScores {
	// For simplicity, we'll return a balanced spread
	// In a full implementation, this would be interactive
	return AbilityScores{
		Strength:     13,
		Dexterity:    14,
		Constitution: 15,
		Intelligence: 12,
		Wisdom:       10,
		Charisma:     8,
	}
}

// generateStandardArrayAbilities generates abilities using the standard array
func generateStandardArrayAbilities() AbilityScores {
	return AbilityScores{
		Strength:     15,
		Dexterity:    14,
		Constitution: 13,
		Intelligence: 12,
		Wisdom:       10,
		Charisma:     8,
	}
}

// GetMethodName returns the name of the ability generation method
func GetMethodName(method AbilityGenerationMethod) string {
	switch method {
	case DiceRoll:
		return "Dice Roll (4d6 drop lowest)"
	case PointBuy:
		return "Point Buy (27 points)"
	case StandardArray:
		return "Standard Array (15,14,13,12,10,8)"
	default:
		return "Unknown"
	}
}

// GetAllMethods returns all available ability generation methods
func GetAllMethods() []AbilityGenerationMethod {
	return []AbilityGenerationMethod{DiceRoll, PointBuy, StandardArray}
}

// PrintAbilityScores prints the ability scores in a formatted way
func PrintAbilityScores(abilities AbilityScores) {
	fmt.Printf("Strength:     %d (%+d)\n", abilities.Strength, abilities.GetStrengthModifier())
	fmt.Printf("Dexterity:    %d (%+d)\n", abilities.Dexterity, abilities.GetDexterityModifier())
	fmt.Printf("Constitution: %d (%+d)\n", abilities.Constitution, abilities.GetConstitutionModifier())
	fmt.Printf("Intelligence: %d (%+d)\n", abilities.Intelligence, abilities.GetIntelligenceModifier())
	fmt.Printf("Wisdom:       %d (%+d)\n", abilities.Wisdom, abilities.GetWisdomModifier())
	fmt.Printf("Charisma:     %d (%+d)\n", abilities.Charisma, abilities.GetCharismaModifier())
}
