package character

// Class represents a D&D 5e class
type Class struct {
	Name           string
	Description    string
	HitDie         int
	PrimaryAbility string
	SavingThrows   []string
	Skills         []string
}

// GetAllClasses returns all available classes
func GetAllClasses() []Class {
	return []Class{
		{
			Name:           "Fighter",
			Description:    "A master of martial combat",
			HitDie:         10,
			PrimaryAbility: "Strength or Dexterity",
			SavingThrows:   []string{"Strength", "Constitution"},
			Skills:         []string{"Acrobatics", "Animal Handling", "Athletics", "History", "Insight", "Intimidation", "Perception", "Survival"},
		},
		{
			Name:           "Wizard",
			Description:    "A scholarly magic-user",
			HitDie:         6,
			PrimaryAbility: "Intelligence",
			SavingThrows:   []string{"Intelligence", "Wisdom"},
			Skills:         []string{"Arcana", "History", "Insight", "Investigation", "Medicine", "Religion"},
		},
		{
			Name:           "Rogue",
			Description:    "A scoundrel who uses stealth and trickery",
			HitDie:         8,
			PrimaryAbility: "Dexterity",
			SavingThrows:   []string{"Dexterity", "Intelligence"},
			Skills:         []string{"Acrobatics", "Athletics", "Deception", "Insight", "Intimidation", "Investigation", "Perception", "Performance", "Persuasion", "Sleight of Hand", "Stealth"},
		},
		{
			Name:           "Cleric",
			Description:    "A priestly champion who wields divine magic",
			HitDie:         8,
			PrimaryAbility: "Wisdom",
			SavingThrows:   []string{"Wisdom", "Charisma"},
			Skills:         []string{"History", "Insight", "Medicine", "Persuasion", "Religion"},
		},
		{
			Name:           "Ranger",
			Description:    "A warrior of the wilderness",
			HitDie:         10,
			PrimaryAbility: "Dexterity and Wisdom",
			SavingThrows:   []string{"Strength", "Dexterity"},
			Skills:         []string{"Animal Handling", "Athletics", "Insight", "Investigation", "Nature", "Perception", "Stealth", "Survival"},
		},
		{
			Name:           "Paladin",
			Description:    "A holy warrior bound to a sacred oath",
			HitDie:         10,
			PrimaryAbility: "Strength and Charisma",
			SavingThrows:   []string{"Wisdom", "Charisma"},
			Skills:         []string{"Athletics", "Insight", "Intimidation", "Medicine", "Persuasion", "Religion"},
		},
		{
			Name:           "Barbarian",
			Description:    "A fierce warrior of primitive background",
			HitDie:         12,
			PrimaryAbility: "Strength",
			SavingThrows:   []string{"Strength", "Constitution"},
			Skills:         []string{"Animal Handling", "Athletics", "Intimidation", "Nature", "Perception", "Survival"},
		},
		{
			Name:           "Bard",
			Description:    "A master of song, speech, and the magic they contain",
			HitDie:         8,
			PrimaryAbility: "Charisma",
			SavingThrows:   []string{"Dexterity", "Charisma"},
			Skills:         []string{"Any three of your choice"},
		},
		{
			Name:           "Druid",
			Description:    "A priest of nature, wielding elemental forces",
			HitDie:         8,
			PrimaryAbility: "Wisdom",
			SavingThrows:   []string{"Intelligence", "Wisdom"},
			Skills:         []string{"Arcana", "Animal Handling", "Insight", "Medicine", "Nature", "Perception", "Religion", "Survival"},
		},
		{
			Name:           "Monk",
			Description:    "A master of martial arts, harnessing inner power",
			HitDie:         8,
			PrimaryAbility: "Dexterity and Wisdom",
			SavingThrows:   []string{"Strength", "Dexterity"},
			Skills:         []string{"Acrobatics", "Athletics", "History", "Insight", "Religion", "Stealth"},
		},
		{
			Name:           "Sorcerer",
			Description:    "A spellcaster who draws on inherent magic",
			HitDie:         6,
			PrimaryAbility: "Charisma",
			SavingThrows:   []string{"Constitution", "Charisma"},
			Skills:         []string{"Arcana", "Deception", "Insight", "Intimidation", "Persuasion", "Religion"},
		},
		{
			Name:           "Warlock",
			Description:    "A wielder of magic derived from a bargain with an extraplanar entity",
			HitDie:         8,
			PrimaryAbility: "Charisma",
			SavingThrows:   []string{"Wisdom", "Charisma"},
			Skills:         []string{"Arcana", "Deception", "History", "Intimidation", "Investigation", "Nature", "Religion"},
		},
	}
}

// GetClassByName returns a class by name
func GetClassByName(name string) (Class, bool) {
	classes := GetAllClasses()
	for _, class := range classes {
		if class.Name == name {
			return class, true
		}
	}
	return Class{}, false
}
