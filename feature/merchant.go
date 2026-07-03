package feature

import (
	"Projet-Red/models"
	"Projet-Red/utils"
	"fmt"
)

//Affichage de marché avec les achats..

func Market(p *models.Player) {
	utils.Clear()
	AfficherMarchand()
	fmt.Println("\nAppuie sur Entrée pour rencontrer Barnabé...")
	utils.ReadInput()
	for {
		utils.Clear()
		fmt.Println("==============================")
		fmt.Println("=== \033[96mBarnabé le Colporteur\033[0m ===")
		fmt.Println("==============================")
		fmt.Println(" ")
		fmt.Println("Vous avez actuellement ", p.Money, " PO")
		fmt.Println(" ")
		fmt.Println("Voici les consommables que je vend !")
		fmt.Println(" ")
		fmt.Println("1) Potion de", utils.Red, "santé ", utils.Reset, "10 PO")
		fmt.Println("2) Potion de", utils.Green, "poison ", utils.Reset, "10 PO")
		fmt.Println("3) Potion de", utils.Blue, "mana ", utils.Reset, "12 PO")
		fmt.Println("4) Attaque spéciale", utils.Red, SkillDeClasse(p.Class), utils.Reset, "55 PO")
		fmt.Println(" ")
		fmt.Println("Voici l'amélioration de votre inventaire !")
		fmt.Println(" ")
		fmt.Println("5) Sac à dos amélioré", utils.Blue, "+ 10 places ", utils.Reset, "30 PO")
		fmt.Println(" ")
		fmt.Println("Voici les objets pour faire votre armure chez le forgeron !")
		fmt.Println(" ")
		fmt.Println("6) Fer 8 PO")
		fmt.Println("7) Acier 12 PO")
		fmt.Println("8) Cuir 5 PO")
		fmt.Println("9) Bois 6 PO")
		fmt.Println(" ")
		fmt.Println("10) Je ne souhaite rien acheter au final...")
		fmt.Println(" ")
		choix := utils.ReadChoice("Ton choix : ", []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"})
		switch choix {
		case "1":
			utils.Clear()
			if !p.PotionGratuiteUtilisee {
				fmt.Println("C'est ta première potion, elle est offerte par la maison !")
				Acheter(p, "Potion de santé", 0)
				p.PotionGratuiteUtilisee = true
			} else {
				fmt.Println(" ")
				Acheter(p, "Potion de santé", 10)
				fmt.Println(" ")
			}
		case "2":
			utils.Clear()
			Acheter(p, "Potion de poison", 10)
		case "3":
			utils.Clear()
			Acheter(p, "Potion de mana", 12)

		case "4":
			utils.Clear()
			if len(p.Skill) > 0 {
				fmt.Println("Tu as déjà ta capacité !")
				break
			}
			if HowManyMoney(p, 55) {
				p.Money -= 55
				p.Skill = append(p.Skill, SkillDeClasse(p.Class))
				fmt.Println("Vous avez acheté votre attaque spéciale !")
			} else {
				fmt.Println("Pas assez d'or !")
			}
		case "5":
			utils.Clear()
			if p.InventoryMax >= 40 {
				fmt.Println("Tu as déjà 3 sacs, c'est le max !")
			} else if HowManyMoney(p, 30) {
				p.Money -= 30
				p.InventoryMax += 10
				fmt.Println("Vous avez acheté un sac à dos amélioré + 10 places !")
			} else {
				fmt.Println("Pas assez d'or !")
			}
		case "6":
			utils.Clear()
			Acheter(p, "Fer", 8)
		case "7":
			utils.Clear()
			Acheter(p, "Acier", 12)
		case "8":
			utils.Clear()
			Acheter(p, "Cuir", 5)
		case "9":
			utils.Clear()
			Acheter(p, "Bois", 6)
		case "10":
			utils.Clear()
			return
		}
		if choix != "10" {
			fmt.Println("\nAppuie sur Entrée pour continuer...")
			utils.ReadInput()
		}
	}
}

// Si j'ai assez d'argent par rapport au prix 

func HowManyMoney(p *models.Player, cost int) bool {
	return p.Money >= cost
}

// Achat des items + vérification si j'ai assez de place + assez d'argent

func Acheter(p *models.Player, nom string, prix int) {
	if len(p.Inventory) >= p.InventoryMax {
		fmt.Println("Ton sac est plein, fais de la place !")
	} else if HowManyMoney(p, prix) {
		p.Money -= prix
		AddItem(p, nom, 1)
		fmt.Println("Vous avez acheté :", nom)
	} else {
		fmt.Println("Pas assez d'or !")
	}
}

// Apparation en boutique de l'attaque spé en fonction de ta classe 

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
