// Local Storage Management for D&D Character Creator

const STORAGE_KEY = 'dnd_characters';

// Character storage functions
function getCharacters() {
    try {
        const stored = localStorage.getItem(STORAGE_KEY);
        return stored ? JSON.parse(stored) : [];
    } catch (error) {
        console.error('Error loading characters:', error);
        return [];
    }
}

function saveCharacter(character) {
    try {
        const characters = getCharacters();
        
        // Generate unique ID
        const id = Date.now() + Math.random();
        character.id = id;
        character.created = new Date().toISOString();
        character.updated = new Date().toISOString();
        
        characters.push(character);
        localStorage.setItem(STORAGE_KEY, JSON.stringify(characters));
        
        return id;
    } catch (error) {
        console.error('Error saving character:', error);
        throw new Error('Failed to save character. Storage may be full.');
    }
}

function updateCharacter(id, updatedCharacter) {
    try {
        const characters = getCharacters();
        const index = characters.findIndex(char => char.id === id);
        
        if (index === -1) {
            throw new Error('Character not found');
        }
        
        updatedCharacter.id = id;
        updatedCharacter.updated = new Date().toISOString();
        // Preserve creation date
        updatedCharacter.created = characters[index].created;
        
        characters[index] = updatedCharacter;
        localStorage.setItem(STORAGE_KEY, JSON.stringify(characters));
        
        return true;
    } catch (error) {
        console.error('Error updating character:', error);
        throw new Error('Failed to update character');
    }
}

function deleteCharacter(id) {
    try {
        const characters = getCharacters();
        const filteredCharacters = characters.filter(char => char.id !== id);
        
        if (filteredCharacters.length === characters.length) {
            throw new Error('Character not found');
        }
        
        localStorage.setItem(STORAGE_KEY, JSON.stringify(filteredCharacters));
        return true;
    } catch (error) {
        console.error('Error deleting character:', error);
        throw new Error('Failed to delete character');
    }
}

function getCharacter(id) {
    try {
        const characters = getCharacters();
        return characters.find(char => char.id === id) || null;
    } catch (error) {
        console.error('Error getting character:', error);
        return null;
    }
}

function exportCharacter(id) {
    try {
        const character = getCharacter(id);
        if (!character) {
            throw new Error('Character not found');
        }
        
        const dataStr = JSON.stringify(character, null, 2);
        const dataBlob = new Blob([dataStr], {type: 'application/json'});
        
        const link = document.createElement('a');
        link.href = URL.createObjectURL(dataBlob);
        link.download = `${character.name.replace(/[^a-z0-9]/gi, '_').toLowerCase()}-character.json`;
        link.click();
        
        return true;
    } catch (error) {
        console.error('Error exporting character:', error);
        throw new Error('Failed to export character');
    }
}

function importCharacter(characterData) {
    try {
        // Validate character data
        if (!characterData.name || !characterData.race || !characterData.class) {
            throw new Error('Invalid character data');
        }
        
        // Remove ID to generate new one
        delete characterData.id;
        
        return saveCharacter(characterData);
    } catch (error) {
        console.error('Error importing character:', error);
        throw new Error('Failed to import character');
    }
}

// Storage utility functions
function getStorageInfo() {
    try {
        const characters = getCharacters();
        const dataSize = new Blob([localStorage.getItem(STORAGE_KEY) || '']).size;
        
        return {
            characterCount: characters.length,
            storageUsed: dataSize,
            storageUsedFormatted: formatBytes(dataSize),
            lastUpdated: characters.length > 0 ? 
                Math.max(...characters.map(c => new Date(c.updated || c.created).getTime())) : null
        };
    } catch (error) {
        console.error('Error getting storage info:', error);
        return {
            characterCount: 0,
            storageUsed: 0,
            storageUsedFormatted: '0 B',
            lastUpdated: null
        };
    }
}

function formatBytes(bytes, decimals = 2) {
    if (bytes === 0) return '0 B';
    
    const k = 1024;
    const dm = decimals < 0 ? 0 : decimals;
    const sizes = ['B', 'KB', 'MB', 'GB'];
    
    const i = Math.floor(Math.log(bytes) / Math.log(k));
    
    return parseFloat((bytes / Math.pow(k, i)).toFixed(dm)) + ' ' + sizes[i];
}

