package character

import "fmt"

// Character represents a D&D 5e character
type Character struct {
	Name       string
	Race       Race
	Class      Class
	Background Background
	Abilities  AbilityScores
	Level      int
	Spells     []string // List of spell names the character knows
}

// AbilityScores represents the six core ability scores
type AbilityScores struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
}

// GetModifier calculates the ability modifier for a given score
func (a *AbilityScores) GetModifier(score int) int {
	return (score - 10) / 2
}

// GetStrengthModifier returns the strength modifier
func (a *AbilityScores) GetStrengthModifier() int {
	return a.GetModifier(a.Strength)
}

// GetDexterityModifier returns the dexterity modifier
func (a *AbilityScores) GetDexterityModifier() int {
	return a.GetModifier(a.Dexterity)
}

// GetConstitutionModifier returns the constitution modifier
func (a *AbilityScores) GetConstitutionModifier() int {
	return a.GetModifier(a.Constitution)
}

// GetIntelligenceModifier returns the intelligence modifier
func (a *AbilityScores) GetIntelligenceModifier() int {
	return a.GetModifier(a.Intelligence)
}

// GetWisdomModifier returns the wisdom modifier
func (a *AbilityScores) GetWisdomModifier() int {
	return a.GetModifier(a.Wisdom)
}

// GetCharismaModifier returns the charisma modifier
func (a *AbilityScores) GetCharismaModifier() int {
	return a.GetModifier(a.Charisma)
}

// ApplyRacialBonuses applies racial ability score bonuses
func (a *AbilityScores) ApplyRacialBonuses(race Race) {
	a.Strength += race.StrengthBonus
	a.Dexterity += race.DexterityBonus
	a.Constitution += race.ConstitutionBonus
	a.Intelligence += race.IntelligenceBonus
	a.Wisdom += race.WisdomBonus
	a.Charisma += race.CharismaBonus
}

// String returns a formatted string representation of the character
func (c *Character) String() string {
	return fmt.Sprintf(`
=== CHARACTER SHEET ===
Name: %s
Race: %s
Class: %s
Background: %s
Level: %d

ABILITY SCORES:
Strength:     %d (%+d)
Dexterity:    %d (%+d)
Constitution: %d (%+d)
Intelligence: %d (%+d)
Wisdom:       %d (%+d)
Charisma:     %d (%+d)

Hit Points: %d
Armor Class: %d (base 10 + Dex modifier)
`,
		c.Name,
		c.Race.Name,
		c.Class.Name,
		c.Background.Name,
		c.Level,
		c.Abilities.Strength, c.Abilities.GetStrengthModifier(),
		c.Abilities.Dexterity, c.Abilities.GetDexterityModifier(),
		c.Abilities.Constitution, c.Abilities.GetConstitutionModifier(),
		c.Abilities.Intelligence, c.Abilities.GetIntelligenceModifier(),
		c.Abilities.Wisdom, c.Abilities.GetWisdomModifier(),
		c.Abilities.Charisma, c.Abilities.GetCharismaModifier(),
		c.GetHitPoints(),
		10+c.Abilities.GetDexterityModifier(),
	)
}

// GetHitPoints calculates the character's hit points
func (c *Character) GetHitPoints() int {
	return c.Class.HitDie + c.Abilities.GetConstitutionModifier()
}
