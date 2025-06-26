// D&D 5e Spells Data
const SPELLS = {
    Wizard: {
        cantrips: [
            {
                name: "Fire Bolt",
                school: "Evocation",
                description: "You hurl a mote of fire at a creature or object within range. Make a ranged spell attack against the target. On a hit, the target takes 1d10 fire damage."
            },
            {
                name: "Mage Hand",
                school: "Transmutation",
                description: "A spectral, floating hand appears at a point you choose within range. The hand lasts for the duration or until you dismiss it as an action."
            },
            {
                name: "Minor Illusion",
                school: "Illusion",
                description: "You create a sound or an image of an object within range that lasts for the duration. The illusion also ends if you dismiss it as an action or cast this spell again."
            },
            {
                name: "Prestidigitation",
                school: "Transmutation",
                description: "This spell is a minor magical trick that novice spellcasters use for practice. You create one of several minor effects within range."
            },
            {
                name: "Ray of Frost",
                school: "Evocation",
                description: "A frigid beam of blue-white light streaks toward a creature within range. Make a ranged spell attack against the target."
            },
            {
                name: "Light",
                school: "Evocation",
                description: "You touch one object that is no larger than 10 feet in any dimension. Until the spell ends, the object sheds bright light in a 20-foot radius."
            }
        ],
        firstLevel: [
            {
                name: "Magic Missile",
                school: "Evocation",
                description: "You create three glowing darts of magical force. Each dart hits a creature of your choice that you can see within range."
            },
            {
                name: "Shield",
                school: "Abjuration",
                description: "An invisible barrier of magical force appears and protects you. Until the start of your next turn, you have a +5 bonus to AC."
            },
            {
                name: "Burning Hands",
                school: "Evocation",
                description: "As you hold your hands with thumbs touching and fingers spread, a thin sheet of flames shoots forth from your outstretched fingertips."
            },
            {
                name: "Detect Magic",
                school: "Divination",
                description: "For the duration, you sense the presence of magic within 30 feet of you."
            },
            {
                name: "Sleep",
                school: "Enchantment",
                description: "This spell sends creatures into a magical slumber. Roll 5d8; the total is how many hit points of creatures this spell can affect."
            }
        ]
    },
    Cleric: {
        cantrips: [
            {
                name: "Sacred Flame",
                school: "Evocation",
                description: "Flame-like radiance descends on a creature that you can see within range. The target must succeed on a Dexterity saving throw or take 1d8 radiant damage."
            },
            {
                name: "Guidance",
                school: "Divination",
                description: "You touch one willing creature. Once before the spell ends, the target can roll a d4 and add the number rolled to one ability check of its choice."
            },
            {
                name: "Light",
                school: "Evocation",
                description: "You touch one object that is no larger than 10 feet in any dimension. Until the spell ends, the object sheds bright light in a 20-foot radius."
            },
            {
                name: "Thaumaturgy",
                school: "Transmutation",
                description: "You manifest a minor wonder, a sign of supernatural power, within range. You create one of several effects within range."
            },
            {
                name: "Spare the Dying",
                school: "Necromancy",
                description: "You touch a living creature that has 0 hit points. The creature becomes stable."
            },
            {
                name: "Resistance",
                school: "Abjuration",
                description: "You touch one willing creature. Once before the spell ends, the target can roll a d4 and add the number rolled to one saving throw of its choice."
            }
        ],
        firstLevel: [
            {
                name: "Cure Wounds",
                school: "Evocation",
                description: "A creature you touch regains a number of hit points equal to 1d8 + your spellcasting ability modifier."
            },
            {
                name: "Healing Word",
                school: "Evocation",
                description: "A creature of your choice that you can see within range regains hit points equal to 1d4 + your spellcasting ability modifier."
            },
            {
                name: "Bless",
                school: "Enchantment",
                description: "You bless up to three creatures of your choice within range. Whenever a target makes an attack roll or a saving throw before the spell ends, the target can roll a d4 and add the number rolled to the attack roll or saving throw."
            },
            {
                name: "Command",
                school: "Enchantment",
                description: "You speak a one-word command to a creature you can see within range. The target must succeed on a Wisdom saving throw or follow the command on its next turn."
            },
            {
                name: "Detect Evil and Good",
                school: "Divination",
                description: "For the duration, you know if there is an aberration, celestial, elemental, fey, fiend, or undead within 30 feet of you."
            }
        ]
    },
    Bard: {
        cantrips: [
            {
                name: "Vicious Mockery",
                school: "Enchantment",
                description: "You unleash a string of insults laced with subtle enchantments at a creature you can see within range."
            },
            {
                name: "Minor Illusion",
                school: "Illusion",
                description: "You create a sound or an image of an object within range that lasts for the duration."
            },
            {
                name: "Prestidigitation",
                school: "Transmutation",
                description: "This spell is a minor magical trick that novice spellcasters use for practice."
            },
            {
                name: "Mage Hand",
                school: "Transmutation",
                description: "A spectral, floating hand appears at a point you choose within range."
            }
        ],
        firstLevel: [
            {
                name: "Healing Word",
                school: "Evocation",
                description: "A creature of your choice that you can see within range regains hit points equal to 1d4 + your spellcasting ability modifier."
            },
            {
                name: "Thunderwave",
                school: "Evocation",
                description: "A wave of thunderous force sweeps out from you. Each creature in a 15-foot cube originating from you must make a Constitution saving throw."
            },
            {
                name: "Charm Person",
                school: "Enchantment",
                description: "You attempt to charm a humanoid you can see within range. It must make a Wisdom saving throw, and it does so with advantage if you or your companions are fighting it."
            },
            {
                name: "Dissonant Whispers",
                school: "Enchantment",
                description: "You whisper a discordant melody that only one creature of your choice within range can hear, wracking it with terrible pain."
            },
            {
                name: "Faerie Fire",
                school: "Evocation",
                description: "Each object in a 20-foot cube within range is outlined in blue, green, or violet light (your choice)."
            },
            {
                name: "Identify",
                school: "Divination",
                description: "You choose one object that you must touch throughout the casting of the spell. If it is a magic item or some other magic-imbued object, you learn its properties and how to use them."
            }
        ]
    },
    Sorcerer: {
        cantrips: [
            {
                name: "Fire Bolt",
                school: "Evocation",
                description: "You hurl a mote of fire at a creature or object within range. Make a ranged spell attack against the target."
            },
            {
                name: "Ray of Frost",
                school: "Evocation",
                description: "A frigid beam of blue-white light streaks toward a creature within range."
            },
            {
                name: "Shocking Grasp",
                school: "Evocation",
                description: "Lightning springs from your hand to deliver a shock to a creature you try to touch."
            },
            {
                name: "Mage Hand",
                school: "Transmutation",
                description: "A spectral, floating hand appears at a point you choose within range."
            },
            {
                name: "Minor Illusion",
                school: "Illusion",
                description: "You create a sound or an image of an object within range that lasts for the duration."
            },
            {
                name: "Prestidigitation",
                school: "Transmutation",
                description: "This spell is a minor magical trick that novice spellcasters use for practice."
            }
        ],
        firstLevel: [
            {
                name: "Magic Missile",
                school: "Evocation",
                description: "You create three glowing darts of magical force. Each dart hits a creature of your choice that you can see within range."
            },
            {
                name: "Shield",
                school: "Abjuration",
                description: "An invisible barrier of magical force appears and protects you."
            },
            {
                name: "Burning Hands",
                school: "Evocation",
                description: "A thin sheet of flames shoots forth from your outstretched fingertips."
            },
            {
                name: "Charm Person",
                school: "Enchantment",
                description: "You attempt to charm a humanoid you can see within range."
            },
            {
                name: "Color Spray",
                school: "Illusion",
                description: "A dazzling array of flashing, colored light springs from your hand."
            }
        ]
    },
    Warlock: {
        cantrips: [
            {
                name: "Eldritch Blast",
                school: "Evocation",
                description: "A beam of crackling energy streaks toward a creature within range. Make a ranged spell attack against the target."
            },
            {
                name: "Minor Illusion",
                school: "Illusion",
                description: "You create a sound or an image of an object within range that lasts for the duration."
            },
            {
                name: "Prestidigitation",
                school: "Transmutation",
                description: "This spell is a minor magical trick that novice spellcasters use for practice."
            },
            {
                name: "Mage Hand",
                school: "Transmutation",
                description: "A spectral, floating hand appears at a point you choose within range."
            }
        ],
        firstLevel: [
            {
                name: "Hex",
                school: "Enchantment",
                description: "You place a curse on a creature that you can see within range. Until the spell ends, you deal an extra 1d6 necrotic damage to the target whenever you hit it with an attack."
            },
            {
                name: "Arms of Hadar",
                school: "Conjuration",
                description: "You invoke the power of Hadar, the Dark Hunger. Tendrils of dark energy erupt from you and batter all creatures within 10 feet of you."
            },
            {
                name: "Charm Person",
                school: "Enchantment",
                description: "You attempt to charm a humanoid you can see within range."
            },
            {
                name: "Hellish Rebuke",
                school: "Evocation",
                description: "You point your finger, and the creature that damaged you is momentarily surrounded by hellish flames."
            },
            {
                name: "Protection from Evil and Good",
                school: "Abjuration",
                description: "Until the spell ends, one willing creature you touch is protected against certain types of creatures."
            }
        ]
    },
    Druid: {
        cantrips: [
            {
                name: "Druidcraft",
                school: "Transmutation",
                description: "Whispering to the spirits of nature, you create one of several minor effects within range."
            },
            {
                name: "Guidance",
                school: "Divination",
                description: "You touch one willing creature. Once before the spell ends, the target can roll a d4 and add the number rolled to one ability check of its choice."
            },
            {
                name: "Produce Flame",
                school: "Conjuration",
                description: "A flickering flame appears in your hand. The flame remains there for the duration and harms neither you nor your equipment."
            },
            {
                name: "Resistance",
                school: "Abjuration",
                description: "You touch one willing creature. Once before the spell ends, the target can roll a d4 and add the number rolled to one saving throw of its choice."
            }
        ],
        firstLevel: [
            {
                name: "Cure Wounds",
                school: "Evocation",
                description: "A creature you touch regains a number of hit points equal to 1d8 + your spellcasting ability modifier."
            },
            {
                name: "Entangle",
                school: "Conjuration",
                description: "Grasping weeds and vines sprout from the ground in a 20-foot square starting from a point within range."
            },
            {
                name: "Faerie Fire",
                school: "Evocation",
                description: "Each object in a 20-foot cube within range is outlined in blue, green, or violet light."
            },
            {
                name: "Goodberry",
                school: "Transmutation",
                description: "Up to ten berries appear in your hand and are infused with magic for the duration."
            },
            {
                name: "Speak with Animals",
                school: "Divination",
                description: "You gain the ability to comprehend and verbally communicate with beasts for the duration."
            }
        ]
    }
};

