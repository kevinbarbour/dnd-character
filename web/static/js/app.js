// D&D Character Creator JavaScript

document.addEventListener('DOMContentLoaded', function() {
    // Initialize character creation form if it exists
    if (document.getElementById('characterForm')) {
        initCharacterCreation();
    }
});

function initCharacterCreation() {
    let selectedMethod = null;
    
    // Method selection
    const methodCards = document.querySelectorAll('.method-card');
    const generateBtn = document.getElementById('generateBtn');
    const abilityScoresDiv = document.getElementById('abilityScores');
    
    methodCards.forEach(card => {
        card.addEventListener('click', function() {
            // Remove selected class from all cards
            methodCards.forEach(c => c.classList.remove('selected'));
            
            // Add selected class to clicked card
            this.classList.add('selected');
            
            // Store selected method
            selectedMethod = this.dataset.method;
            
            // Enable generate button
            generateBtn.disabled = false;
        });
    });
    
    // Generate ability scores
    generateBtn.addEventListener('click', function() {
        if (!selectedMethod) return;
        
        // Show loading state
        this.innerHTML = '<span class="spinner-border spinner-border-sm me-2"></span>Generating...';
        this.disabled = true;
        
        // Make API call to generate abilities
        fetch('/api/generate-abilities', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ method: selectedMethod })
        })
        .then(response => response.json())
        .then(data => {
            // Populate ability score inputs
            document.getElementById('strength').value = data.Strength;
            document.getElementById('dexterity').value = data.Dexterity;
            document.getElementById('constitution').value = data.Constitution;
            document.getElementById('intelligence').value = data.Intelligence;
            document.getElementById('wisdom').value = data.Wisdom;
            document.getElementById('charisma').value = data.Charisma;
            
            // Show ability scores section
            abilityScoresDiv.style.display = 'block';
            abilityScoresDiv.classList.add('fade-in');
            
            // Reset button
            this.innerHTML = '<i class="bi bi-arrow-clockwise"></i> Generate New Scores';
            this.disabled = false;
            
            // Update form validation
            updateFormValidation();
        })
        .catch(error => {
            console.error('Error generating abilities:', error);
            alert('Failed to generate ability scores. Please try again.');
            
            // Reset button
            this.innerHTML = '<i class="bi bi-arrow-clockwise"></i> Generate Ability Scores';
            this.disabled = false;
        });
    });
    
    // Race selection info
    const raceSelect = document.getElementById('race');
    const raceInfo = document.getElementById('raceInfo');
    
    raceSelect.addEventListener('change', function() {
        const raceName = this.value;
        if (raceName && window.raceData && window.raceData[raceName]) {
            const race = window.raceData[raceName];
            const bonuses = [];
            
            Object.entries(race.bonuses).forEach(([ability, bonus]) => {
                if (bonus > 0) {
                    bonuses.push(`+${bonus} ${ability.charAt(0).toUpperCase() + ability.slice(1)}`);
                }
            });
            
            raceInfo.innerHTML = `
                <strong>Racial Bonuses:</strong> ${bonuses.join(', ')}<br>
                <strong>Size:</strong> ${race.size} | <strong>Speed:</strong> ${race.speed} ft<br>
                <strong>Traits:</strong> ${race.traits.join(', ')}
            `;
        } else {
            raceInfo.innerHTML = '';
        }
        
        updateFormValidation();
        updateCharacterPreview();
    });
    
    // Class selection info
    const classSelect = document.getElementById('class');
    const classInfo = document.getElementById('classInfo');
    
    classSelect.addEventListener('change', function() {
        const className = this.value;
        if (className && window.classData && window.classData[className]) {
            const classData = window.classData[className];
            
            classInfo.innerHTML = `
                <strong>Hit Die:</strong> d${classData.hitDie} | <strong>Primary Ability:</strong> ${classData.primaryAbility}<br>
                <strong>Saving Throws:</strong> ${classData.savingThrows.join(', ')}
            `;
        } else {
            classInfo.innerHTML = '';
        }
        
        updateFormValidation();
        updateCharacterPreview();
    });
    
    // Background selection info
    const backgroundSelect = document.getElementById('background');
    const backgroundInfo = document.getElementById('backgroundInfo');
    
    backgroundSelect.addEventListener('change', function() {
        const backgroundName = this.value;
        if (backgroundName && window.backgroundData && window.backgroundData[backgroundName]) {
            const background = window.backgroundData[backgroundName];
            
            backgroundInfo.innerHTML = `
                <strong>Skills:</strong> ${background.skills.join(', ')}<br>
                <strong>Feature:</strong> ${background.feature}
            `;
        } else {
            backgroundInfo.innerHTML = '';
        }
        
        updateFormValidation();
        updateCharacterPreview();
    });
    
    // Character name input
    const nameInput = document.getElementById('name');
    nameInput.addEventListener('input', function() {
        updateFormValidation();
        updateCharacterPreview();
    });
    
    // Ability score inputs
    const abilityInputs = document.querySelectorAll('#abilityScores input');
    abilityInputs.forEach(input => {
        input.addEventListener('input', function() {
            updateFormValidation();
            updateCharacterPreview();
        });
    });
    
    function updateFormValidation() {
        const form = document.getElementById('characterForm');
        const submitBtn = document.getElementById('submitBtn');
        
        // Check if all required fields are filled
        const isValid = form.checkValidity();
        submitBtn.disabled = !isValid;
    }
    
    function updateCharacterPreview() {
        const name = document.getElementById('name').value;
        const race = document.getElementById('race').value;
        const className = document.getElementById('class').value;
        const background = document.getElementById('background').value;
        
        const strength = parseInt(document.getElementById('strength').value) || 0;
        const dexterity = parseInt(document.getElementById('dexterity').value) || 0;
        const constitution = parseInt(document.getElementById('constitution').value) || 0;
        const intelligence = parseInt(document.getElementById('intelligence').value) || 0;
        const wisdom = parseInt(document.getElementById('wisdom').value) || 0;
        const charisma = parseInt(document.getElementById('charisma').value) || 0;
        
        if (name && race && className && background && strength && dexterity && constitution && intelligence && wisdom && charisma) {
            // Calculate modifiers
            const getModifier = (score) => Math.floor((score - 10) / 2);
            
            // Apply racial bonuses for preview
            let finalAbilities = {
                strength: strength,
                dexterity: dexterity,
                constitution: constitution,
                intelligence: intelligence,
                wisdom: wisdom,
                charisma: charisma
            };
            
            if (window.raceData && window.raceData[race]) {
                const raceBonuses = window.raceData[race].bonuses;
                finalAbilities.strength += raceBonuses.strength;
                finalAbilities.dexterity += raceBonuses.dexterity;
                finalAbilities.constitution += raceBonuses.constitution;
                finalAbilities.intelligence += raceBonuses.intelligence;
                finalAbilities.wisdom += raceBonuses.wisdom;
                finalAbilities.charisma += raceBonuses.charisma;
            }
            
            // Calculate derived stats
            const hitDie = window.classData && window.classData[className] ? window.classData[className].hitDie : 8;
            const hitPoints = hitDie + getModifier(finalAbilities.constitution);
            const armorClass = 10 + getModifier(finalAbilities.dexterity);
            
            const previewContent = `
                <div class="row">
                    <div class="col-md-6">
                        <h6><i class="bi bi-person"></i> ${name}</h6>
                        <p class="mb-1">Level 1 ${race} ${className}</p>
                        <p class="text-muted small">${background}</p>
                    </div>
                    <div class="col-md-6">
                        <div class="row g-2">
                            <div class="col-6">
                                <small class="text-muted">Hit Points</small>
                                <div class="fw-bold text-success">${hitPoints}</div>
                            </div>
                            <div class="col-6">
                                <small class="text-muted">Armor Class</small>
                                <div class="fw-bold text-info">${armorClass}</div>
                            </div>
                        </div>
                    </div>
                </div>
                <hr>
                <div class="row g-2">
                    <div class="col-2 text-center">
                        <small class="text-muted d-block">STR</small>
                        <span class="badge bg-primary">${finalAbilities.strength}</span>
                        <small class="d-block">${getModifier(finalAbilities.strength) >= 0 ? '+' : ''}${getModifier(finalAbilities.strength)}</small>
                    </div>
                    <div class="col-2 text-center">
                        <small class="text-muted d-block">DEX</small>
                        <span class="badge bg-success">${finalAbilities.dexterity}</span>
                        <small class="d-block">${getModifier(finalAbilities.dexterity) >= 0 ? '+' : ''}${getModifier(finalAbilities.dexterity)}</small>
                    </div>
                    <div class="col-2 text-center">
                        <small class="text-muted d-block">CON</small>
                        <span class="badge bg-danger">${finalAbilities.constitution}</span>
                        <small class="d-block">${getModifier(finalAbilities.constitution) >= 0 ? '+' : ''}${getModifier(finalAbilities.constitution)}</small>
                    </div>
                    <div class="col-2 text-center">
                        <small class="text-muted d-block">INT</small>
                        <span class="badge bg-info">${finalAbilities.intelligence}</span>
                        <small class="d-block">${getModifier(finalAbilities.intelligence) >= 0 ? '+' : ''}${getModifier(finalAbilities.intelligence)}</small>
                    </div>
                    <div class="col-2 text-center">
                        <small class="text-muted d-block">WIS</small>
                        <span class="badge bg-warning">${finalAbilities.wisdom}</span>
                        <small class="d-block">${getModifier(finalAbilities.wisdom) >= 0 ? '+' : ''}${getModifier(finalAbilities.wisdom)}</small>
                    </div>
                    <div class="col-2 text-center">
                        <small class="text-muted d-block">CHA</small>
                        <span class="badge bg-secondary">${finalAbilities.charisma}</span>
                        <small class="d-block">${getModifier(finalAbilities.charisma) >= 0 ? '+' : ''}${getModifier(finalAbilities.charisma)}</small>
                    </div>
                </div>
            `;
            
            document.getElementById('previewContent').innerHTML = previewContent;
            document.getElementById('characterPreview').style.display = 'block';
            document.getElementById('characterPreview').classList.add('fade-in');
        } else {
            document.getElementById('characterPreview').style.display = 'none';
        }
    }
}

