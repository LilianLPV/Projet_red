package game

import (
	"fmt"
	"time"

	"Projet-Red/models"
	"Projet-Red/utils"
)
//Introduction du jeu 
func IntroLore(p *models.Player) {
	time.Sleep(600 * time.Millisecond)
	utils.SlowPrint("Voici ton histoire, "+p.Name, 50*time.Millisecond)
	utils.SlowPrint("Tout a commencé il y a longtemps. ", 50*time.Millisecond)
	utils.SlowPrint("Il y a eu un homme.", 50*time.Millisecond)
	utils.SlowPrint("Celui que tu considérais comme ton père.", 50*time.Millisecond)

	time.Sleep(800 * time.Millisecond)

	utils.SlowPrint("Il t’a tout appris.", 50*time.Millisecond)
	utils.SlowPrint("Le combat. La discipline. La survie.", 50*time.Millisecond)

	time.Sleep(700 * time.Millisecond)

	utils.SlowPrint("Puis, un jour… "+utils.Purple+"il a disparu."+utils.Reset, 55*time.Millisecond)

	time.Sleep(500 * time.Millisecond)

	utils.SlowPrint("Sans bruit.", 50*time.Millisecond)
	utils.SlowPrint("Sans trace.", 50*time.Millisecond)

	time.Sleep(700 * time.Millisecond)

	utils.SlowPrint("Seulement un message laissé derrière lui :", 55*time.Millisecond)

	utils.SlowPrint(utils.Cyan+"\"Je dois partir. Désolé. Ne cherche pas à me retrouver.\""+utils.Reset, 45*time.Millisecond)

	time.Sleep(900 * time.Millisecond)

	utils.SlowPrint("Tu avais 12 ans.", 50*time.Millisecond)
	utils.SlowPrint("Et tu as obéi.", 50*time.Millisecond)

	time.Sleep(800 * time.Millisecond)

	utils.SlowPrint("Mais quelque chose dans ce message ne t’a jamais semblé juste.", 50*time.Millisecond)

	time.Sleep(900 * time.Millisecond)

	utils.SlowPrint("En soufflant tes seize bougies tu réalises quelque chose.", 50*time.Millisecond)

	utils.SlowPrint("Ce message sonnait faux.", 55*time.Millisecond)

	time.Sleep(700 * time.Millisecond)

	utils.SlowPrint(utils.Yellow+"Tu dois le retrouver."+utils.Reset, 70*time.Millisecond)

	time.Sleep(1200 * time.Millisecond)

	utils.SlowPrint(utils.Cyan+"..."+utils.Reset, 80*time.Millisecond)

	time.Sleep(600 * time.Millisecond)
}
func StartMenu(p *models.Player) {
	utils.SlowPrint("=== QUE FAIS-TU ? ===", 40*time.Millisecond)
	fmt.Println()

	utils.SlowPrint(utils.Cyan+"1)"+utils.Reset+" Partir à sa recherche", 35*time.Millisecond)
	utils.SlowPrint(utils.Cyan+"2)"+utils.Reset+" Rester ici, ignorer le message", 35*time.Millisecond)

	fmt.Println()

	choix := utils.ReadChoice("Ton choix : ", []string{"1", "2"})
	switch choix {

	case "1":
		utils.SlowPrint("Tu resserres ton équipement.", 40*time.Millisecond)
		utils.SlowPrint("Le voyage commence.", 60*time.Millisecond)
		MainMenu(p)

	case "2":
		utils.SlowPrint("Tu regardes le message une dernière fois...", 40*time.Millisecond)
		time.Sleep(600 * time.Millisecond)

		utils.SlowPrint("Tu le ranges.", 50*time.Millisecond)
		utils.SlowPrint("Peut-être… qu’il vaut mieux rester ici.", 50*time.Millisecond)

		time.Sleep(800 * time.Millisecond)

		utils.SlowPrint(utils.Purple+"Parfois, l’inaction est aussi un choix."+utils.Reset, 50*time.Millisecond)

		time.Sleep(600 * time.Millisecond)

		utils.SlowPrint("Fin.", 80*time.Millisecond)
		return
	}
}


