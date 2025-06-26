# D&D 5th Edition Character Creator

A command-line application for creating and managing Dungeons & Dragons 5th Edition characters, built in Go with SQLite database persistence.

## Features

### Character Creation
- **Interactive Menu System**: Navigate with arrow keys and Enter to select
- **Multiple Ability Score Generation Methods**:
  - Dice Roll (4d6 drop lowest)
  - Point Buy (27 points)
  - Standard Array (15,14,13,12,10,8)
- **Complete D&D 5e Content**:
  - 9 Races: Human, Elf, Dwarf, Halfling, Dragonborn, Gnome, Half-Elf, Half-Orc, Tiefling
  - 12 Classes: Fighter, Wizard, Rogue, Cleric, Ranger, Paladin, Barbarian, Bard, Druid, Monk, Sorcerer, Warlock
  - 6 Backgrounds: Acolyte, Criminal, Folk Hero, Noble, Sage, Soldier
- **Automatic Calculations**: Ability modifiers, hit points, armor class
- **Racial Bonuses**: Automatically applied to ability scores

### Database Persistence
- **SQLite Database**: Simple, file-based storage (no external database required)
- **Character Management**:
  - Save characters after creation
  - List all saved characters
  - View detailed character sheets
  - Edit existing characters (race, class, background)
  - Delete characters with confirmation
- **Timestamps**: Track creation and modification dates

### User Interface
- **Beautiful CLI**: Unicode characters and colors for enhanced visual appeal
- **Error Handling**: Comprehensive validation and user-friendly error messages
- **Navigation**: Intuitive menu system with clear instructions

## Installation

### Prerequisites
- Go 1.19 or later
- CGO enabled (required for SQLite)

### Build from Source
```bash
git clone <repository-url>
cd dungeon-and-dragon
go mod download
go build -o dnd-character-creator
```

## Usage

Run the application:
```bash
./dnd-character-creator
```

### Main Menu Options
1. **Create New Character**: Step-by-step character creation process
2. **Manage Saved Characters**: View, edit, or delete existing characters
3. **Exit**: Close the application

### Character Creation Flow
1. Enter character name
2. Choose ability score generation method
3. Generate and review ability scores
4. Select race (racial bonuses automatically applied)
5. Select class
6. Select background
7. Review and confirm character
8. Choose to save to database

### Character Management
- **List All Characters**: See all saved characters with basic info
- **View Character Details**: Display complete character sheet
- **Edit Character**: Modify race, class, and background
- **Delete Character**: Remove character with confirmation

## Database

The application creates a SQLite database file (`dnd_characters.db`) in the same directory as the executable. This file contains:

- Character data (name, race, class, background, level, ability scores)
- Creation and modification timestamps
- Automatic ID assignment

## Project Structure

```
dungeon-and-dragon/
├── main.go                           # Application entry point and main logic
├── go.mod                           # Go module dependencies
├── dnd-character-creator            # Compiled executable
├── dnd_characters.db                # SQLite database (created on first run)
└── internal/
    ├── character/                   # Character data structures and logic
    │   ├── character.go             # Core character struct and methods
    │   ├── race.go                  # Race definitions and bonuses
    │   ├── class.go                 # Class definitions and features
    │   ├── background.go            # Background definitions
    │   └── abilities.go             # Ability score generation methods
    ├── database/                    # Database layer
    │   ├── database.go              # SQLite connection and initialization
    │   └── character_repository.go  # CRUD operations for characters
    └── ui/                          # User interface
        └── menu.go                  # Interactive menu system and prompts
```

## Dependencies

- **github.com/manifoldco/promptui**: Interactive CLI prompts and menus
- **github.com/mattn/go-sqlite3**: SQLite database driver

## D&D 5e Rules Implementation

The application implements core D&D 5th Edition rules:

- **Ability Scores**: Six core abilities (STR, DEX, CON, INT, WIS, CHA)
- **Ability Modifiers**: Calculated as (score - 10) / 2
- **Racial Bonuses**: Applied automatically based on selected race
- **Hit Points**: Class hit die + Constitution modifier
- **Armor Class**: Base 10 + Dexterity modifier

## Examples

### Creating a Character
```
🐉 D&D 5th Edition Character Creator 🐉

What would you like to do?
▶ Create New Character
  Manage Saved Characters
  Exit

Enter your character's name: Thorin

Choose ability score generation method?
🎲 Dice Roll (4d6 drop lowest)
  Point Buy (27 points)
  Standard Array (15,14,13,12,10,8)

Generated ability scores using Dice Roll:
Strength:     16 (+3)
Dexterity:    12 (+1)
Constitution: 15 (+2)
Intelligence: 10 (+0)
Wisdom:       13 (+1)
Charisma:     8 (-1)

Choose your character's race?
🧝 Dwarf - Hardy and resilient

Applied Dwarf racial bonuses to ability scores!
Constitution: 17 (+3)

Choose your character's class?
🗡 Fighter - A master of martial combat

Choose your character's background?
📜 Soldier - War has been your life

=== CHARACTER SHEET ===
Name: Thorin
Race: Dwarf
Class: Fighter
Background: Soldier
Level: 1

ABILITY SCORES:
Strength:     16 (+3)
Dexterity:    12 (+1)
Constitution: 17 (+3)
Intelligence: 10 (+0)
Wisdom:       13 (+1)
Charisma:     8 (-1)

Hit Points: 13
Armor Class: 11

Save this character to database? (Y/n): Y
💾 Character saved successfully with ID: 1
```

## License

This project is for educational and personal use. D&D 5th Edition content is owned by Wizards of the Coast.
