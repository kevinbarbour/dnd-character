// D&D 5e Classes Data
const CLASSES = [
    {
        name: "Fighter",
        description: "A master of martial combat, skilled with a variety of weapons and armor. Fighters learn the basics of all combat styles.",
        hitDie: 10,
        primaryAbility: ["Strength", "Dexterity"],
        savingThrowProficiencies: ["Strength", "Constitution"],
        skillChoices: 2,
        skillOptions: ["Acrobatics", "Animal Handling", "Athletics", "History", "Insight", "Intimidation", "Perception", "Survival"],
        isSpellcaster: false,
        features: [
            "Fighting Style",
            "Second Wind"
        ]
    },
    {
        name: "Wizard",
        description: "A scholarly magic-user capable of manipulating the structures of spellcasting to cast spells of explosive fire, arcing lightning, subtle deception, and brute-force mind control.",
        hitDie: 6,
        primaryAbility: ["Intelligence"],
        savingThrowProficiencies: ["Intelligence", "Wisdom"],
        skillChoices: 2,
        skillOptions: ["Arcana", "History", "Insight", "Investigation", "Medicine", "Religion"],
        isSpellcaster: true,
        spellcastingAbility: "Intelligence",
        cantripsKnown: 3,
        spellsKnown: 2,
        spellSlots: {
            1: 2
        },
        features: [
            "Spellcasting",
            "Arcane Recovery"
        ]
    },
    {
        name: "Cleric",
        description: "A priestly champion who wields divine magic in service of a higher power. Clerics combine the helpful magic of healing and inspiring their allies with spells that harm and hinder foes.",
        hitDie: 8,
        primaryAbility: ["Wisdom"],
        savingThrowProficiencies: ["Wisdom", "Charisma"],
        skillChoices: 2,
        skillOptions: ["History", "Insight", "Medicine", "Persuasion", "Religion"],
        isSpellcaster: true,
        spellcastingAbility: "Wisdom",
        cantripsKnown: 3,
        spellsKnown: 2,
        spellSlots: {
            1: 2
        },
        features: [
            "Spellcasting",
            "Divine Domain"
        ]
    },
    {
        name: "Rogue",
        description: "A scoundrel who uses stealth and trickery to overcome obstacles. Rogues have a knack for finding the solution to just about any problem, demonstrating a resourcefulness and versatility that is the cornerstone of any successful adventuring party.",
        hitDie: 8,
        primaryAbility: ["Dexterity"],
        savingThrowProficiencies: ["Dexterity", "Intelligence"],
        skillChoices: 4,
        skillOptions: ["Acrobatics", "Athletics", "Deception", "Insight", "Intimidation", "Investigation", "Perception", "Performance", "Persuasion", "Sleight of Hand", "Stealth"],
        isSpellcaster: false,
        features: [
            "Expertise",
            "Sneak Attack",
            "Thieves' Cant"
        ]
    },
    {
        name: "Ranger",
        description: "A warrior of the wilderness, rangers specialize in hunting the monsters that threaten the edges of civilization—humanoid raiders, rampaging beasts and monstrosities, terrible giants, and deadly dragons.",
        hitDie: 10,
        primaryAbility: ["Dexterity", "Wisdom"],
        savingThrowProficiencies: ["Strength", "Dexterity"],
        skillChoices: 3,
        skillOptions: ["Animal Handling", "Athletics", "Insight", "Investigation", "Nature", "Perception", "Stealth", "Survival"],
        isSpellcaster: false, // Rangers get spells at level 2
        features: [
            "Favored Enemy",
            "Natural Explorer"
        ]
    },
    {
        name: "Bard",
        description: "A master of song, speech, and the magic they contain. Bards say that the multiverse was spoken into existence, that the words of the gods gave it shape, and that echoes of these primordial Words of Creation still resound throughout the cosmos.",
        hitDie: 8,
        primaryAbility: ["Charisma"],
        savingThrowProficiencies: ["Dexterity", "Charisma"],
        skillChoices: 3,
        skillOptions: ["Deception", "History", "Investigation", "Persuasion", "Religion", "Sleight of Hand"],
        isSpellcaster: true,
        spellcastingAbility: "Charisma",
        cantripsKnown: 2,
        spellsKnown: 4,
        spellSlots: {
            1: 2
        },
        features: [
            "Spellcasting",
            "Bardic Inspiration"
        ]
    },
    {
        name: "Barbarian",
        description: "A fierce warrior of primitive background who can enter a battle rage. For some, their rage springs from a communion with fierce animal spirits. Others draw from a roiling reservoir of anger at a world full of pain.",
        hitDie: 12,
        primaryAbility: ["Strength"],
        savingThrowProficiencies: ["Strength", "Constitution"],
        skillChoices: 2,
        skillOptions: ["Animal Handling", "Athletics", "Intimidation", "Nature", "Perception", "Survival"],
        isSpellcaster: false,
        features: [
            "Rage",
            "Unarmored Defense"
        ]
    },
    {
        name: "Sorcerer",
        description: "A spellcaster who draws on inherent magic from a gift or bloodline. Sorcerers have no use for the spellbooks and ancient knowledge that wizards rely on, nor do they rely on a patron to grant their spells as warlocks do.",
        hitDie: 6,
        primaryAbility: ["Charisma"],
        savingThrowProficiencies: ["Constitution", "Charisma"],
        skillChoices: 2,
        skillOptions: ["Arcana", "Deception", "Insight", "Intimidation", "Persuasion", "Religion"],
        isSpellcaster: true,
        spellcastingAbility: "Charisma",
        cantripsKnown: 4,
        spellsKnown: 2,
        spellSlots: {
            1: 2
        },
        features: [
            "Spellcasting",
            "Sorcerous Origin"
        ]
    },
    {
        name: "Warlock",
        description: "A wielder of magic that is derived from a bargain with an extraplanar entity. Warlocks are driven by an insatiable need for knowledge and power, which compels them into their pacts and shapes their lives.",
        hitDie: 8,
        primaryAbility: ["Charisma"],
        savingThrowProficiencies: ["Wisdom", "Charisma"],
        skillChoices: 2,
        skillOptions: ["Arcana", "Deception", "History", "Intimidation", "Investigation", "Nature", "Religion"],
        isSpellcaster: true,
        spellcastingAbility: "Charisma",
        cantripsKnown: 2,
        spellsKnown: 2,
        spellSlots: {
            1: 1
        },
        features: [
            "Otherworldly Patron",
            "Pact Magic"
        ]
    },
    {
        name: "Paladin",
        description: "A holy warrior bound to a sacred oath. Whether sworn before a god's altar and the witness of a priest, in a sacred glade before nature spirits and fey beings, or in a moment of desperation and grief with the dead as the only witness, a paladin's oath is a powerful bond.",
        hitDie: 10,
        primaryAbility: ["Strength", "Charisma"],
        savingThrowProficiencies: ["Wisdom", "Charisma"],
        skillChoices: 2,
        skillOptions: ["Athletics", "Insight", "Intimidation", "Medicine", "Persuasion", "Religion"],
        isSpellcaster: false, // Paladins get spells at level 2
        features: [
            "Divine Sense",
            "Lay on Hands"
        ]
    },
    {
        name: "Druid",
        description: "A priest of the Old Faith, wielding the elemental forces of nature and transforming into animal forms. Druids revere nature above all, gaining their spells and other magical powers either from the force of nature itself or from a nature deity.",
        hitDie: 8,
        primaryAbility: ["Wisdom"],
        savingThrowProficiencies: ["Intelligence", "Wisdom"],
        skillChoices: 2,
        skillOptions: ["Arcana", "Animal Handling", "Insight", "Medicine", "Nature", "Perception", "Religion", "Survival"],
        isSpellcaster: true,
        spellcastingAbility: "Wisdom",
        cantripsKnown: 2,
        spellsKnown: 2,
        spellSlots: {
            1: 2
        },
        features: [
            "Druidcraft",
            "Natural Recovery"
        ]
    },
    {
        name: "Monk",
        description: "A master of martial arts, harnessing the power of the body in pursuit of physical and spiritual perfection. Monks make careful study of a magical energy that most monastic traditions call ki.",
        hitDie: 8,
        primaryAbility: ["Dexterity", "Wisdom"],
        savingThrowProficiencies: ["Strength", "Dexterity"],
        skillChoices: 2,
        skillOptions: ["Acrobatics", "Athletics", "History", "Insight", "Religion", "Stealth"],
        isSpellcaster: false,
        features: [
            "Unarmored Defense",
            "Martial Arts"
        ]
    }
];

