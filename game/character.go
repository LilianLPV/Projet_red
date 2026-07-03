package game

import (
	"Projet-Red/models"
	"Projet-Red/utils"
	"fmt"
)


//Info Joueur 
func DisplayInfo(p *models.Player) {
	fmt.Printf("Nom : %v\n", p.Name)
	fmt.Printf("Classe : %v\n", p.Class)
	fmt.Printf("Level : %d\n", p.Level)
	fmt.Printf("XP : %d / %d\n", p.XP, p.XPMax)
	fmt.Printf("PV : %d / %d\n", p.HP, p.HPMax)
	fmt.Printf("Money : %v\n", p.Money)
	fmt.Printf("Mana : %d / %d\n", p.Mana, p.ManaMax)
	fmt.Printf("Speed : %d\n", p.Speed)
	fmt.Printf("StunChance : %d %%\n", p.StunChance)
	fmt.Printf("Dodge : %d %%\n", p.Dodge)
	fmt.Printf("Attaque de base : %s (%d dégâts)\n", p.BaseAttack, p.Damage)
	for _, s := range p.Skill {
		fmt.Printf("Skill : %v\n", s)
	}
	fmt.Printf("Inventory : %d / %d\n", len(p.Inventory), p.InventoryMax)
	if p.Equipment.Head != nil {
		fmt.Printf("Tête : %s\n", p.Equipment.Head.Name)
	} else {
		fmt.Println("Tête : (vide)")
	}
	if p.Equipment.Torso != nil {
		fmt.Printf("Torse : %s\n", p.Equipment.Torso.Name)
	} else {
		fmt.Println("Torse : (vide)")
	}
	if p.Equipment.Feet != nil {
		fmt.Printf("Pied : %s\n", p.Equipment.Feet.Name)
	} else {
		fmt.Println("Pied : (vide)")
	}

}
// Création du Personnages
func CharCreation() models.Player {
	nom := utils.ReadName("Entre ton nom : ")
	fmt.Println("1) Soldat")
	fmt.Println("2) Moine")
	fmt.Println("3) Alchimiste")
	fmt.Println("4) Samouraï")

	choix := utils.ReadChoice("Ton choix : ", []string{"1", "2", "3", "4"})
	var classe string
	switch choix {
	case "1":
		classe = "Soldat"
	case "2":
		classe = "Moine"
	case "3":
		classe = "Alchimiste"
	case "4":
		classe = "Samouraï"
	}
	return models.Init(nom, classe)
}