// Utility functions
function showToast(message, type = 'info') {
    // Create toast element
    const toast = document.createElement('div');
    toast.className = `alert alert-${type} alert-dismissible fade show position-fixed`;
    toast.style.top = '20px';
    toast.style.right = '20px';
    toast.style.zIndex = '9999';
    toast.innerHTML = `
        ${message}
        <button type="button" class="btn-close" data-bs-dismiss="alert"></button>
    `;
    
    document.body.appendChild(toast);
    
    // Auto-remove after 5 seconds
    setTimeout(() => {
        if (toast.parentNode) {
            toast.parentNode.removeChild(toast);
        }
    }, 5000);
}

// Form submission handling
document.addEventListener('submit', function(e) {
    const form = e.target;
    if (form.id === 'characterForm') {
        const submitBtn = form.querySelector('button[type="submit"]');
        if (submitBtn) {
            submitBtn.innerHTML = '<span class="spinner-border spinner-border-sm me-2"></span>Creating Character...';
            submitBtn.disabled = true;
        }
    }
});

// Character card animations
document.addEventListener('DOMContentLoaded', function() {
    const cards = document.querySelectorAll('.character-card');
    cards.forEach((card, index) => {
        card.style.animationDelay = `${index * 0.1}s`;
        card.classList.add('fade-in');
    });
});

// Confirmation dialogs
function confirmDelete(characterName) {
    return confirm(`Are you sure you want to delete "${characterName}"? This action cannot be undone.`);
}

// Keyboard navigation for method cards
document.addEventListener('keydown', function(e) {
    if (e.target.classList.contains('method-card')) {
        if (e.key === 'Enter' || e.key === ' ') {
            e.preventDefault();
            e.target.click();
        }
    }
});

// Add accessibility attributes
document.addEventListener('DOMContentLoaded', function() {
    const methodCards = document.querySelectorAll('.method-card');
    methodCards.forEach(card => {
        card.setAttribute('tabindex', '0');
        card.setAttribute('role', 'button');
        card.setAttribute('aria-label', `Select ${card.querySelector('h6').textContent} method`);
    });
});
