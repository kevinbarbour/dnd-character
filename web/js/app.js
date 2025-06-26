// D&D Character Creator - Main Application Logic

// Ability score generation methods
const ABILITY_GENERATION_METHODS = {
    DICE_ROLL: 'dice',
    POINT_BUY: 'pointbuy',
    STANDARD_ARRAY: 'standard'
};

// Standard array values
const STANDARD_ARRAY = [15, 14, 13, 12, 10, 8];

// Point buy system
const POINT_BUY_COSTS = {
    8: 0, 9: 1, 10: 2, 11: 3, 12: 4, 13: 5, 14: 7, 15: 9
};

// Utility functions
function getAbilityModifier(score) {
    return Math.floor((score - 10) / 2);
}

function formatModifier(modifier) {
    return modifier >= 0 ? `+${modifier}` : `${modifier}`;
}

function rollDice(sides, count = 1) {
    let total = 0;
    for (let i = 0; i < count; i++) {
        total += Math.floor(Math.random() * sides) + 1;
    }
    return total;
}

function roll4d6DropLowest() {
    const rolls = [];
    for (let i = 0; i < 4; i++) {
        rolls.push(rollDice(6));
    }
    rolls.sort((a, b) => b - a);
    return rolls.slice(0, 3).reduce((sum, roll) => sum + roll, 0);
}

// Ability score generation functions
function generateAbilityScores(method) {
    const abilities = {
        strength: 10,
        dexterity: 10,
        constitution: 10,
        intelligence: 10,
        wisdom: 10,
        charisma: 10
    };

    switch (method) {
        case ABILITY_GENERATION_METHODS.DICE_ROLL:
            abilities.strength = roll4d6DropLowest();
            abilities.dexterity = roll4d6DropLowest();
            abilities.constitution = roll4d6DropLowest();
            abilities.intelligence = roll4d6DropLowest();
            abilities.wisdom = roll4d6DropLowest();
            abilities.charisma = roll4d6DropLowest();
            break;

        case ABILITY_GENERATION_METHODS.STANDARD_ARRAY:
            // Return standard array - user will assign manually
            const shuffledArray = [...STANDARD_ARRAY].sort(() => Math.random() - 0.5);
            abilities.strength = shuffledArray[0];
            abilities.dexterity = shuffledArray[1];
            abilities.constitution = shuffledArray[2];
            abilities.intelligence = shuffledArray[3];
            abilities.wisdom = shuffledArray[4];
            abilities.charisma = shuffledArray[5];
            break;

        case ABILITY_GENERATION_METHODS.POINT_BUY:
            // Start with 8s for point buy
            abilities.strength = 8;
            abilities.dexterity = 8;
            abilities.constitution = 8;
            abilities.intelligence = 8;
            abilities.wisdom = 8;
            abilities.charisma = 8;
            break;
    }

    return abilities;
}

// Character creation functions
function createCharacter(characterData) {
    const character = {
        name: characterData.name,
        race: characterData.race,
        class: characterData.class,
        background: characterData.background,
        level: 1,
        abilities: characterData.abilities,
        hitPoints: calculateHitPoints(characterData.class, characterData.abilities.constitution),
        armorClass: calculateArmorClass(characterData.abilities.dexterity),
        proficiencyBonus: 2, // Level 1
        spells: characterData.spells || [],
        equipment: getStartingEquipment(characterData.class, characterData.background),
        skills: characterData.skills || [],
        savingThrows: getSavingThrowProficiencies(characterData.class)
    };

    return character;
}

function calculateHitPoints(className, constitution) {
    const cls = getClassByName(className);
    const constitutionModifier = getAbilityModifier(constitution);
    return cls.hitDie + constitutionModifier;
}

function calculateArmorClass(dexterity) {
    // Base AC (10 + Dex modifier) - assumes no armor
    return 10 + getAbilityModifier(dexterity);
}

function getSavingThrowProficiencies(className) {
    const cls = getClassByName(className);
    return cls ? cls.savingThrowProficiencies : [];
}

