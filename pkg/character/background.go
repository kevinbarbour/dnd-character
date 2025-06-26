package character

// Background represents a D&D 5e background
type Background struct {
	Name        string
	Description string
	Skills      []string
	Languages   []string
	Equipment   []string
	Feature     string
}

// GetAllBackgrounds returns all available backgrounds
func GetAllBackgrounds() []Background {
	return []Background{
		{
			Name:        "Acolyte",
			Description: "You have spent your life in the service of a temple",
			Skills:      []string{"Insight", "Religion"},
			Languages:   []string{"Two of your choice"},
			Equipment:   []string{"Holy symbol", "Prayer book", "Incense", "Vestments", "Common clothes", "Belt pouch with 15 gp"},
			Feature:     "Shelter of the Faithful",
		},
		{
			Name:        "Criminal",
			Description: "You are an experienced criminal with a history of breaking the law",
			Skills:      []string{"Deception", "Stealth"},
			Languages:   []string{},
			Equipment:   []string{"Crowbar", "Dark common clothes with hood", "Belt pouch with 15 gp"},
			Feature:     "Criminal Contact",
		},
		{
			Name:        "Folk Hero",
			Description: "You come from a humble social rank, but you are destined for so much more",
			Skills:      []string{"Animal Handling", "Survival"},
			Languages:   []string{},
			Equipment:   []string{"Artisan's tools", "Shovel", "Iron pot", "Common clothes", "Belt pouch with 10 gp"},
			Feature:     "Rustic Hospitality",
		},
		{
			Name:        "Noble",
			Description: "You understand wealth, power, and privilege",
			Skills:      []string{"History", "Persuasion"},
			Languages:   []string{"One of your choice"},
			Equipment:   []string{"Fine clothes", "Signet ring", "Scroll of pedigree", "Purse with 25 gp"},
			Feature:     "Position of Privilege",
		},
		{
			Name:        "Sage",
			Description: "You spent years learning the lore of the multiverse",
			Skills:      []string{"Arcana", "History"},
			Languages:   []string{"Two of your choice"},
			Equipment:   []string{"Bottle of black ink", "Quill", "Small knife", "Letter from colleague", "Common clothes", "Belt pouch with 10 gp"},
			Feature:     "Researcher",
		},
		{
			Name:        "Soldier",
			Description: "War has been your life for as long as you care to remember",
			Skills:      []string{"Athletics", "Intimidation"},
			Languages:   []string{},
			Equipment:   []string{"Insignia of rank", "Trophy from fallen enemy", "Deck of cards", "Common clothes", "Belt pouch with 10 gp"},
			Feature:     "Military Rank",
		},
	}
}

// GetBackgroundByName returns a background by name
func GetBackgroundByName(name string) (Background, bool) {
	backgrounds := GetAllBackgrounds()
	for _, background := range backgrounds {
		if background.Name == name {
			return background, true
		}
	}
	return Background{}, false
}
