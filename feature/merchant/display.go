package merchant

import (
	"fmt"

	"Projet-Red/feature/character"
	"Projet-Red/utils"
)

// DisplayMarketIntro affiche l'illustration ASCII du marché.
func DisplayMarketIntro() {
	fmt.Println(utils.Yellow + `
+----------------------------------------------------+
|    |  |                                            |
|  +------+              _...._                      |
|  |Marché|            /        \                    |
|  +------+           |  o    o  |                   |
|                     |    __    |                   |
|                      \  \__/  /                    |
|     _                 '-....-'                     |
|    |_|             _.-'      '-._                  |
|  +-----+          /              \                 |
|  |     |         /                \                |
|  |     |        |                  |               |
+----------------------------------------------------+
 \                                                  /
  \           +------------------------+           /
   \          |                        |          /
    \         |        Barnabé         |         /
     \        |                        |        /
      \       +------------------------+       /
       +--------------------------------------+
` + utils.Reset)
}

func PromptMeetMerchant() {
	fmt.Println("\nAppuie sur Entrée pour rencontrer Barnabé...")
	utils.ReadInput()
}

func DisplayMarketMenu(p *character.Player) {
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
}

func ReadMarketChoice() string {
	return utils.ReadChoice("Ton choix : ", []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"})
}

func PromptContinue() {
	fmt.Println("\nAppuie sur Entrée pour continuer...")
	utils.ReadInput()
}

func DisplayBlankLine() {
	fmt.Println(" ")
}

func DisplayFreePotion() {
	fmt.Println("C'est ta première potion, elle est offerte par la maison !")
}

func DisplayAlreadyHaveSkill() {
	fmt.Println("Tu as déjà ta capacité !")
}

func DisplaySkillPurchased() {
	fmt.Println("Vous avez acheté votre attaque spéciale !")
}

func DisplayNotEnoughGold() {
	fmt.Println("Pas assez d'or !")
}

func DisplayMaxBackpacks() {
	fmt.Println("Tu as déjà 3 sacs, c'est le max !")
}

func DisplayBackpackUpgraded() {
	fmt.Println("Vous avez acheté un sac à dos amélioré + 10 places !")
}

func DisplayBagFull() {
	fmt.Println("Ton sac est plein, fais de la place !")
}

func DisplayPurchase(nom string) {
	fmt.Println("Vous avez acheté :", nom)
}