// L'histoire du jeu 
func StartHistory(p *models.Player) {
	switch p.Progression {
	case 0:
		utils.SlowPrint("Vous vous rendez à Henan", 50*time.Millisecond)
		utils.SlowPrint("Vous voyez un monastère Shaolin", 50*time.Millisecond)
		time.Sleep(600 * time.Millisecond)
		utils.SlowPrint("Que faites-vous ? Entrez-vous dans le monastère ?", 50*time.Millisecond)

		utils.SlowPrint(utils.Cyan+"1)"+utils.Reset+" Oui", 35*time.Millisecond)
		utils.SlowPrint(utils.Cyan+"2)"+utils.Reset+" Non", 35*time.Millisecond)

		choix := utils.ReadChoice("Ton choix : ", []string{"1", "2"})
		switch choix {
		case "1":
			utils.SlowPrint("Vous entrez dans le monastère. Un moine s'y trouve.", 40*time.Millisecond)
			utils.SlowPrint("Il se dirige vers vous et vous demande comment vous vous appelez.", 60*time.Millisecond)
			utils.SlowPrint("Enchanté, "+p.Name, 60*time.Millisecond)
			utils.SlowPrint("Vous lui indiquez que vous êtes à la recherche d'un moine du nom de Sora.", 60*time.Millisecond)
			utils.SlowPrint("Il vous dit avoir croisé ce moine il y a trois ans.", 60*time.Millisecond)
			utils.SlowPrint("Sora lui avait parlé d'une quête qu'il devait accomplir.", 60*time.Millisecond)
			time.Sleep(600 * time.Millisecond)
			utils.SlowPrint("Vous lui demandez de quoi il s'agissait.", 60*time.Millisecond)
			utils.SlowPrint("Il vous baratine et vous raconte une histoire qui ne tient pas la route.", 60*time.Millisecond)
			utils.SlowPrint("Vous décidez de le menacer pour connaître la vraie histoire.", 60*time.Millisecond)
			utils.SlowPrint("Il refuse de céder à cette menace et vous provoque en duel.", 60*time.Millisecond)
			if Combat(p, models.MonsterDefs["Moine renégat"]) {
				utils.SlowPrint("Le moine, vaincu, lâche un indice avant de s'effondrer :", 50*time.Millisecond)
				utils.SlowPrint("Diriger vous aux égoûts du monastère", 50*time.Millisecond)
				p.Progression = 1
				utils.SlowPrint("Vous décidez de faire une pause avant de continuer votre quête...", 50*time.Millisecond)
				return
			}
		case "2":
			utils.SlowPrint("Vous décidez de faire demi-tour, mais un garde vous interpelle au loin.", 40*time.Millisecond)
			utils.SlowPrint("Il vous demande ce que vous faites ici.", 60*time.Millisecond)
			utils.SlowPrint("Vous lui répondez que vous cherchez un moine nommé Sora.", 60*time.Millisecond)
			utils.SlowPrint("Il vous dit que vous semblez louche.", 60*time.Millisecond)
			utils.SlowPrint("Pour être sûr de vos propos, il décide de vous attaquer en duel.", 60*time.Millisecond)
			if Combat(p, models.MonsterDefs["Garde"]) {
				utils.SlowPrint("Le garde, vaincu, lâche un indice avant de s'effondrer :", 50*time.Millisecond)
				utils.SlowPrint("Diriger vous aux égoûts du monastère", 50*time.Millisecond)
				p.Progression = 1
				utils.SlowPrint("Vous décidez de faire une pause avant de continuer votre quête...", 50*time.Millisecond)
				return
			}
		}
	case 1:
		utils.SlowPrint("Suite à ce dur combat, vous arrivez à trouver les égoûts du monastère.", 60*time.Millisecond)
		utils.SlowPrint("Une fois entré, vous entendez des bruits étranges.", 60*time.Millisecond)
		utils.SlowPrint("Des rats vous foncent dessus vous devez vous défendre !", 60*time.Millisecond)
		if Combat(p, models.MonsterDefs["Rat-agile"]) {
			utils.SlowPrint("Vous avez repoussé les rats et pouvez poursuivre votre exploration.", 50*time.Millisecond)
			p.Progression = 2
			return
		}
	case 2:
		utils.SlowPrint("Une fois tuée ces rats vous vous enfoncez dans les égoûts.", 60*time.Millisecond)
		utils.SlowPrint("Vous voyez tout au bout du couloir un chaman.", 60*time.Millisecond)
		utils.SlowPrint("Vous lui demandez ce qu'il fait ici et est-ce qu'il connaît Sora.", 60*time.Millisecond)
		utils.SlowPrint("Il vous dit que oui il connaît cet homme, il vous demande de vous asseoir.", 60*time.Millisecond)
		utils.SlowPrint("Une fois assis il vous demande de lui donner vos mains, vous le faites.", 60*time.Millisecond)
		utils.SlowPrint("Il vous raconte votre histoire dans les moindres détails..., cela vous choque.", 60*time.Millisecond)
		utils.SlowPrint("Vous lui demandez si Sora est dans les parages, et si oui ou il se trouve ?", 60*time.Millisecond)
		utils.SlowPrint("Il vous répond que oui il suffit d'ouvrir la trappe qui se trouve à côté de lui.", 60*time.Millisecond)
		utils.SlowPrint("Vous êtes super heureux de cette nouvelle vous le remerciez !!", 60*time.Millisecond)
		utils.SlowPrint("Cependant ce shaman vous demande de l'argent contre les informations qu'il vous a données", 60*time.Millisecond)
		utils.SlowPrint("Deux choix s'offrent à vous : lui donnée 30 PO ou l'affronter", 60*time.Millisecond)
		utils.SlowPrint(utils.Cyan+"1)"+utils.Reset+" Combattre le shaman", 35*time.Millisecond)
		utils.SlowPrint(utils.Cyan+"2)"+utils.Reset+" Donner les 30 PO", 35*time.Millisecond)
		choix := utils.ReadChoice("Ton choix : ", []string{"1", "2"})
		switch choix {
		case "1":
			utils.SlowPrint("Le combat, alors. Qu'il en soit ainsi.", 60*time.Millisecond)
			if Combat(p, models.MonsterDefs["Shaman"]) {
				utils.SlowPrint("Le shaman, vaincu, l'accès vers la trappe est libre !", 50*time.Millisecond)
				p.Progression = 3
				utils.SlowPrint("Vous décidez de faire une pause avant de continuer votre quête...", 50*time.Millisecond)
				return
			}
		case "2":
			if p.Money >= 30 {
				p.Money -= 30
				utils.SlowPrint("Vous lui donnez les 30 pièces d'or !", 60*time.Millisecond)
				p.Progression = 3
				utils.SlowPrint("Vous décidez de faire une pause avant de continuer votre quête...", 50*time.Millisecond)
				return
			} else {
				utils.SlowPrint("Tu as essayé de la mettre à l'envers tu n'as pas assez de pièces d'or !", 60*time.Millisecond)
				if Combat(p, models.MonsterDefs["Shaman"]) {
					utils.SlowPrint("Le shaman, vaincu, l'accès vers la trappe est libre !", 50*time.Millisecond)
					p.Progression = 3
					utils.SlowPrint("Vous décidez de faire une pause avant de continuer votre quête...", 50*time.Millisecond)
					return
				}
			}

		}
	case 3:
		utils.SlowPrint("Vous entrez dans cette trappe, vous découvrez une salle avec une lumière aveuglante.", 60*time.Millisecond)
		utils.SlowPrint("Vos yeux se sont habitués à cette lumière et vous apercevez une personne...", 60*time.Millisecond)
		utils.SlowPrint("Cette personne, c'est lui. Sora. Vous le reconnaissez immédiatement. Il est assis sur une chaise, il vous attend.", 60*time.Millisecond)
		utils.SlowPrint("Sora se lève. Il vous dit : \"Félicitations pour être arrivé jusqu'ici, "+p.Name+".\"", 60*time.Millisecond)
		utils.SlowPrint("\"Je vais t'enseigner une dernière leçon.\"", 60*time.Millisecond)
		if Combat(p, models.MonsterDefs["Sora"]) {
			utils.SlowPrint("Vous avez battu Sora !", 50*time.Millisecond)
			utils.SlowPrint("Il est étonné de voir que vous avez autant ...progressé", 50*time.Millisecond)
			utils.SlowPrint("Il est tellement fier de vous.", 50*time.Millisecond)
			utils.SlowPrint("Vous décidez de retourner à ... votre village", 50*time.Millisecond)
			p.Progression = 4
			utils.SlowPrint("Votre quête est terminée", 50*time.Millisecond)
		}

	}
}
