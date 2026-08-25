package blacksmith

import (
	"Projet-Red/feature/character"
	"Projet-Red/feature/inventory"
	"Projet-Red/utils"
)

// equipmentSlot retourne un pointeur vers l'emplacement d'équipement correspondant au slot donné.
func equipmentSlot(p *character.Player, slot string) **inventory.Item {
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

// Craft forge un équipement à partir d'une recette, si le joueur a l'or, les matériaux,
// et un emplacement libre. Retourne true si la forge a réussi.
func Craft(p *character.Player, r Recipe) bool {
	slot := equipmentSlot(p, r.Slot)
	if slot == nil {
		return false
	}
	if *slot != nil {
		DisplaySlotTaken()
		return false
	}
	if !character.CanAfford(p, r.Cost) {
		DisplayNotEnoughGold(r.Name)
		return false
	}
	for material, quantite := range r.Materials {
		if character.QuantityOf(p, material) < quantite {
			DisplayMissingMaterials(r.Name)
			return false
		}
	}

	p.Money -= r.Cost
	for material, quantite := range r.Materials {
		character.RemoveItem(p, material, quantite)
	}
	*slot = &inventory.Item{Name: r.Name}
	r.Apply(p)
	DisplayCraftSuccess(r.Name)
	return true
}

func Blacksmith(p *character.Player) {
	utils.Clear()
	DisplayForgeIntro()
	PromptMeetBlacksmith()
	utils.Clear()
	for {
		DisplayBlacksmithMenu(p)
		choice := ReadForgeChoice()
		switch choice {
		case "1":
			utils.Clear()
			Craft(p, RecipeDefs[0])
		case "2":
			utils.Clear()
			Craft(p, RecipeDefs[1])
		case "3":
			utils.Clear()
			Craft(p, RecipeDefs[2])
		case "4":
			utils.Clear()
			return
		}
	}
}
