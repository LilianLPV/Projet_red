package game

import (
	"Projet-Red/feature"
	"Projet-Red/models"
	"Projet-Red/utils"
	"fmt"
	"math/rand/v2"
)

// Fonction du combat tour par tour
func Combat(p *models.Player, e models.Monster) bool {
	utils.Clear()
	feature.AfficherCombat()
	fmt.Println("\nAppuie sur Entrée pour engager le combat...")
	utils.ReadInput()

	cs := &models.CombatState{Tour: 1}
	joueurRapide := p.Speed >= e.Speed

	for p.HP > 0 && e.HP > 0 {
		fmt.Println()
		fmt.Println(utils.Cyan+"=== Tour", cs.Tour, "==="+utils.Reset)

		if joueurRapide {
			if TourJoueur(p, &e, cs) {
				fmt.Println("Vous prenez la fuite !")
				return false
			}
			if e.HP <= 0 {
				handleVictoire(p, &e)
				return true
			}
			TourEnnemi(p, &e, cs)
			if p.HP <= 0 {
				handleDefaite(p)
				return false
			}
		} else {
			TourEnnemi(p, &e, cs)
			if p.HP <= 0 {
				handleDefaite(p)
				return false
			}
			if TourJoueur(p, &e, cs) {
				fmt.Println("Vous prenez la fuite !")
				return false
			}
			if e.HP <= 0 {
				handleVictoire(p, &e)
				return true
			}
		}
		cs.Tour++
	}
	return false
}

// Effet du joueur en combat
func TourJoueur(p *models.Player, e *models.Monster, cs *models.CombatState) bool {
	if cs.JoueurEtourdi {
		fmt.Println("Vous êtes étourdi et passez votre tour !")
		cs.JoueurEtourdi = false
	} else {
		if CharTurn(p, e, cs) {
			fmt.Println("Vous prenez la fuite !")
			fmt.Println()
			return true
		}
		if utils.Chance(p.StunChance) {
			cs.EnnemiEtourdi = true
			fmt.Println("Vous étourdissez", e.Name, "!")
		}
	}
	return false
}

// Effet de l'ennemie en combat
func TourEnnemi(p *models.Player, e *models.Monster, cs *models.CombatState) {
	if cs.EnnemiEmpoisonne > 0 {
		degatsPoison := 5
		e.HP -= degatsPoison
		cs.EnnemiEmpoisonne--
		fmt.Println(e.Name, "subit", degatsPoison, "dégâts de poison ! PV :", e.HP)
		if e.HP <= 0 {
			return
		}
	}
	if cs.EnnemiEtourdi {
		fmt.Println(e.Name, "est étourdi et passe son tour !")
		cs.EnnemiEtourdi = false
	} else {
		if utils.Chance(p.Dodge) {
			fmt.Println("Vous esquivez l'attaque de", e.Name, "!")
		} else {
			degats := e.Damage
			if cs.Tour%5 == 0 {
				degats = e.Damage * 2
				fmt.Println(e.Name, "déchaîne une attaque critique !")
			}
			p.HP -= degats
			fmt.Println(e.Name, "vous inflige", degats, "dégâts. PV:", p.HP)
			if utils.Chance(e.StunChance) {
				cs.JoueurEtourdi = true
				fmt.Println(e.Name, "vous étourdit !")
			}
		}
	}
}

// Combat tour par tour