function getStartingEquipment(className, backgroundName) {
    const equipment = [];
    
    // Add class equipment (simplified)
    const classEquipment = {
        'Fighter': ['Chain mail', 'Shield', 'Longsword', 'Light crossbow'],
        'Wizard': ['Quarterstaff', 'Dagger', 'Spellbook', 'Component pouch'],
        'Cleric': ['Scale mail', 'Shield', 'Mace', 'Holy symbol'],
        'Rogue': ['Leather armor', 'Shortsword', 'Thieves\' tools', 'Dagger'],
        'Ranger': ['Scale mail', 'Shortsword', 'Longbow', 'Quiver'],
        'Bard': ['Leather armor', 'Rapier', 'Lute', 'Dagger'],
        'Barbarian': ['Shield', 'Handaxe', 'Javelin', 'Explorer\'s pack'],
        'Sorcerer': ['Light crossbow', 'Dagger', 'Component pouch', 'Explorer\'s pack'],
        'Warlock': ['Light crossbow', 'Simple weapon', 'Leather armor', 'Component pouch'],
        'Paladin': ['Chain mail', 'Shield', 'Longsword', 'Javelin'],
        'Druid': ['Leather armor', 'Shield', 'Scimitar', 'Druidcraft focus'],
        'Monk': ['Shortsword', 'Dart', 'Explorer\'s pack', 'Dungeoneer\'s pack']
    };

    if (classEquipment[className]) {
        equipment.push(...classEquipment[className]);
    }

    // Add background equipment
    const backgroundEquipment = getBackgroundEquipment(backgroundName);
    if (backgroundEquipment) {
        equipment.push(...backgroundEquipment);
    }

    return equipment;
}

// Spell validation functions
function validateSpellSelection(className, selectedSpells) {
    const spellcastingInfo = getSpellcastingInfo(className);
    if (!spellcastingInfo) {
        return { valid: true, errors: [] };
    }

    const errors = [];
    const cantrips = selectedSpells.filter(spell => {
        const spellData = getSpellByName(spell, className);
        return spellData && getSpellsForClass(className).cantrips.includes(spellData);
    });

    const firstLevelSpells = selectedSpells.filter(spell => {
        const spellData = getSpellByName(spell, className);
        return spellData && getSpellsForClass(className).firstLevel.includes(spellData);
    });

    // Check cantrip count
    if (cantrips.length > spellcastingInfo.cantripsKnown) {
        errors.push(`Too many cantrips selected. Maximum: ${spellcastingInfo.cantripsKnown}`);
    }

    // Check spell count
    if (firstLevelSpells.length > spellcastingInfo.spellsKnown) {
        errors.push(`Too many spells selected. Maximum: ${spellcastingInfo.spellsKnown}`);
    }

    return {
        valid: errors.length === 0,
        errors: errors
    };
}

// Form validation functions
function validateCharacterForm(formData) {
    const errors = [];

    if (!formData.name || formData.name.trim().length === 0) {
        errors.push('Character name is required');
    }

    if (!formData.race) {
        errors.push('Race selection is required');
    }

    if (!formData.class) {
        errors.push('Class selection is required');
    }

    if (!formData.background) {
        errors.push('Background selection is required');
    }

    // Validate ability scores
    const abilities = ['strength', 'dexterity', 'constitution', 'intelligence', 'wisdom', 'charisma'];
    abilities.forEach(ability => {
        const score = parseInt(formData.abilities[ability]);
        if (isNaN(score) || score < 3 || score > 20) {
            errors.push(`${ability.charAt(0).toUpperCase() + ability.slice(1)} must be between 3 and 20`);
        }
    });

    // Validate spells if spellcaster
    if (isSpellcaster(formData.class)) {
        const spellValidation = validateSpellSelection(formData.class, formData.spells || []);
        if (!spellValidation.valid) {
            errors.push(...spellValidation.errors);
        }
    }

    return {
        valid: errors.length === 0,
        errors: errors
    };
}

// UI Helper functions
function showAlert(message, type = 'info') {
    const alertDiv = document.createElement('div');
    alertDiv.className = `alert alert-${type} alert-dismissible fade show`;
    alertDiv.innerHTML = `
        ${message}
        <button type="button" class="btn-close" data-bs-dismiss="alert"></button>
    `;

    // Insert at top of main content
    const container = document.querySelector('.container');
    if (container) {
        container.insertBefore(alertDiv, container.firstChild);
        
        // Auto-dismiss after 5 seconds
        setTimeout(() => {
            if (alertDiv.parentNode) {
                alertDiv.remove();
            }
        }, 5000);
    }
}

function showLoading(element, show = true) {
    if (show) {
        element.classList.add('loading');
        element.style.pointerEvents = 'none';
    } else {
        element.classList.remove('loading');
        element.style.pointerEvents = '';
    }
}

