package feature

import (
	"Projet-Red/models"
	"fmt"
)

// Ajout de l'item dans l'inventaire

func AddItem(p *models.Player, nom string, quantite int) {
	for i := range p.Inventory {
		if p.Inventory[i].Name == nom {
			p.Inventory[i].Quantity += quantite
			return
		}
	}
	if len(p.Inventory) >= p.InventoryMax {
		fmt.Println("Votre Inventaire est plein !")
		return
	}
	p.Inventory = append(p.Inventory, models.Item{Name: nom, Quantity: quantite})
}

// Suppresion de l'item dans l'inventaire

func RemoveItem(p *models.Player, nom string, quantite int) {
	for i := range p.Inventory {
		if p.Inventory[i].Name == nom {
			p.Inventory[i].Quantity -= quantite
			if p.Inventory[i].Quantity <= 0 {
				p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
			}
			return
		}
	}
}