func CharTurn(p *models.Player, e *models.Monster, cs *models.CombatState) bool {
	actionFaite := false
	for !actionFaite {
		fmt.Println()
		fmt.Println("=== COMBAT ===")
		fmt.Println("Vous :", p.HP, "/", p.HPMax, "PV  |  Mana:", p.Mana, "/", p.ManaMax)
		fmt.Println(e.Name, ":", e.HP, "/", e.HPMax, "PV")
		fmt.Println(" ")
		fmt.Println("C'est à vous de combattre !")
		fmt.Println(" ")
		fmt.Println("\033[96m1.\033[0m Attaquer", p.BaseAttack)
		if len(p.Skill) > 0 {
			fmt.Println("\033[96m2.\033[0m Attaque spéciale :", p.Skill[0])
		} else {
			fmt.Println("\033[96m2.\033[0m Attaque spéciale : (non achetée)")
		}
		fmt.Println("\033[96m3.\033[0m Inventaire")

		fmt.Println("\033[96m4.\033[0m Fuir le combat")

		fmt.Print("Choisissez une option: ")
		choix := utils.ReadChoice("Ton choix : ", []string{"1", "2", "3", "4"})
		switch choix {
		case "1":
			utils.Clear()
			if utils.Chance(e.Dodge) {
				fmt.Println(e.Name, "esquive votre attaque !")
			} else {
				e.HP -= p.Damage
				fmt.Println("Vous infligez", p.Damage, "dégâts !")
			}
			actionFaite = true
		case "2":
			utils.Clear()
			if len(p.Skill) == 0 {
				fmt.Println("Tu n'as pas d'attaque spéciale ! Achète-la chez le marchand.")
				continue
			}
			attaque := models.AttackDefs[p.Skill[0]]
			if p.Mana < attaque.ManaCost {
				fmt.Println("Pas assez de mana ! (", p.Mana, "/", attaque.ManaCost, ")")
				continue
			}
			p.Mana -= attaque.ManaCost
			if utils.Chance(e.Dodge) {
				fmt.Println(e.Name, "esquive votre attaque !")
			} else {
				e.HP -= attaque.Damage
				fmt.Println("Vous lancez", p.Skill[0], "et infligez", attaque.Damage, "dégâts ! (-", attaque.ManaCost, "mana)")
			}
			actionFaite = true
		case "3":
			utils.Clear()
			fmt.Println("--- Inventaire ---")
			fmt.Println("1) Potion de santé (x", feature.QuantiteItem(p, "Potion de santé"), ")")
			fmt.Println("2) Potion de mana (x", feature.QuantiteItem(p, "Potion de mana"), ")")
			fmt.Println("3) Potion de poison (x", feature.QuantiteItem(p, "Potion de poison"), ")")
			fmt.Println("4) Retour")
			sousChoix := utils.ReadChoice("Ton choix : ", []string{"1", "2", "3", "4"})
			switch sousChoix {
			case "1":
				if feature.QuantiteItem(p, "Potion de santé") > 0 {
					feature.TakePot(p)
					actionFaite = true
				} else {
					fmt.Println("Tu n'as pas de potion de santé !")
				}
			case "2":
				if feature.QuantiteItem(p, "Potion de mana") > 0 {
					feature.ManaPot(p)
					actionFaite = true
				} else {
					fmt.Println("Tu n'as pas de potion de mana !")
				}
			case "3":
				if feature.QuantiteItem(p, "Potion de poison") > 0 {
					if feature.PoisonPot(p, e, cs) {
						actionFaite = true
					}
				} else {
					fmt.Println("Tu n'as pas de potion de poison !")
				}
			case "4":
				utils.Clear()
			}
		case "4":
			utils.Clear()
			return true
		}
	}
	return false
}

//Le nombre d'xp gagné et les stats qu'il gagne a chaque level

const Levelmax = 5

func XPWin(p *models.Player, montant int) {
	if p.Level >= Levelmax {
		return
	}
	p.XP += montant
	fmt.Println("Vous gagnez", montant, "XP !")
	if p.XP >= p.XPMax {
		p.XP -= p.XPMax
		p.Level++
		p.XPMax += 50
		p.HPMax += 10
		p.HP = p.HPMax
		p.Damage += 2
		if p.Level >= Levelmax {
			p.XP = 0 
			fmt.Println("NIVEAU SUPÉRIEUR ! Vous êtes niveau", p.Level, "(niveau maximum atteint)")
		} else {
			p.XPMax += 50
			fmt.Println("NIVEAU SUPÉRIEUR ! Vous êtes niveau", p.Level)
		}
	}
}

//En cas de victoire

func handleVictoire(p *models.Player, e *models.Monster) {
	or := 15 + rand.IntN(16)
	p.Money += or
	fmt.Println("Vous avez vaincu", e.Name, "et gagné", or, "PO !")
	XPWin(p, e.XPReward)
}

// En cas de défaite
func handleDefaite(p *models.Player) {
	if p.DejaRessuscite {
		fmt.Println("Vous succombez définitivement...")
		return
	}
	p.HP = p.HPMax / 2
	p.Mana = p.ManaMax
	p.DejaRessuscite = true
	fmt.Println("Vous êtes tombé... mais vous vous relevez une dernière fois avec", p.HP, "PV !")
}
