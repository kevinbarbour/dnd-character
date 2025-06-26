// D&D 5e Backgrounds Data
const BACKGROUNDS = [
    {
        name: "Acolyte",
        description: "You have spent your life in the service of a temple to a specific god or pantheon of gods. You act as an intermediary between the realm of the holy and the mortal world, performing sacred rites and offering sacrifices in order to conduct worshipers into the presence of the divine.",
        skillProficiencies: ["Insight", "Religion"],
        languages: 2,
        equipment: [
            "Holy symbol",
            "Prayer book or prayer wheel",
            "5 sticks of incense",
            "Vestments",
            "Common clothes",
            "Belt pouch with 15 gp"
        ],
        feature: "Shelter of the Faithful"
    },
    {
        name: "Criminal",
        description: "You are an experienced criminal with a history of breaking the law. You have spent a lot of time among other criminals and still have contacts within the criminal underworld. You're far closer than most people to the world of murder, theft, and violence that pervades the underbelly of civilization.",
        skillProficiencies: ["Deception", "Stealth"],
        toolProficiencies: ["Thieves' tools", "Gaming set"],
        languages: 0,
        equipment: [
            "Crowbar",
            "Dark common clothes with hood",
            "Belt pouch with 15 gp"
        ],
        feature: "Criminal Contact"
    },
    {
        name: "Folk Hero",
        description: "You come from a humble social rank, but you are destined for so much more. Already the people of your home village regard you as their champion, and your destiny calls you to stand against the tyrants and monsters that threaten the common folk everywhere.",
        skillProficiencies: ["Animal Handling", "Survival"],
        toolProficiencies: ["Artisan's tools", "Vehicles (land)"],
        languages: 0,
        equipment: [
            "Artisan's tools",
            "Shovel",
            "Iron pot",
            "Common clothes",
            "Belt pouch with 10 gp"
        ],
        feature: "Rustic Hospitality"
    },
    {
        name: "Noble",
        description: "You understand wealth, power, and privilege. You carry a noble title, and your family owns land, collects taxes, and wields significant political influence. You might be a pampered aristocrat unfamiliar with work or discomfort, a former merchant just elevated to the nobility, or a disinherited scoundrel with a disproportionate sense of entitlement.",
        skillProficiencies: ["History", "Persuasion"],
        toolProficiencies: ["Gaming set"],
        languages: 1,
        equipment: [
            "Signet ring",
            "Scroll of pedigree",
            "Fine clothes",
            "Purse with 25 gp"
        ],
        feature: "Position of Privilege"
    },
    {
        name: "Sage",
        description: "You spent years learning the lore of the multiverse. You scoured manuscripts, studied scrolls, and listened to the greatest experts on the subjects that interest you. Your efforts have made you a master in your fields of study.",
        skillProficiencies: ["Arcana", "History"],
        languages: 2,
        equipment: [
            "Bottle of black ink",
            "Quill",
            "Small knife",
            "Letter from colleague",
            "Common clothes",
            "Belt pouch with 10 gp"
        ],
        feature: "Researcher"
    },
    {
        name: "Soldier",
        description: "War has been your life for as long as you care to remember. You trained as a youth, studied the use of weapons and armor, learned basic survival techniques, including how to stay alive on the battlefield. You might have been part of a standing national army or a mercenary company, or perhaps a member of a local militia who rose to prominence during a recent war.",
        skillProficiencies: ["Athletics", "Intimidation"],
        toolProficiencies: ["Gaming set", "Vehicles (land)"],
        languages: 0,
        equipment: [
            "Insignia of rank",
            "Trophy from fallen enemy",
            "Deck of cards",
            "Common clothes",
            "Belt pouch with 10 gp"
        ],
        feature: "Military Rank"
    },
    {
        name: "Charlatan",
        description: "You have always had a way with people. You know what makes them tick, you can tease out their hearts' desires after a few minutes of conversation, and with a few leading questions you can read them like they were children's books. It's a useful talent, and one that you're perfectly willing to use for your advantage.",
        skillProficiencies: ["Deception", "Sleight of Hand"],
        toolProficiencies: ["Forgery kit", "Disguise kit"],
        languages: 0,
        equipment: [
            "Forgery kit",
            "Disguise kit",
            "Signet ring of fictional person",
            "Fine clothes",
            "Belt pouch with 15 gp"
        ],
        feature: "False Identity"
    },
    {
        name: "Entertainer",
        description: "You thrive in front of an audience. You know how to entrance them, entertain them, and even inspire them. Your poetics can stir the hearts of those who hear you, awakening grief or joy, laughter or anger. Your music raises their spirits or captures their sorrow. Your dance steps captivate, your humor cuts to the quick.",
        skillProficiencies: ["Acrobatics", "Performance"],
        toolProficiencies: ["Disguise kit", "Musical instrument"],
        languages: 0,
        equipment: [
            "Musical instrument",
            "Favor of an admirer",
            "Costume",
            "Belt pouch with 15 gp"
        ],
        feature: "By Popular Demand"
    },
    {
        name: "Guild Artisan",
        description: "You are a member of an artisan's guild, skilled in a particular field and closely associated with other artisans. You are a well-established part of the mercantile world, freed by talent and wealth from the constraints of a feudal social order.",
        skillProficiencies: ["Insight", "Persuasion"],
        toolProficiencies: ["Artisan's tools"],
        languages: 1,
        equipment: [
            "Artisan's tools",
            "Letter of introduction from guild",
            "Traveler's clothes",
            "Belt pouch with 15 gp"
        ],
        feature: "Guild Membership"
    },
    {
        name: "Hermit",
        description: "You lived in seclusion—either in a sheltered community such as a monastery, or entirely alone—for a formative part of your life. In your time apart from the clamor of society, you found quiet, solitude, and perhaps some of the answers you were looking for.",
        skillProficiencies: ["Medicine", "Religion"],
        toolProficiencies: ["Herbalism kit"],
        languages: 1,
        equipment: [
            "Herbalism kit",
            "Scroll case with spiritual writings",
            "Winter blanket",
            "Belt pouch with 5 gp"
        ],
        feature: "Discovery"
    },
    {
        name: "Outlander",
        description: "You grew up in the wilds, far from civilization and the comforts of town and technology. You've witnessed the migration of herds larger than forests, survived weather more extreme than any city-dweller could comprehend, and enjoyed the solitude of being the only thinking creature for miles in any direction.",
        skillProficiencies: ["Athletics", "Survival"],
        toolProficiencies: ["Musical instrument"],
        languages: 1,
        equipment: [
            "Staff",
            "Hunting trap",
            "Traveler's clothes",
            "Belt pouch with 10 gp"
        ],
        feature: "Wanderer"
    },
    {
        name: "Sailor",
        description: "You sailed on a seagoing vessel for years. In that time, you faced down mighty storms, monsters of the deep, and those who wanted to sink your craft to the bottomless depths. Your first love is the distant line of the horizon, but the time has come to try your hand at something new.",
        skillProficiencies: ["Athletics", "Perception"],
        toolProficiencies: ["Navigator's tools", "Vehicles (water)"],
        languages: 0,
        equipment: [
            "Belaying pin",
            "50 feet of silk rope",
            "Lucky charm",
            "Common clothes",
            "Belt pouch with 10 gp"
        ],
        feature: "Ship's Passage"
    }
];

// Helper functions for background data
function getBackgroundByName(name) {
    return BACKGROUNDS.find(bg => bg.name.toLowerCase() === name.toLowerCase());
}

function getBackgroundSkills(backgroundName) {
    const background = getBackgroundByName(backgroundName);
    return background ? background.skillProficiencies : [];
}

function getBackgroundEquipment(backgroundName) {
    const background = getBackgroundByName(backgroundName);
    return background ? background.equipment : [];
}

function getBackgroundFeature(backgroundName) {
    const background = getBackgroundByName(backgroundName);
    return background ? background.feature : "";
}

// Export for use in other scripts
window.BACKGROUNDS = BACKGROUNDS;
window.getBackgroundByName = getBackgroundByName;
window.getBackgroundSkills = getBackgroundSkills;
window.getBackgroundEquipment = getBackgroundEquipment;
window.getBackgroundFeature = getBackgroundFeature;
