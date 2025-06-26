package character

// Race represents a D&D 5e race
type Race struct {
	Name              string
	Description       string
	StrengthBonus     int
	DexterityBonus    int
	ConstitutionBonus int
	IntelligenceBonus int
	WisdomBonus       int
	CharismaBonus     int
	Size              string
	Speed             int
	Traits            []string
}

// GetAllRaces returns all available races
func GetAllRaces() []Race {
	return []Race{
		{
			Name:              "Human",
			Description:       "Versatile and ambitious",
			StrengthBonus:     1,
			DexterityBonus:    1,
			ConstitutionBonus: 1,
			IntelligenceBonus: 1,
			WisdomBonus:       1,
			CharismaBonus:     1,
			Size:              "Medium",
			Speed:             30,
			Traits:            []string{"Extra Language", "Extra Skill"},
		},
		{
			Name:           "Elf",
			Description:    "Graceful and magical",
			DexterityBonus: 2,
			Size:           "Medium",
			Speed:          30,
			Traits:         []string{"Darkvision", "Keen Senses", "Fey Ancestry", "Trance"},
		},
		{
			Name:              "Dwarf",
			Description:       "Hardy and resilient",
			ConstitutionBonus: 2,
			Size:              "Medium",
			Speed:             25,
			Traits:            []string{"Darkvision", "Dwarven Resilience", "Stonecunning"},
		},
		{
			Name:           "Halfling",
			Description:    "Small and lucky",
			DexterityBonus: 2,
			Size:           "Small",
			Speed:          25,
			Traits:         []string{"Lucky", "Brave", "Halfling Nimbleness"},
		},
		{
			Name:          "Dragonborn",
			Description:   "Draconic heritage",
			StrengthBonus: 2,
			CharismaBonus: 1,
			Size:          "Medium",
			Speed:         30,
			Traits:        []string{"Draconic Ancestry", "Breath Weapon", "Damage Resistance"},
		},
		{
			Name:              "Gnome",
			Description:       "Small and clever",
			IntelligenceBonus: 2,
			Size:              "Small",
			Speed:             25,
			Traits:            []string{"Darkvision", "Gnome Cunning"},
		},
		{
			Name:          "Half-Elf",
			Description:   "Between two worlds",
			CharismaBonus: 2,
			Size:          "Medium",
			Speed:         30,
			Traits:        []string{"Darkvision", "Fey Ancestry", "Two Extra Skills"},
		},
		{
			Name:              "Half-Orc",
			Description:       "Fierce and strong",
			StrengthBonus:     2,
			ConstitutionBonus: 1,
			Size:              "Medium",
			Speed:             30,
			Traits:            []string{"Darkvision", "Relentless Endurance", "Savage Attacks"},
		},
		{
			Name:              "Tiefling",
			Description:       "Infernal heritage",
			IntelligenceBonus: 1,
			CharismaBonus:     2,
			Size:              "Medium",
			Speed:             30,
			Traits:            []string{"Darkvision", "Hellish Resistance", "Infernal Legacy"},
		},
	}
}

// GetRaceByName returns a race by name
func GetRaceByName(name string) (Race, bool) {
	races := GetAllRaces()
	for _, race := range races {
		if race.Name == name {
			return race, true
		}
	}
	return Race{}, false
}
