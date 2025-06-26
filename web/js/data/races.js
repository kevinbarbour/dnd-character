// D&D 5e Races Data
const RACES = [
    {
        name: "Human",
        description: "Humans are the most adaptable and ambitious people among the common races. Whatever drives them, humans are the innovators, the achievers, and the pioneers of the worlds.",
        abilityScoreIncrease: {
            strength: 1,
            dexterity: 1,
            constitution: 1,
            intelligence: 1,
            wisdom: 1,
            charisma: 1
        },
        size: "Medium",
        speed: 30,
        traits: [
            "Extra Language",
            "Extra Skill"
        ]
    },
    {
        name: "Elf",
        description: "Elves are a magical people of otherworldly grace, living in places of ethereal beauty, in the midst of ancient forests or in silvery spires glittering with faerie light.",
        abilityScoreIncrease: {
            dexterity: 2
        },
        size: "Medium",
        speed: 30,
        traits: [
            "Darkvision",
            "Keen Senses",
            "Fey Ancestry",
            "Trance"
        ]
    },
    {
        name: "Dwarf",
        description: "Bold and hardy, dwarves are known as skilled warriors, miners, and workers of stone and metal. Though they stand well under 5 feet tall, dwarves are so broad and compact that they can weigh as much as a human standing nearly two feet taller.",
        abilityScoreIncrease: {
            constitution: 2
        },
        size: "Medium",
        speed: 25,
        traits: [
            "Darkvision",
            "Dwarven Resilience",
            "Dwarven Combat Training",
            "Stonecunning"
        ]
    },
    {
        name: "Halfling",
        description: "The comforts of home are the goals of most halflings' lives: a place to settle in peace and quiet, far from marauding monsters and clashing armies; a blazing fire and a generous meal; fine drink and fine conversation.",
        abilityScoreIncrease: {
            dexterity: 2
        },
        size: "Small",
        speed: 25,
        traits: [
            "Lucky",
            "Brave",
            "Halfling Nimbleness"
        ]
    },
    {
        name: "Dragonborn",
        description: "Born of dragons, as their name proclaims, the dragonborn walk proudly through a world that greets them with fearful incomprehension. Shaped by draconic gods or the dragons themselves, dragonborn originally hatched from dragon eggs as a unique race, combining the best attributes of dragons and humanoids.",
        abilityScoreIncrease: {
            strength: 2,
            charisma: 1
        },
        size: "Medium",
        speed: 30,
        traits: [
            "Draconic Ancestry",
            "Breath Weapon",
            "Damage Resistance"
        ]
    },
    {
        name: "Gnome",
        description: "A gnome's energy and enthusiasm for living shines through every inch of his or her tiny body. Gnomes average slightly over 3 feet tall and weigh 40 to 45 pounds. Their tan or brown faces are usually adorned with broad smiles (beneath their prodigious noses), and their bright eyes shine with excitement.",
        abilityScoreIncrease: {
            intelligence: 2
        },
        size: "Small",
        speed: 25,
        traits: [
            "Darkvision",
            "Gnome Cunning"
        ]
    },
    {
        name: "Half-Elf",
        description: "Walking in two worlds but truly belonging to neither, half-elves combine what some say are the best qualities of their elf and human parents: human curiosity, inventiveness, and ambition tempered by the refined senses, love of nature, and artistic tastes of the elves.",
        abilityScoreIncrease: {
            charisma: 2,
            // Player chooses two different abilities to increase by 1
        },
        size: "Medium",
        speed: 30,
        traits: [
            "Darkvision",
            "Fey Ancestry",
            "Extra Skill Versatility"
        ]
    },
    {
        name: "Half-Orc",
        description: "Whether united under the leadership of a mighty warlock or having fought to a standstill after years of conflict, orc and human tribes sometimes form alliances, joining forces into a larger horde to the terror of civilized lands nearby. When these alliances are sealed by marriages, half-orcs are born.",
        abilityScoreIncrease: {
            strength: 2,
            constitution: 1
        },
        size: "Medium",
        speed: 30,
        traits: [
            "Darkvision",
            "Relentless Endurance",
            "Savage Attacks"
        ]
    },
    {
        name: "Tiefling",
        description: "To be greeted with stares and whispers, to suffer violence and insult on the street, to see mistrust and fear in every eye: this is the lot of the tiefling. And to twist the knife, tieflings know that this is because a pact struck generations ago infused the essence of Asmodeus—overlord of the Nine Hells—into their bloodline.",
        abilityScoreIncrease: {
            intelligence: 1,
            charisma: 2
        },
        size: "Medium",
        speed: 30,
        traits: [
            "Darkvision",
            "Hellish Resistance",
            "Infernal Legacy"
        ]
    }
];

// Helper functions for race data
function getRaceByName(name) {
    return RACES.find(race => race.name.toLowerCase() === name.toLowerCase());
}

function getRaceAbilityBonus(raceName) {
    const race = getRaceByName(raceName);
    return race ? race.abilityScoreIncrease : {};
}

function applyRacialBonuses(abilities, raceName) {
    const bonuses = getRaceAbilityBonus(raceName);
    const result = { ...abilities };
    
    Object.keys(bonuses).forEach(ability => {
        if (result[ability] !== undefined) {
            result[ability] += bonuses[ability];
        }
    });
    
    return result;
}

// Export for use in other scripts
window.RACES = RACES;
window.getRaceByName = getRaceByName;
window.getRaceAbilityBonus = getRaceAbilityBonus;
window.applyRacialBonuses = applyRacialBonuses;