// Helper functions for spell data
function getSpellsForClass(className) {
    return SPELLS[className] || { cantrips: [], firstLevel: [] };
}

function getSpellByName(spellName, className = null) {
    if (className) {
        const classSpells = getSpellsForClass(className);
        return [...classSpells.cantrips, ...classSpells.firstLevel]
            .find(spell => spell.name.toLowerCase() === spellName.toLowerCase());
    }
    
    // Search all classes
    for (const cls of Object.keys(SPELLS)) {
        const classSpells = getSpellsForClass(cls);
        const spell = [...classSpells.cantrips, ...classSpells.firstLevel]
            .find(spell => spell.name.toLowerCase() === spellName.toLowerCase());
        if (spell) return spell;
    }
    
    return null;
}

function getSpellSchools() {
    const schools = new Set();
    Object.values(SPELLS).forEach(classSpells => {
        [...classSpells.cantrips, ...classSpells.firstLevel].forEach(spell => {
            schools.add(spell.school);
        });
    });
    return Array.from(schools).sort();
}

function filterSpellsBySchool(className, school) {
    const classSpells = getSpellsForClass(className);
    return {
        cantrips: classSpells.cantrips.filter(spell => spell.school === school),
        firstLevel: classSpells.firstLevel.filter(spell => spell.school === school)
    };
}

