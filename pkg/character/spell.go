package character

// Spell represents a D&D 5e spell
type Spell struct {
	Name        string
	Level       int
	School      string
	Description string
}

// SpellList represents spells available to a class
type SpellList struct {
	Cantrips    []Spell
	FirstLevel  []Spell
	BonusSpells []Spell // For subclass bonus spells
}

// GetSpellsForClass returns the spell list for a given class
func GetSpellsForClass(className string) SpellList {
	switch className {
	case "Wizard":
		return SpellList{
			Cantrips: []Spell{
				{Name: "Mage Hand", Level: 0, School: "Transmutation", Description: "Create a spectral hand that can manipulate objects at a distance."},
				{Name: "Prestidigitation", Level: 0, School: "Transmutation", Description: "Perform minor magical tricks and effects."},
				{Name: "Fire Bolt", Level: 0, School: "Evocation", Description: "Hurl a mote of fire at a creature or object."},
				{Name: "Minor Illusion", Level: 0, School: "Illusion", Description: "Create a sound or image that lasts for 1 minute."},
				{Name: "Light", Level: 0, School: "Evocation", Description: "Touch an object to make it shed bright light."},
				{Name: "Ray of Frost", Level: 0, School: "Evocation", Description: "Frigid beam that damages and slows a target."},
			},
			FirstLevel: []Spell{
				{Name: "Magic Missile", Level: 1, School: "Evocation", Description: "Create three darts of magical force that automatically hit their targets."},
				{Name: "Shield", Level: 1, School: "Abjuration", Description: "Invisible barrier gives +5 AC and immunity to magic missile."},
				{Name: "Detect Magic", Level: 1, School: "Divination", Description: "Sense the presence of magic within 30 feet."},
				{Name: "Identify", Level: 1, School: "Divination", Description: "Learn the properties of a magical item or spell."},
				{Name: "Sleep", Level: 1, School: "Enchantment", Description: "Put creatures to sleep in a 20-foot radius."},
				{Name: "Burning Hands", Level: 1, School: "Evocation", Description: "Flames sweep out from your hands in a 15-foot cone."},
			},
		}
	case "Cleric":
		return SpellList{
			Cantrips: []Spell{
				{Name: "Sacred Flame", Level: 0, School: "Evocation", Description: "Flame strikes a target with radiant damage."},
				{Name: "Guidance", Level: 0, School: "Divination", Description: "Touch a creature to give it a bonus to one ability check."},
				{Name: "Light", Level: 0, School: "Evocation", Description: "Touch an object to make it shed bright light."},
				{Name: "Thaumaturgy", Level: 0, School: "Transmutation", Description: "Manifest minor wonders that show divine power."},
				{Name: "Spare the Dying", Level: 0, School: "Necromancy", Description: "Stabilize a dying creature."},
				{Name: "Resistance", Level: 0, School: "Abjuration", Description: "Touch a creature to give it resistance to one damage type."},
			},
			FirstLevel: []Spell{
				{Name: "Cure Wounds", Level: 1, School: "Evocation", Description: "Heal a creature by touching it."},
				{Name: "Bless", Level: 1, School: "Enchantment", Description: "Up to three creatures gain bonuses to attack rolls and saving throws."},
				{Name: "Healing Word", Level: 1, School: "Evocation", Description: "Heal a creature at range as a bonus action."},
				{Name: "Guiding Bolt", Level: 1, School: "Evocation", Description: "Flash of light deals radiant damage and marks the target."},
				{Name: "Shield of Faith", Level: 1, School: "Abjuration", Description: "Give a creature +2 AC for 10 minutes."},
				{Name: "Command", Level: 1, School: "Enchantment", Description: "Force a creature to follow a one-word command."},
			},
		}
	case "Bard":
		return SpellList{
			Cantrips: []Spell{
				{Name: "Vicious Mockery", Level: 0, School: "Enchantment", Description: "Insult a creature to deal psychic damage and impose disadvantage."},
				{Name: "Minor Illusion", Level: 0, School: "Illusion", Description: "Create a sound or image that lasts for 1 minute."},
				{Name: "Prestidigitation", Level: 0, School: "Transmutation", Description: "Perform minor magical tricks and effects."},
				{Name: "Mage Hand", Level: 0, School: "Transmutation", Description: "Create a spectral hand that can manipulate objects at a distance."},
				{Name: "Thunderclap", Level: 0, School: "Evocation", Description: "Create a burst of thunderous sound around you."},
				{Name: "Dancing Lights", Level: 0, School: "Evocation", Description: "Create up to four lights that you can move around."},
			},
			FirstLevel: []Spell{
				{Name: "Healing Word", Level: 1, School: "Evocation", Description: "Heal a creature at range as a bonus action."},
				{Name: "Thunderwave", Level: 1, School: "Evocation", Description: "Thunder damages and pushes creatures in a 15-foot cube."},
				{Name: "Charm Person", Level: 1, School: "Enchantment", Description: "Make a humanoid regard you as a friendly acquaintance."},
				{Name: "Dissonant Whispers", Level: 1, School: "Enchantment", Description: "Whisper that deals psychic damage and forces movement."},
				{Name: "Faerie Fire", Level: 1, School: "Evocation", Description: "Outline creatures in light, giving advantage on attacks against them."},
				{Name: "Sleep", Level: 1, School: "Enchantment", Description: "Put creatures to sleep in a 20-foot radius."},
			},
		}
	case "Sorcerer":
		return SpellList{
			Cantrips: []Spell{
				{Name: "Fire Bolt", Level: 0, School: "Evocation", Description: "Hurl a mote of fire at a creature or object."},
				{Name: "Ray of Frost", Level: 0, School: "Evocation", Description: "Frigid beam that damages and slows a target."},
				{Name: "Mage Hand", Level: 0, School: "Transmutation", Description: "Create a spectral hand that can manipulate objects at a distance."},
				{Name: "Prestidigitation", Level: 0, School: "Transmutation", Description: "Perform minor magical tricks and effects."},
				{Name: "Light", Level: 0, School: "Evocation", Description: "Touch an object to make it shed bright light."},
				{Name: "Minor Illusion", Level: 0, School: "Illusion", Description: "Create a sound or image that lasts for 1 minute."},
			},
			FirstLevel: []Spell{
				{Name: "Magic Missile", Level: 1, School: "Evocation", Description: "Create three darts of magical force that automatically hit their targets."},
				{Name: "Shield", Level: 1, School: "Abjuration", Description: "Invisible barrier gives +5 AC and immunity to magic missile."},
				{Name: "Burning Hands", Level: 1, School: "Evocation", Description: "Flames sweep out from your hands in a 15-foot cone."},
				{Name: "Charm Person", Level: 1, School: "Enchantment", Description: "Make a humanoid regard you as a friendly acquaintance."},
				{Name: "Sleep", Level: 1, School: "Enchantment", Description: "Put creatures to sleep in a 20-foot radius."},
				{Name: "Thunderwave", Level: 1, School: "Evocation", Description: "Thunder damages and pushes creatures in a 15-foot cube."},
			},
		}
	case "Warlock":
		return SpellList{
			Cantrips: []Spell{
				{Name: "Eldritch Blast", Level: 0, School: "Evocation", Description: "Crackling beam of energy that grows stronger as you level."},
				{Name: "Mage Hand", Level: 0, School: "Transmutation", Description: "Create a spectral hand that can manipulate objects at a distance."},
				{Name: "Minor Illusion", Level: 0, School: "Illusion", Description: "Create a sound or image that lasts for 1 minute."},
				{Name: "Prestidigitation", Level: 0, School: "Transmutation", Description: "Perform minor magical tricks and effects."},
				{Name: "Chill Touch", Level: 0, School: "Necromancy", Description: "Ghostly hand deals necrotic damage and prevents healing."},
				{Name: "Thaumaturgy", Level: 0, School: "Transmutation", Description: "Manifest minor wonders that show supernatural power."},
			},
			FirstLevel: []Spell{
				{Name: "Hex", Level: 1, School: "Enchantment", Description: "Curse a target to deal extra damage and impose disadvantage."},
				{Name: "Armor of Agathys", Level: 1, School: "Abjuration", Description: "Gain temporary hit points and damage attackers with cold."},
				{Name: "Charm Person", Level: 1, School: "Enchantment", Description: "Make a humanoid regard you as a friendly acquaintance."},
				{Name: "Arms of Hadar", Level: 1, School: "Conjuration", Description: "Tentacles burst from the ground to restrain creatures."},
				{Name: "Burning Hands", Level: 1, School: "Evocation", Description: "Flames sweep out from your hands in a 15-foot cone."},
				{Name: "Comprehend Languages", Level: 1, School: "Divination", Description: "Understand any spoken or written language."},
			},
		}
	case "Druid":
		return SpellList{
			Cantrips: []Spell{
				{Name: "Druidcraft", Level: 0, School: "Transmutation", Description: "Create minor nature effects like blooming flowers or weather prediction."},
				{Name: "Guidance", Level: 0, School: "Divination", Description: "Touch a creature to give it a bonus to one ability check."},
				{Name: "Produce Flame", Level: 0, School: "Conjuration", Description: "Create a flame in your hand that can be thrown or used for light."},
				{Name: "Resistance", Level: 0, School: "Abjuration", Description: "Touch a creature to give it resistance to one damage type."},
				{Name: "Shillelagh", Level: 0, School: "Transmutation", Description: "Make a club or quarterstaff magical and use Wisdom for attacks."},
				{Name: "Thorn Whip", Level: 0, School: "Transmutation", Description: "Create a whip of thorns that pulls creatures toward you."},
			},
			FirstLevel: []Spell{
				{Name: "Cure Wounds", Level: 1, School: "Evocation", Description: "Heal a creature by touching it."},
				{Name: "Entangle", Level: 1, School: "Conjuration", Description: "Grasping weeds and vines restrain creatures in a 20-foot square."},
				{Name: "Faerie Fire", Level: 1, School: "Evocation", Description: "Outline creatures in light, giving advantage on attacks against them."},
				{Name: "Goodberry", Level: 1, School: "Transmutation", Description: "Create berries that provide nourishment and healing."},
				{Name: "Speak with Animals", Level: 1, School: "Divination", Description: "Communicate with beasts for 10 minutes."},
				{Name: "Thunderwave", Level: 1, School: "Evocation", Description: "Thunder damages and pushes creatures in a 15-foot cube."},
			},
		}
	default:
		return SpellList{} // Non-spellcasting classes
	}
}

// IsSpellcaster returns true if the class can cast spells at level 1
func IsSpellcaster(className string) bool {
	spellcasters := []string{"Wizard", "Cleric", "Bard", "Sorcerer", "Warlock", "Druid"}
	for _, class := range spellcasters {
		if class == className {
			return true
		}
	}
	return false
}

// GetSpellAllocation returns how many cantrips and spells a level 1 character gets
func GetSpellAllocation(className string) (cantrips int, spells int) {
	switch className {
	case "Wizard":
		return 3, 2 // 3 cantrips, 2 first-level spells
	case "Cleric":
		return 3, 2
	case "Bard":
		return 2, 4 // Bards know more spells
	case "Sorcerer":
		return 4, 2 // Sorcerers get more cantrips
	case "Warlock":
		return 2, 2 // Warlocks have fewer spell slots but regain on short rest
	case "Druid":
		return 2, 2
	default:
		return 0, 0 // Non-spellcasters
	}
}
