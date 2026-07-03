package feature

import (
	"Projet-Red/models"
	"Projet-Red/utils"
	"fmt"
)


// Pour équiper un un équipement
func equipmentSlot(p *models.Player, slot string) **models.Item {
	switch slot {
	case "Head":
		return &p.Equipment.Head
	case "Torso":
		return &p.Equipment.Torso
	case "Feet":
		return &p.Equipment.Feet
	}
	return nil
}
// Pour crafter un équipement
func Craft(p *models.Player, r models.Recipe) bool {
	slot := equipmentSlot(p, r.Slot)
	if slot == nil {
		return false
	}
	if *slot != nil {
		fmt.Println("Tu as déjà un équipement à cet emplacement !")
		return false
	}
	if !HowManyMoney(p, r.Cost) {
		fmt.Println("Pas assez d'or pour forger le", r.Name, "!")
		return false
	}
	for materiau, quantite := range r.Materials {
		if QuantiteItem(p, materiau) < quantite {
			fmt.Println("Tu n'as malheureusement pas les matériaux pour faire ce magnifique", r.Name, "!")
			return false
		}
	}

	p.Money -= r.Cost
	for materiau, quantite := range r.Materials {
		RemoveItem(p, materiau, quantite)
	}
	*slot = &models.Item{Name: r.Name}
	r.Apply(p)
	fmt.Println("Voilà votre", r.Name, "est fini, il vous sera très utile pour la suite !")
	return true
}
// Le forgeron qui affiche les craft et les resultats
func Forgeron(p *models.Player) {
	utils.Clear()
	AfficherForgeron()
	fmt.Println("\nAppuie sur Entrée pour rencontrer Thorin...")
	utils.ReadInput()
	utils.Clear()
	for {
		fmt.Println("==============================")
		fmt.Println("=== \033[96mThorin le Marteau\033[0m ===")
		fmt.Println("==============================")
		fmt.Println(" ")
		fmt.Println("Vous avez actuellement ", p.Money, " PO")
		fmt.Println(" ")
		fmt.Println("Voici toutes les armures que je peux vous forger en échange de quelques pièces d'or ahahah !! ")
		fmt.Println(" ")
		fmt.Println(utils.Bold + "1) Menpo" + utils.Reset + " — masque de samouraï")
		fmt.Println("   Effet : +3 HP max, +2% étourdissement")
		fmt.Println("   Recette : 1 Fer + 1 Cuir + 15 PO")
		fmt.Println()
		fmt.Println(utils.Bold + "2) Tatami-do" + utils.Reset + " — plastron de samouraï")
		fmt.Println("   Effet : +12 HP max")
		fmt.Println("   Recette : 2 Acier + 1 Bois + 1 Cuir + 30 PO")
		fmt.Println()
		fmt.Println(utils.Bold + "3) Waraji" + utils.Reset + " — sandales en bois de cerisier")
		fmt.Println("   Effet : +3% esquive")
		fmt.Println("   Recette : 2 Cuir + 1 Bois + 20 PO")
		fmt.Println()
		fmt.Println("4) Je ne souhaite rien forger au final...")
		fmt.Println(" ")
		choix := utils.ReadChoice("Ton choix : ", []string{"1", "2", "3", "4"})
		switch choix {
		case "1":
			utils.Clear()
			Craft(p, models.RecipeDefs[0])
		case "2":
			utils.Clear()
			Craft(p, models.RecipeDefs[1])
		case "3":
			utils.Clear()
			Craft(p, models.RecipeDefs[2])
		case "4":
			utils.Clear()
			return
		}
	}
}
// Savoir si il a assez de d'item pour craft
func QuantiteItem(p *models.Player, nom string) int {
	for _, item := range p.Inventory {
		if item.Name == nom {
			return item.Quantity
		}
	}
	return 0
}