function searchSpells(query, className = null) {
    const searchTerm = query.toLowerCase();
    const results = { cantrips: [], firstLevel: [] };
    
    const classesToSearch = className ? [className] : Object.keys(SPELLS);
    
    classesToSearch.forEach(cls => {
        const classSpells = getSpellsForClass(cls);
        
        classSpells.cantrips.forEach(spell => {
            if (spell.name.toLowerCase().includes(searchTerm) || 
                spell.description.toLowerCase().includes(searchTerm) ||
                spell.school.toLowerCase().includes(searchTerm)) {
                results.cantrips.push({ ...spell, class: cls });
            }
        });
        
        classSpells.firstLevel.forEach(spell => {
            if (spell.name.toLowerCase().includes(searchTerm) || 
                spell.description.toLowerCase().includes(searchTerm) ||
                spell.school.toLowerCase().includes(searchTerm)) {
                results.firstLevel.push({ ...spell, class: cls });
            }
        });
    });
    
    return results;
}

// Export for use in other scripts
window.SPELLS = SPELLS;
window.getSpellsForClass = getSpellsForClass;
window.getSpellByName = getSpellByName;
window.getSpellSchools = getSpellSchools;
window.filterSpellsBySchool = filterSpellsBySchool;
window.searchSpells = searchSpells;
