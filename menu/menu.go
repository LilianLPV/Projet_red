package menu

import (
	"fmt"
	"time"

	"Projet-Red/feature/blacksmith"
	"Projet-Red/feature/character"
	"Projet-Red/feature/combat"
	"Projet-Red/feature/merchant"
	"Projet-Red/feature/story"
	"Projet-Red/utils"
)

// MainMenu affiche le menu principal du jeu et redirige vers chaque fonctionnalité.
func MainMenu(p *character.Player) {
	for {
		time.Sleep(800 * time.Millisecond)
		utils.Clear()
		displayHeader()
		choix := readMainChoice()
		switch choix {
		case "1":
			utils.Clear()
			character.DisplayInfo(p)
			fmt.Println("\nAppuie sur Entrée pour continuer...")
			utils.ReadInput()
			utils.Clear()
		case "2":
			utils.Clear()
			character.AccessInventory(p)
		case "3":
			utils.Clear()
			merchant.Market(p)
		case "4":
			utils.Clear()
			blacksmith.Blacksmith(p)
		case "5":
			fmt.Println("=== Mode entraînement ===")
			mannequin := combat.MonsterDefs["Mannequin"]
			combat.Combat(p, mannequin)
		case "6":
			utils.Clear()
			story.StartHistory(p)
		case "7":
			utils.Clear()
			return
		}
	}
}

func displayHeader() {
	fmt.Println("==============================")
	fmt.Println("======= \033[96mJeu Projet Red\033[0m =======")
	fmt.Println("==============================")
	fmt.Println(" ")
	fmt.Println("\033[96m1\033[0m: 	Information personnage")
	fmt.Println("\033[96m2\033[0m: 	Inventaire")
	fmt.Println("\033[96m3\033[0m: 	Marchand")
	fmt.Println("\033[96m4\033[0m: 	Forgeron")
	fmt.Println("\033[96m5\033[0m: 	Partir a l'entrainement")
	fmt.Println("\033[96m6\033[0m: 	Continuer le périple !")
	fmt.Println("\033[96m7\033[0m: 	Quitter")
}

func readMainChoice() string {
	return utils.ReadChoice("Ton choix : ", []string{"1", "2", "3", "4", "5", "6", "7"})
}
