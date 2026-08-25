package blacksmith

import (
	"fmt"

	"Projet-Red/feature/character"
	"Projet-Red/utils"
)

// DisplayForgeIntro affiche l'illustration ASCII de la forge.
func DisplayForgeIntro() {
	fmt.Println(utils.Yellow + `
+----------------------------------------------+
|   ||                                         |
| +----------+            _______              |
| | Forgeron |          /         \            |
| +----------+         |   O   O   |           |
|                      |           |           |
|                       \  \___/  /            |
|      _____             \_______/____         |
|    /     /|           /             \        |
|   /_____/ /          /               \       |
|         \/          /                 |      |
|         /          |                  |      |
|        /           |                  |      |
+----------------------------------------------+
|                                              |
|          +------------------------+          |
|          |                        |          |
|          |         Thorin         |          |
|          |                        |          |
|          +------------------------+          |
|                                              |
+----------------------------------------------+
` + utils.Reset)
}

// PromptMeetBlacksmith invite le joueur à entrer chez le forgeron.
func PromptMeetBlacksmith() {
	fmt.Println("\nAppuie sur Entrée pour rencontrer Thorin...")
	utils.ReadInput()
}

// DisplayForgeMenu affiche les recettes disponibles chez le forgeron.
func DisplayBlacksmithMenu(p *character.Player) {
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
}

// ReadForgeChoice lit le choix du joueur chez le forgeron.
func ReadForgeChoice() string {
	return utils.ReadChoice("Ton choix : ", []string{"1", "2", "3", "4"})
}

// DisplaySlotTaken prévient que l'emplacement d'équipement est déjà occupé.
func DisplaySlotTaken() {
	fmt.Println("Tu as déjà un équipement à cet emplacement !")
}

// DisplayNotEnoughGold prévient que le joueur n'a pas assez d'or pour forger.
func DisplayNotEnoughGold(name string) {
	fmt.Println("Pas assez d'or pour forger le", name, "!")
}

// DisplayMissingMaterials prévient que le joueur n'a pas les matériaux nécessaires.
func DisplayMissingMaterials(name string) {
	fmt.Println("Tu n'as malheureusement pas les matériaux pour faire ce magnifique", name, "!")
}

// DisplayCraftSuccess confirme la réussite de la forge.
func DisplayCraftSuccess(name string) {
	fmt.Println("Voilà votre", name, "est fini, il vous sera très utile pour la suite !")
}
