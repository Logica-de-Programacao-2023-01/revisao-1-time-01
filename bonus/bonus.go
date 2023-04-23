package bonus

import "fmt"

func CalculateDamage(characterLevel, enemyLevel int) (int, error) {
	var dano int

	if characterLevel <= 0 || enemyLevel <= 0 {
		return 0, fmt.Errorf("nível inválido")
	}

	if characterLevel >= 100 {
		dano = characterLevel * 20
		if characterLevel > enemyLevel+20 {
			dano *= 5
		} else if characterLevel < enemyLevel+20 {
			dano *= 2
		}
	} else if enemyLevel >= 100 {
		dano = characterLevel * 2
		if characterLevel > enemyLevel+20 {
			dano *= 5
		} else if characterLevel < enemyLevel+20 {
			dano *= 2
		}
	} else if characterLevel > enemyLevel {
		dano = characterLevel * 10
	} else if characterLevel < enemyLevel {
		dano = characterLevel * 5
	} else if characterLevel == enemyLevel {
		dano = characterLevel * 7
	}

	return dano, nil
}
