package game

import (
	"Projet-Red/feature"
	"Projet-Red/models"
	"Projet-Red/utils"
	"fmt"
)


//Pour avoir accès a l'inventaire
func AccessInventory(p *models.Player) {
	utils.Clear()
	fmt.Println("Voici votre Inventaire :")
	if len(p.Inventory) == 0 {
		fmt.Println("Votre Inventaire est vide !")
	} else {
		for _, i := range p.Inventory {
			fmt.Println(i.Name, i.Quantity)
		}
	}
	fmt.Println(" ")
	fmt.Println("1) Boire une potion de santé")
	fmt.Println("2) Boire une potion de mana")
	fmt.Println("3) Retour")
	choix := utils.ReadChoice("Ton choix : ", []string{"1", "2", "3"})
	switch choix {
	case "1":
		feature.TakePot(p)
	case "2":
		feature.ManaPot(p)
	case "3":
		return
	}
}

