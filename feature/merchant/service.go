package merchant

import (
	"Projet-Red/feature/character"
	"Projet-Red/utils"
)

// Market ouvre la boutique de Barnabé le Colporteur.
func Market(p *character.Player) {
	utils.Clear()
	DisplayMarketIntro()
	PromptMeetMerchant()
	for {
		utils.Clear()
		DisplayMarketMenu(p)
		choix := ReadMarketChoice()
		switch choix {
		case "1":
			utils.Clear()
			if !p.PotionGratuiteUtilisee {
				DisplayFreePotion()
				Buy(p, "Potion de santé", 0)
				p.PotionGratuiteUtilisee = true
			} else {
				DisplayBlankLine()
				Buy(p, "Potion de santé", 10)
				DisplayBlankLine()
			}
		case "2":
			utils.Clear()
			Buy(p, "Potion de poison", 10)
		case "3":
			utils.Clear()
			Buy(p, "Potion de mana", 12)
		case "4":
			utils.Clear()
			if len(p.Skill) > 0 {
				DisplayAlreadyHaveSkill()
				break
			}
			if character.CanAfford(p, 55) {
				p.Money -= 55
				p.Skill = append(p.Skill, SkillDeClasse(p.Class))
				DisplaySkillPurchased()
			} else {
				DisplayNotEnoughGold()
			}
		case "5":
			utils.Clear()
			if p.InventoryMax >= 40 {
				DisplayMaxBackpacks()
			} else if character.CanAfford(p, 30) {
				p.Money -= 30
				p.InventoryMax += 10
				DisplayBackpackUpgraded()
			} else {
				DisplayNotEnoughGold()
			}
		case "6":
			utils.Clear()
			Buy(p, "Fer", 8)
		case "7":
			utils.Clear()
			Buy(p, "Acier", 12)
		case "8":
			utils.Clear()
			Buy(p, "Cuir", 5)
		case "9":
			utils.Clear()
			Buy(p, "Bois", 6)
		case "10":
			utils.Clear()
			return
		}
		if choix != "10" {
			PromptContinue()
		}
	}
}

// Acheter tente d'acheter un objet : vérifie la place dans le sac et l'or disponible.
func Buy(p *character.Player, nom string, prix int) {
	if len(p.Inventory) >= p.InventoryMax {
		DisplayBagFull()
	} else if character.CanAfford(p, prix) {
		p.Money -= prix
		character.AddItem(p, nom, 1)
		DisplayPurchase(nom)
	} else {
		DisplayNotEnoughGold()
	}
}

// SkillDeClasse retourne le nom de l'attaque spéciale correspondant à la classe du joueur.
func SkillDeClasse(classe string) string {
	switch classe {
	case "Soldat":
		return "Sabre"
	case "Moine":
		return "Décharge d'énergie"
	case "Alchimiste":
		return "Potion de dégâts instantané"
	case "Samouraï":
		return "Katana"
	}
	return ""
}