function clearAllCharacters() {
    try {
        localStorage.removeItem(STORAGE_KEY);
        return true;
    } catch (error) {
        console.error('Error clearing characters:', error);
        throw new Error('Failed to clear character data');
    }
}

// Backup and restore functions
function createBackup() {
    try {
        const characters = getCharacters();
        const backup = {
            version: '1.0',
            created: new Date().toISOString(),
            characters: characters
        };
        
        const dataStr = JSON.stringify(backup, null, 2);
        const dataBlob = new Blob([dataStr], {type: 'application/json'});
        
        const link = document.createElement('a');
        link.href = URL.createObjectURL(dataBlob);
        link.download = `dnd-characters-backup-${new Date().toISOString().split('T')[0]}.json`;
        link.click();
        
        return true;
    } catch (error) {
        console.error('Error creating backup:', error);
        throw new Error('Failed to create backup');
    }
}

function restoreBackup(backupData) {
    try {
        // Validate backup format
        if (!backupData.characters || !Array.isArray(backupData.characters)) {
            throw new Error('Invalid backup format');
        }
        
        // Merge with existing characters (don't overwrite)
        const existingCharacters = getCharacters();
        const restoredCharacters = backupData.characters.map(char => {
            // Generate new IDs to avoid conflicts
            delete char.id;
            char.id = Date.now() + Math.random();
            char.restored = new Date().toISOString();
            return char;
        });
        
        const allCharacters = [...existingCharacters, ...restoredCharacters];
        localStorage.setItem(STORAGE_KEY, JSON.stringify(allCharacters));
        
        return restoredCharacters.length;
    } catch (error) {
        console.error('Error restoring backup:', error);
        throw new Error('Failed to restore backup');
    }
}

// Search and filter functions
function searchCharacters(query) {
    try {
        const characters = getCharacters();
        const lowercaseQuery = query.toLowerCase();
        
        return characters.filter(character => 
            character.name.toLowerCase().includes(lowercaseQuery) ||
            character.race.toLowerCase().includes(lowercaseQuery) ||
            character.class.toLowerCase().includes(lowercaseQuery) ||
            character.background.toLowerCase().includes(lowercaseQuery)
        );
    } catch (error) {
        console.error('Error searching characters:', error);
        return [];
    }
}

function filterCharacters(filters) {
    try {
        const characters = getCharacters();
        
        return characters.filter(character => {
            if (filters.race && character.race !== filters.race) return false;
            if (filters.class && character.class !== filters.class) return false;
            if (filters.background && character.background !== filters.background) return false;
            if (filters.level && character.level !== filters.level) return false;
            return true;
        });
    } catch (error) {
        console.error('Error filtering characters:', error);
        return [];
    }
}

// Statistics functions
function getCharacterStats() {
    try {
        const characters = getCharacters();
        
        const stats = {
            total: characters.length,
            byRace: {},
            byClass: {},
            byBackground: {},
            byLevel: {},
            averageLevel: 0
        };
        
        characters.forEach(character => {
            // Count by race
            stats.byRace[character.race] = (stats.byRace[character.race] || 0) + 1;
            
            // Count by class
            stats.byClass[character.class] = (stats.byClass[character.class] || 0) + 1;
            
            // Count by background
            stats.byBackground[character.background] = (stats.byBackground[character.background] || 0) + 1;
            
            // Count by level
            const level = character.level || 1;
            stats.byLevel[level] = (stats.byLevel[level] || 0) + 1;
        });
        
        // Calculate average level
        if (characters.length > 0) {
            const totalLevels = characters.reduce((sum, char) => sum + (char.level || 1), 0);
            stats.averageLevel = Math.round((totalLevels / characters.length) * 10) / 10;
        }
        
        return stats;
    } catch (error) {
        console.error('Error getting character stats:', error);
        return {
            total: 0,
            byRace: {},
            byClass: {},
            byBackground: {},
            byLevel: {},
            averageLevel: 0
        };
    }
}

// Export all functions for use in other scripts
window.CharacterStorage = {
    getCharacters,
    saveCharacter,
    updateCharacter,
    deleteCharacter,
    getCharacter,
    exportCharacter,
    importCharacter,
    getStorageInfo,
    clearAllCharacters,
    createBackup,
    restoreBackup,
    searchCharacters,
    filterCharacters,
    getCharacterStats,
    formatBytes
};
