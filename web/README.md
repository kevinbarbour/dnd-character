# D&D 5e Character Creator

A pure client-side Dungeons & Dragons 5th Edition character creation tool with full spell selection capabilities. Built with vanilla HTML, CSS, and JavaScript for GitHub Pages deployment.

## 🎲 Features

### Complete D&D 5e System
- **9 Races**: Human, Elf, Dwarf, Halfling, Dragonborn, Gnome, Half-Elf, Half-Orc, Tiefling
- **12 Classes**: Fighter, Wizard, Cleric, Rogue, Ranger, Bard, Barbarian, Sorcerer, Warlock, Paladin, Druid, Monk
- **12 Backgrounds**: Acolyte, Criminal, Folk Hero, Noble, Sage, Soldier, Charlatan, Entertainer, Guild Artisan, Hermit, Outlander, Sailor

### Spell System ✨
- **6 Spellcasting Classes** with authentic D&D spells
- **Interactive spell selection** with beautiful card interface
- **Class-specific spell lists** and limitations
- **Spell validation** to ensure proper allocation

| Class | Cantrips | 1st Level Spells |
|-------|----------|------------------|
| Wizard | 3 | 2 |
| Cleric | 3 | 2 |
| Bard | 2 | 4 |
| Sorcerer | 4 | 2 |
| Warlock | 2 | 2 |
| Druid | 2 | 2 |

### Character Creation
- **Ability Score Generation**: Dice roll (4d6 drop lowest), Point Buy, Standard Array
- **Racial Bonuses**: Automatic application of racial ability score increases
- **Starting Equipment**: Class and background-based equipment
- **Character Validation**: Comprehensive form validation

### Local Storage & Privacy
- **Client-side only** - no server required
- **Local storage** - characters saved in your browser
- **Export/Import** - JSON backup and restore
- **Privacy-first** - data never leaves your device

## 🚀 GitHub Pages Deployment

### Quick Setup
1. **Fork this repository**
2. **Enable GitHub Pages**:
   - Go to repository Settings
   - Scroll to "Pages" section
   - Set source to "Deploy from a branch"
   - Select branch: `main`
   - Select folder: `/web`
3. **Access your site** at `https://yourusername.github.io/repository-name`

### Custom Domain (Optional)
1. Add a `CNAME` file to the `/web` directory with your domain
2. Configure DNS to point to `yourusername.github.io`
3. Enable HTTPS in repository settings

## 📁 Project Structure

```
web/
├── index.html              # Home page
├── characters.html         # Character list (to be created)
├── create.html            # Character creation (to be created)
├── character.html         # Character view (to be created)
├── css/
│   └── style.css          # D&D-themed styling
├── js/
│   ├── app.js             # Main application logic
│   ├── storage.js         # LocalStorage management
│   └── data/
│       ├── races.js       # D&D races data
│       ├── classes.js     # D&D classes data
│       ├── backgrounds.js # D&D backgrounds data
│       └── spells.js      # Complete spell database
└── README.md              # This file
```

## 🎮 Usage

### Creating Characters
1. **Visit the site** and click "Create New Character"
2. **Enter character name** and select race, class, background
3. **Generate ability scores** using your preferred method
4. **Select spells** if your class is a spellcaster
5. **Save character** - stored locally in your browser

### Managing Characters
- **View all characters** on the characters page
- **Export individual characters** as JSON files
- **Import characters** from JSON files
- **Delete characters** with confirmation

### Backup & Restore
- **Export all characters** from the home page
- **Import characters** to merge with existing collection
- **Clear all data** with double confirmation

## 🔧 Technical Details

### Browser Compatibility
- **Modern browsers** (Chrome 60+, Firefox 55+, Safari 12+, Edge 79+)
- **JavaScript ES6+** features used
- **LocalStorage** required (5-10MB limit)
- **No server** or internet connection required after initial load

### Data Storage
Characters are stored in `localStorage` with this structure:
```javascript
{
  id: 1640995200000.123,
  name: "Gandalf the Grey",
  race: "Human",
  class: "Wizard",
  background: "Sage",
  level: 1,
  abilities: {
    strength: 10,
    dexterity: 14,
    constitution: 12,
    intelligence: 16,
    wisdom: 13,
    charisma: 11
  },
  spells: ["Fire Bolt", "Mage Hand", "Minor Illusion", "Magic Missile", "Shield"],
  hitPoints: 7,
  armorClass: 12,
  equipment: [...],
  created: "2024-06-26T22:00:00.000Z",
  updated: "2024-06-26T22:00:00.000Z"
}
```

### Performance
- **Lightweight** - ~200KB total size
- **Fast loading** - static files served from CDN
- **Offline capable** - works without internet after first load
- **Mobile friendly** - responsive Bootstrap 5 design

## 🎨 Customization

### Styling
Edit `css/style.css` to customize:
- Color scheme (CSS custom properties in `:root`)
- Layout and spacing
- Component styling
- Print styles for character sheets

### Data
Modify files in `js/data/` to:
- Add new races, classes, or backgrounds
- Expand spell lists
- Customize starting equipment
- Add homebrew content

### Features
Extend `js/app.js` to add:
- Character leveling system
- Equipment management
- Spell slot tracking
- Combat calculator

## 🐛 Troubleshooting

### Characters Not Saving
- Check if localStorage is enabled in your browser
- Verify you're not in private/incognito mode
- Clear browser cache and try again

### Spells Not Loading
- Ensure JavaScript is enabled
- Check browser console for errors
- Verify all script files are loading correctly

### Export/Import Issues
- Use modern browser with File API support
- Check JSON file format is valid
- Ensure file size is under browser limits

## 📝 License

This project is for educational and personal use. Dungeons & Dragons content is owned by Wizards of the Coast.

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## 🎯 Roadmap

- [ ] Character view/edit pages
- [ ] Spell slot tracking
- [ ] Subclass support
- [ ] Equipment management
- [ ] Character export to PDF
- [ ] Dice roller integration
- [ ] Party management
- [ ] Campaign notes

---

**Ready to create your next D&D character?** 🗡️✨

Visit the live site or deploy your own copy to GitHub Pages!
