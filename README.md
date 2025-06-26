# D&D 5e Character Creator

A comprehensive Dungeons & Dragons 5th Edition character creation tool built with Go, featuring both CLI and web interfaces with full spell selection capabilities.

## 🎲 Features

### Core Character Creation
- **Complete D&D 5e System**: All official races, classes, and backgrounds
- **Ability Score Generation**: Dice roll, point buy, and standard array methods
- **Racial Bonuses**: Automatic application of racial ability score bonuses
- **Character Persistence**: Save and manage multiple characters

### Spell System ✨
- **6 Spellcasting Classes**: Wizard, Cleric, Bard, Sorcerer, Warlock, Druid
- **Interactive Spell Selection**: Beautiful card-based interface
- **Authentic Spell Lists**: Class-appropriate cantrips and 1st level spells
- **Smart Validation**: Enforces proper spell allocation per class

### Web Interface
- **Modern Bootstrap 5 Design**: Responsive and mobile-friendly
- **Interactive Forms**: Dynamic spell selection based on class choice
- **Character Management**: Create, view, and delete characters
- **Beautiful Character Sheets**: Complete stat displays with combat info

## 🚀 Deployment Options

### Option 1: Vercel (Recommended for Production)
Deploy to Vercel with PostgreSQL for a scalable, serverless solution:

```bash
# Install Vercel CLI
npm i -g vercel

# Deploy to Vercel
vercel

# Follow the detailed guide
cat VERCEL_DEPLOYMENT.md
```

**Live Demo**: Your app will be available at `https://your-project.vercel.app`

### Option 2: Local Development
Run locally with SQLite for development and testing:

```bash
# Clone the repository
git clone <your-repo-url>
cd dungeon-and-dragon

# Install dependencies
go mod tidy

# Build the application
go build -o dnd-character-creator

# Run web server
./dnd-character-creator -web -port 8080

# Or run CLI version
./dnd-character-creator
```

## 🛠 Technology Stack

- **Backend**: Go 1.24+ with standard library HTTP server
- **Database**: PostgreSQL (Vercel) or SQLite (local)
- **Frontend**: Bootstrap 5, vanilla JavaScript
- **Deployment**: Vercel serverless functions
- **Styling**: Custom CSS with D&D theming

## 📁 Project Structure

```
/
├── api/                     # Vercel serverless functions
│   ├── index.go            # Home page handler
│   ├── characters/         # Character CRUD operations
│   ├── generate-abilities.go
│   └── spells.go           # Spell data API
├── public/                  # Static assets (Vercel)
│   ├── static/css/         # Stylesheets
│   ├── static/js/          # JavaScript
│   └── templates/          # HTML templates
├── internal/               # Go packages
│   ├── character/          # Character logic & spell system
│   ├── database/           # Database layer
│   └── ui/                 # CLI interface
├── web/                    # Original web handlers (local)
├── vercel.json            # Vercel configuration
└── main.go                # CLI application entry point
```

## 🎮 Usage

### Web Interface
1. **Create Character**: Choose race, class, background, and spells
2. **Spell Selection**: For spellcasters, select appropriate cantrips and spells
3. **View Characters**: Browse your saved character collection
4. **Character Sheets**: View complete character details and stats

### CLI Interface (Local Only)
1. **Interactive Menus**: Navigate with arrow keys and enter
2. **Character Creation**: Step-by-step guided process
3. **Character Management**: View, edit, and delete saved characters
4. **Database Integration**: Persistent character storage

## 🔮 Spell System Details

### Supported Classes
- **Wizard**: 3 cantrips, 2 spells (Intelligence-based)
- **Cleric**: 3 cantrips, 2 spells (Wisdom-based)
- **Bard**: 2 cantrips, 4 spells (Charisma-based)
- **Sorcerer**: 4 cantrips, 2 spells (Charisma-based)
- **Warlock**: 2 cantrips, 2 spells (Charisma-based)
- **Druid**: 2 cantrips, 2 spells (Wisdom-based)

### Spell Features
- **Authentic D&D Spells**: Magic Missile, Fire Bolt, Cure Wounds, etc.
- **School Classification**: Evocation, Transmutation, Enchantment, etc.
- **Interactive Selection**: Click to select/deselect spells
- **Validation**: Prevents over-selection, enforces limits
- **Persistence**: Spells saved with character data

## 🌐 API Endpoints

- `GET /` - Home page
- `GET /characters` - Character list
- `GET /characters/new` - Character creation form
- `POST /characters/new` - Create new character
- `GET /characters/{id}` - View character
- `POST /api/generate-abilities` - Generate ability scores
- `GET /api/spells?class={className}` - Get class spell list

## 🚀 Quick Start (Vercel)

1. **Fork this repository**
2. **Deploy to Vercel**: Connect your GitHub repo
3. **Add PostgreSQL**: Create database in Vercel dashboard
4. **Configure Environment**: Add database connection strings
5. **Deploy**: Push to main branch or use Vercel CLI

Your D&D Character Creator will be live with:
- ✅ Global CDN
- ✅ Automatic HTTPS
- ✅ Serverless scaling
- ✅ PostgreSQL database
- ✅ Full spell system

## 📝 License

This project is for educational and personal use. Dungeons & Dragons content is owned by Wizards of the Coast.

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## 🎯 Roadmap

- [ ] Character view/edit pages for Vercel
- [ ] Spell slot tracking
- [ ] Subclass support with bonus spells
- [ ] Equipment and inventory system
- [ ] Character export (PDF/JSON)
- [ ] Multi-level character progression

---

**Ready to create your next D&D character?** 🗡️✨

Deploy to Vercel for the full experience, or run locally for development!