function formatDate(dateString) {
    const date = new Date(dateString);
    return date.toLocaleDateString() + ' ' + date.toLocaleTimeString();
}

// Character display functions
function getCharacterSummary(character) {
    const race = getRaceByName(character.race);
    const cls = getClassByName(character.class);
    const background = getBackgroundByName(character.background);

    return {
        name: character.name,
        level: character.level,
        race: race ? race.name : character.race,
        class: cls ? cls.name : character.class,
        background: background ? background.name : character.background,
        hitPoints: character.hitPoints,
        armorClass: character.armorClass,
        abilities: character.abilities,
        spells: character.spells || [],
        isSpellcaster: isSpellcaster(character.class)
    };
}

function renderCharacterCard(character) {
    const summary = getCharacterSummary(character);
    
    return `
        <div class="character-card" data-character-id="${character.id}">
            <div class="character-header">
                <div>
                    <h3 class="character-name">${summary.name}</h3>
                    <p class="character-class-race">Level ${summary.level} ${summary.race} ${summary.class}</p>
                    <p class="text-muted">${summary.background}</p>
                </div>
                <div class="character-actions">
                    <button class="btn btn-sm btn-primary" onclick="viewCharacter(${character.id})">
                        <i class="fas fa-eye"></i> View
                    </button>
                    <button class="btn btn-sm btn-info" onclick="shareCharacter(${character.id})">
                        <i class="fas fa-share"></i> Share
                    </button>
                    <button class="btn btn-sm btn-secondary" onclick="exportCharacter(${character.id})">
                        <i class="fas fa-download"></i> Export
                    </button>
                    <button class="btn btn-sm btn-danger" onclick="deleteCharacterConfirm(${character.id})">
                        <i class="fas fa-trash"></i> Delete
                    </button>
                </div>
            </div>
            
            <div class="character-stats">
                <div class="stat-item">
                    <div class="stat-value">${summary.hitPoints}</div>
                    <div class="stat-label">HP</div>
                </div>
                <div class="stat-item">
                    <div class="stat-value">${summary.armorClass}</div>
                    <div class="stat-label">AC</div>
                </div>
                <div class="stat-item">
                    <div class="stat-value">${summary.abilities.strength}</div>
                    <div class="stat-label">STR</div>
                </div>
                <div class="stat-item">
                    <div class="stat-value">${summary.abilities.dexterity}</div>
                    <div class="stat-label">DEX</div>
                </div>
                <div class="stat-item">
                    <div class="stat-value">${summary.abilities.constitution}</div>
                    <div class="stat-label">CON</div>
                </div>
                <div class="stat-item">
                    <div class="stat-value">${summary.abilities.intelligence}</div>
                    <div class="stat-label">INT</div>
                </div>
                <div class="stat-item">
                    <div class="stat-value">${summary.abilities.wisdom}</div>
                    <div class="stat-label">WIS</div>
                </div>
                <div class="stat-item">
                    <div class="stat-value">${summary.abilities.charisma}</div>
                    <div class="stat-label">CHA</div>
                </div>
            </div>
            
            ${summary.isSpellcaster && summary.spells.length > 0 ? `
                <div class="character-spells mt-3">
                    <h6>Spells (${summary.spells.length})</h6>
                    <div class="spell-list">
                        ${summary.spells.slice(0, 3).map(spell => `<span class="badge bg-info me-1">${spell}</span>`).join('')}
                        ${summary.spells.length > 3 ? `<span class="text-muted">+${summary.spells.length - 3} more</span>` : ''}
                    </div>
                </div>
            ` : ''}
            
            <div class="character-meta mt-2">
                <small class="text-muted">
                    Created: ${formatDate(character.created)}
                    ${character.updated !== character.created ? `• Updated: ${formatDate(character.updated)}` : ''}
                </small>
            </div>
        </div>
    `;
}

// Global character management functions
function viewCharacter(id) {
    window.location.href = `character.html?id=${id}`;
}

function deleteCharacterConfirm(id) {
    const character = getCharacter(id);
    if (!character) {
        showAlert('Character not found', 'error');
        return;
    }

    if (confirm(`Are you sure you want to delete "${character.name}"? This cannot be undone.`)) {
        try {
            deleteCharacter(id);
            showAlert(`Character "${character.name}" deleted successfully`, 'success');
            
            // Refresh the current page
            if (window.location.pathname.includes('characters.html')) {
                loadCharactersList();
            } else {
                // Update character count on home page
                updateHomepageStats();
            }
        } catch (error) {
            showAlert('Failed to delete character: ' + error.message, 'danger');
        }
    }
}

