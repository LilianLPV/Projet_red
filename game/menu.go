package game

import (
	"Projet-Red/feature"
	"Projet-Red/models"
	"Projet-Red/utils"
	"fmt"
	"time"
)
//Menu principal
func MainMenu(p *models.Player) {

	for {
		time.Sleep(800 * time.Millisecond)
		utils.Clear()
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
		choix := utils.ReadChoice("Ton choix : ", []string{"1", "2", "3", "4", "5", "6", "7"})
		switch choix {
		case "1":
			utils.Clear()
			DisplayInfo(p)
			fmt.Println("\nAppuie sur Entrée pour continuer...")
			utils.ReadInput()
			utils.Clear()
		case "2":
			utils.Clear()
			AccessInventory(p)
		case "3":
			utils.Clear()
			feature.Market(p)
		case "4":
			utils.Clear()
			feature.Forgeron(p)
		case "5":
			fmt.Println("=== Mode entraînement ===")
			mannequin := models.MonsterDefs["Mannequin"]
			Combat(p, mannequin)
		case "6":
			utils.Clear()
			StartHistory(p)
		case "7":
			utils.Clear()
			return
		}
	}
}
