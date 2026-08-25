package character

import (
	"fmt"

	"Projet-Red/feature/inventory"
	"Projet-Red/utils"
)

// displayClassChoices affiche la liste des classes disponibles à la création du personnage.
func displayClassChoices() {
	fmt.Println("1) Soldat")
	fmt.Println("2) Moine")
	fmt.Println("3) Alchimiste")
	fmt.Println("4) Samouraï")
}

// DisplayInfo affiche la fiche complète du personnage.
func DisplayInfo(p *Player) {
	fmt.Printf("Nom : %v\n", p.Name)
	fmt.Printf("Classe : %v\n", p.Class)
	fmt.Printf("Niveau : %d\n", p.Level)
	fmt.Printf("XP : %d / %d\n", p.XP, p.XPMax)
	fmt.Printf("PV : %d / %d\n", p.HP, p.HPMax)
	fmt.Printf("Argent : %v\n", p.Money)
	fmt.Printf("Mana : %d / %d\n", p.Mana, p.ManaMax)
	fmt.Printf("Vitesse : %d\n", p.Speed)
	fmt.Printf("Chance d'étourdissement : %d %%\n", p.StunChance)
	fmt.Printf("Esquive : %d %%\n", p.Dodge)
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

// DisplayXPGain annonce le gain d'XP.
func DisplayXPGain(amount int) {
	fmt.Println("Vous gagnez", amount, "XP !")
}

// DisplayLevelUp annonce la montée de niveau.
func DisplayLevelUp(level int, maxReached bool) {
	if maxReached {
		fmt.Println("NIVEAU SUPÉRIEUR ! Vous êtes niveau", level, "(niveau maximum atteint)")
	} else {
		fmt.Println("NIVEAU SUPÉRIEUR ! Vous êtes niveau", level)
	}
}

// DisplayNoItem prévient que le joueur ne possède pas l'objet demandé.
func DisplayNoItem(name string) {
	fmt.Println("Vous n'avez pas de", name, "dans votre inventaire")
}

// DisplayHealthPotionUsed confirme la consommation d'une potion de santé.
func DisplayHealthPotionUsed(hp, hpMax int) {
	fmt.Println("Vous avez bu une potion de santé. PV :", hp, "/", hpMax)
}

// DisplayManaPotionUsed confirme la consommation d'une potion de mana.
func DisplayManaPotionUsed(mana, manaMax int) {
	fmt.Println("Vous avez bu une potion de mana. Mana :", mana, "/", manaMax)
}

// PromptWastefulHeal prévient le joueur qu'il va gaspiller une potion de santé et demande confirmation.
func PromptWastefulHeal(current, max int) bool {
	fmt.Println("Vous avez actuellement", current, "/", max, "PV")
	fmt.Println("Il vous manque moins de 50 PV, vous allez gaspiller la potion.")
	fmt.Println("1: Oui   2: Retour")
	return utils.ReadChoice("Ton choix : ", []string{"1", "2"}) == "1"
}

// PromptWastefulMana prévient le joueur qu'il va gaspiller une potion de mana et demande confirmation.
func PromptWastefulMana(current, max int) bool {
	fmt.Println("Vous avez actuellement", current, "/", max, "Mana")
	fmt.Println("Il vous manque moins de 20 mana, vous allez gaspiller la potion.")
	fmt.Println("1: Oui   2: Retour")
	return utils.ReadChoice("Ton choix : ", []string{"1", "2"}) == "1"
}

// DisplayInventoryScreen affiche le contenu du sac et les actions possibles.
func DisplayInventoryScreen(p *Player) {
	fmt.Println("Voici votre Inventaire :")
	inventory.DisplayList(p.Inventory)
	fmt.Println(" ")
	fmt.Println("1) Boire une potion de santé")
	fmt.Println("2) Boire une potion de mana")
	fmt.Println("3) Retour")
}
