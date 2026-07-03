package feature

import (
	"Projet-Red/models"
	"Projet-Red/utils"
	"fmt"
)

// Prendre une potion de vie + avertissement si elle regen trop 

func TakePot(p *models.Player) {
	for _, o := range p.Inventory {
		if o.Name == "Potion de santé" {

			if p.HP > p.HPMax-50 {
				fmt.Println("Vous avez actuellement", p.HP, "/", p.HPMax, "PV")
				fmt.Println("Il vous manque moins de 50 PV, vous allez gaspiller la potion.")
				fmt.Println("1: Oui   2: Retour")
				if utils.ReadChoice("Ton choix : ", []string{"1", "2"}) == "2" {
					return
				}
			}

			p.HP = utils.Clamp(p.HP+50, 0, p.HPMax)
			RemoveItem(p, "Potion de santé", 1)
			fmt.Println("Vous avez bu une potion de santé. PV :", p.HP, "/", p.HPMax)
			return
		}
	}

	fmt.Println("Vous n'avez pas de potion de santé dans votre inventaire")
}

// Potion de poison avec tic et verification qu'il n'y ai pas plusieurs potion de poison lancé en meme temps car inutile

func PoisonPot(p *models.Player, e *models.Monster, cs *models.CombatState) bool {
	for _, o := range p.Inventory {
		if o.Name == "Potion de poison" {
			if cs.EnnemiEmpoisonne > 0 {
				fmt.Println(e.Name, "est déjà empoisonné ! Impossible de cumuler.")
				return false
			}
			cs.EnnemiEmpoisonne = 5
			fmt.Println("Vous jetez la potion sur", e.Name, "!")
			RemoveItem(p, "Potion de poison", 1)
			return true
		}
	}
	fmt.Println("Vous n'avez pas de potion de poison dans votre inventaire")
	return false
}

// prendre une potion de mana + avertissement si elle regen trop 

func ManaPot(p *models.Player) {
	for _, o := range p.Inventory {
		if o.Name == "Potion de mana" {

			if p.Mana > p.ManaMax-20 {
				fmt.Println("Vous avez actuellement", p.Mana, "/", p.ManaMax, "Mana")
				fmt.Println("Il vous manque moins de 20 mana, vous allez gaspiller la potion.")
				fmt.Println("1: Oui   2: Retour")
				if utils.ReadChoice("Ton choix : ", []string{"1", "2"}) == "2" {
					return
				}
			}

			p.Mana = utils.Clamp(p.Mana+20, 0, p.ManaMax)
			RemoveItem(p, "Potion de mana", 1)
			fmt.Println("Vous avez bu une potion de mana. Mana :", p.Mana, "/", p.ManaMax)
			return
		}
	}
	fmt.Println("Vous n'avez pas de potion de mana dans votre inventaire")
}
