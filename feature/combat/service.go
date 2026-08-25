package combat

import (
	"math/rand/v2"

	"Projet-Red/feature/character"
	"Projet-Red/utils"
)

func Combat(p *character.Player, e Monster) bool {
	utils.Clear()
	DisplayCombatIntro()
	PromptEngageCombat()

	cs := &CombatState{Tour: 1}
	playerSpeed := p.Speed >= e.Speed

	for p.HP > 0 && e.HP > 0 {
		DisplayTourHeader(cs.Tour)

		if playerSpeed {
			if playerTurn(p, &e, cs) {
				DisplayFlee()
				return false
			}
			if e.HP <= 0 {
				handleVictory(p, &e)
				return true
			}
			enemyTurn(p, &e, cs)
			if p.HP <= 0 {
				handleDefeat(p)
				return false
			}
		} else {
			enemyTurn(p, &e, cs)
			if p.HP <= 0 {
				handleDefeat(p)
				return false
			}
			if playerTurn(p, &e, cs) {
				DisplayFlee()
				return false
			}
			if e.HP <= 0 {
				handleVictory(p, &e)
				return true
			}
		}
		cs.Tour++
	}
	return false
}

// playerTurn gère le tour du joueur
func playerTurn(p *character.Player, e *Monster, cs *CombatState) bool {
	if cs.PlayerStun {
		DisplayPlayerStunnedSkip()
		cs.PlayerStun = false
		return false
	}
	if charTurn(p, e, cs) {
		return true
	}
	if utils.Chance(p.StunChance) {
		cs.EnemyStun = true
		DisplayPlayerStuns(e.Name)
	}
	return false
}

// enemyTurn gère le tour de l'ennemi
func enemyTurn(p *character.Player, e *Monster, cs *CombatState) {
	if cs.EnemyPoisons > 0 {
		damagePoison := 5
		e.HP -= damagePoison
		cs.EnemyPoisons--
		DisplayEnemyPoisonTick(e.Name, damagePoison, e.HP)
		if e.HP <= 0 {
			return
		}
	}
	if cs.EnemyStun {
		DisplayEnemyStunnedSkip(e.Name)
		cs.EnemyStun = false
		return
	}
	if utils.Chance(p.Dodge) {
		DisplayPlayerDodges(e.Name)
		return
	}
	damage := e.Damage
	if cs.Tour%5 == 0 {
		damage = e.Damage * 2
		DisplayEnemyCritical(e.Name)
	}
	p.HP -= damage
	DisplayEnemyHit(e.Name, damage, p.HP)
	if utils.Chance(e.StunChance) {
		cs.PlayerStun = true
		DisplayEnemyStuns(e.Name)
	}
}

func charTurn(p *character.Player, e *Monster, cs *CombatState) bool {
	for {
		DisplayCombatStatus(p, e)
		DisplayPlayerMenu(p)
		choix := ReadPlayerChoice()
		switch choix {
		case "1":
			utils.Clear()
			attackDeBase(p, e)
			return false
		case "2":
			utils.Clear()
			if attackSpeciale(p, e) {
				return false
			}
		case "3":
			utils.Clear()
			if inventoryCombat(p, e, cs) {
				return false
			}
		case "4":
			utils.Clear()
			return true
		}
	}
}

// attaqueDeBase applique l'attaque de base du joueur sur l'ennemi.
func attackDeBase(p *character.Player, e *Monster) {
	if utils.Chance(e.Dodge) {
		DisplayEnemyDodges(e.Name)
		return
	}
	e.HP -= p.Damage
	DisplayPlayerBaseAttackHit(p.Damage)
}

// attaqueSpeciale applique l'attaque spéciale du joueur (si achetée et si assez de mana).
func attackSpeciale(p *character.Player, e *Monster) bool {
	if len(p.Skill) == 0 {
		DisplaySkillNotBought()
		return false
	}
	attaque := AttackDefs[p.Skill[0]]
	if p.Mana < attaque.ManaCost {
		DisplayNotEnoughMana(p.Mana, attaque.ManaCost)
		return false
	}
	p.Mana -= attaque.ManaCost
	if utils.Chance(e.Dodge) {
		DisplayEnemyDodges(e.Name)
	} else {
		e.HP -= attaque.Damage
		DisplaySkillHit(p.Skill[0], attaque.Damage, attaque.ManaCost)
	}
	return true
}

// inventaireCombat affiche le sous-menu des potions utilisables en combat et applique le choix.
func inventoryCombat(p *character.Player, e *Monster, cs *CombatState) bool {
	utils.Clear()
	DisplayInventoryMenu(p)
	subChoice := ReadInventoryChoice()
	switch subChoice {
	case "1":
		if character.QuantityOf(p, "Potion de santé") > 0 {
			character.DrinkHealthPotion(p)
			return true
		}
		DisplayNoHealthPotion()
	case "2":
		if character.QuantityOf(p, "Potion de mana") > 0 {
			character.DrinkManaPotion(p)
			return true
		}
		DisplayNoManaPotion()
	case "3":
		if character.QuantityOf(p, "Potion de poison") > 0 {
			return PoisonPot(p, e, cs)
		}
		DisplayNoPoisonPotion()
	case "4":
		utils.Clear()
	}
	return false
}

// PoisonPot jette une potion de poison sur l'ennemi.
func PoisonPot(p *character.Player, e *Monster, cs *CombatState) bool {
	if character.QuantityOf(p, "Potion de poison") == 0 {
		DisplayNoPoisonPotion()
		return false
	}
	if cs.EnemyPoisons > 0 {
		DisplayAlreadyPoisoned(e.Name)
		return false
	}
	cs.EnemyPoisons = 5
	DisplayPoisonThrown(e.Name)
	character.RemoveItem(p, "Potion de poison", 1)
	return true
}

// handleVictoire récompense le joueur après avoir vaincu son adversaire.
func handleVictory(p *character.Player, e *Monster) {
	or := 15 + rand.IntN(16)
	p.Money += or
	DisplayVictory(e.Name, or)
	character.XPWin(p, e.XPReward)
}

// handleDefaite gère la défaite du joueur : une résurrection unique à mi-vie, sinon la mort.
func handleDefeat(p *character.Player) {
	if p.DejaRessuscite {
		DisplayDefeatFinal()
		return
	}
	p.HP = p.HPMax / 2
	p.Mana = p.ManaMax
	p.DejaRessuscite = true
	DisplayDefeatRevive(p.HP)
}