// URL sharing functions
function encodeCharacterToURL(character) {
    try {
        // Create a simplified character object for URL sharing
        const shareData = {
            n: character.name,
            r: character.race,
            c: character.class,
            b: character.background,
            l: character.level || 1,
            a: character.abilities,
            s: character.spells || []
        };
        
        // Convert to JSON and encode
        const jsonString = JSON.stringify(shareData);
        const encoded = btoa(jsonString);
        
        return encoded;
    } catch (error) {
        console.error('Failed to encode character:', error);
        return null;
    }
}

function decodeCharacterFromURL(encodedData) {
    try {
        // Decode and parse
        const jsonString = atob(encodedData);
        const shareData = JSON.parse(jsonString);
        
        // Convert back to full character object
        const character = {
            name: shareData.n,
            race: shareData.r,
            class: shareData.c,
            background: shareData.b,
            level: shareData.l || 1,
            abilities: shareData.a,
            spells: shareData.s || [],
            hitPoints: calculateHitPoints(shareData.c, shareData.a.constitution),
            armorClass: calculateArmorClass(shareData.a.dexterity),
            proficiencyBonus: 2, // Level 1
            equipment: getStartingEquipment(shareData.c, shareData.b),
            skills: [],
            savingThrows: getSavingThrowProficiencies(shareData.c),
            created: new Date().toISOString(),
            updated: new Date().toISOString(),
            shared: true // Mark as shared character
        };
        
        return character;
    } catch (error) {
        console.error('Failed to decode character:', error);
        return null;
    }
}

function generateShareURL(character) {
    const encoded = encodeCharacterToURL(character);
    if (!encoded) return null;
    
    const baseURL = window.location.origin + window.location.pathname.replace(/[^/]*$/, '');
    return `${baseURL}character.html?share=${encoded}`;
}

function copyShareURL(character) {
    const shareURL = generateShareURL(character);
    if (!shareURL) {
        showAlert('Failed to generate share URL', 'danger');
        return;
    }
    
    // Copy to clipboard
    if (navigator.clipboard) {
        navigator.clipboard.writeText(shareURL).then(() => {
            showAlert('Share URL copied to clipboard!', 'success');
        }).catch(() => {
            // Fallback for older browsers
            fallbackCopyToClipboard(shareURL);
        });
    } else {
        fallbackCopyToClipboard(shareURL);
    }
}

function fallbackCopyToClipboard(text) {
    const textArea = document.createElement('textarea');
    textArea.value = text;
    textArea.style.position = 'fixed';
    textArea.style.left = '-999999px';
    textArea.style.top = '-999999px';
    document.body.appendChild(textArea);
    textArea.focus();
    textArea.select();
    
    try {
        document.execCommand('copy');
        showAlert('Share URL copied to clipboard!', 'success');
    } catch (err) {
        showAlert('Failed to copy URL. Please copy manually: ' + text, 'warning');
    }
    
    document.body.removeChild(textArea);
}

// Character sharing function
function shareCharacter(id) {
    const character = getCharacter(id);
    if (!character) {
        showAlert('Character not found', 'danger');
        return;
    }
    
    copyShareURL(character);
}

// Export global functions
window.ABILITY_GENERATION_METHODS = ABILITY_GENERATION_METHODS;
window.generateAbilityScores = generateAbilityScores;
window.createCharacter = createCharacter;
window.validateCharacterForm = validateCharacterForm;
window.validateSpellSelection = validateSpellSelection;
window.getAbilityModifier = getAbilityModifier;
window.formatModifier = formatModifier;
window.showAlert = showAlert;
window.showLoading = showLoading;
window.formatDate = formatDate;
window.getCharacterSummary = getCharacterSummary;
window.renderCharacterCard = renderCharacterCard;
window.viewCharacter = viewCharacter;
window.deleteCharacterConfirm = deleteCharacterConfirm;
window.shareCharacter = shareCharacter;
window.encodeCharacterToURL = encodeCharacterToURL;
window.decodeCharacterFromURL = decodeCharacterFromURL;
window.generateShareURL = generateShareURL;
window.copyShareURL = copyShareURL;