// Spellcasting classes for easy reference
const SPELLCASTING_CLASSES = [
    "Wizard",
    "Cleric", 
    "Bard",
    "Sorcerer",
    "Warlock",
    "Druid"
];

// Helper functions for class data
function getClassByName(name) {
    return CLASSES.find(cls => cls.name.toLowerCase() === name.toLowerCase());
}

function isSpellcaster(className) {
    return SPELLCASTING_CLASSES.includes(className);
}

function getSpellcastingInfo(className) {
    const cls = getClassByName(className);
    if (!cls || !cls.isSpellcaster) {
        return null;
    }
    
    return {
        ability: cls.spellcastingAbility,
        cantripsKnown: cls.cantripsKnown,
        spellsKnown: cls.spellsKnown,
        spellSlots: cls.spellSlots
    };
}

function getClassHitPoints(className, constitution) {
    const cls = getClassByName(className);
    if (!cls) return 8; // Default
    
    const constitutionModifier = Math.floor((constitution - 10) / 2);
    return cls.hitDie + constitutionModifier;
}

function getClassSkillChoices(className) {
    const cls = getClassByName(className);
    return cls ? {
        count: cls.skillChoices,
        options: cls.skillOptions
    } : { count: 2, options: [] };
}

// Export for use in other scripts
window.CLASSES = CLASSES;
window.SPELLCASTING_CLASSES = SPELLCASTING_CLASSES;
window.getClassByName = getClassByName;
window.isSpellcaster = isSpellcaster;
window.getSpellcastingInfo = getSpellcastingInfo;
window.getClassHitPoints = getClassHitPoints;
window.getClassSkillChoices = getClassSkillChoices;
