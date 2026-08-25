package story

import (
	"os"

	"Projet-Red/feature/character"
	"Projet-Red/feature/combat"
)

// IntroLore joue le texte d'introduction de l'histoire du jeu.
func IntroLore(p *character.Player) {
	DisplayIntro(p.Name)
}

// StartMenu propose au joueur de partir en quête ou de rester.
func StartMenu(p *character.Player) bool {
	DisplayStartPrompt()
	choix := ReadStartChoice()
	switch choix {
	case "1":
		DisplayStartQuest()
		return true
	case "2":
		DisplayPassiveRemaining()
		os.Exit(0)
	}
	return false
}

// StartHistory fait progresser la quête principale selon l'avancement du joueur.
func StartHistory(p *character.Player) {
	switch p.Progression {
	case 0:
		sceneMonastery(p)
	case 1:
		sceneSewer(p)
	case 2:
		sceneChaman(p)
	case 3:
		sceneSora(p)
	}
}

func sceneMonastery(p *character.Player) {
	DisplayScene0Intro()
	choix := ReadMonasteryChoice()
	switch choix {
	case "1":
		DisplayScene0EnterMonastery(p.Name)
		if combat.Combat(p, combat.MonsterDefs["Moine renégat"]) {
			DisplayScene0MonkDefeated()
			p.Progression = 1
			DisplayPause()
		}
	case "2":
		DisplayScene0TurnAway()
		if combat.Combat(p, combat.MonsterDefs["Garde"]) {
			DisplayScene0GuardDefeated()
			p.Progression = 1
			DisplayPause()
		}
	}
}

func sceneSewer(p *character.Player) {
	DisplayScene1Intro()
	if combat.Combat(p, combat.MonsterDefs["Rat-agile"]) {
		DisplayScene1Cleared()
		p.Progression = 2
	}
}

func sceneChaman(p *character.Player) {
	DisplayScene2Intro()
	choix := ReadShamanChoice()
	switch choix {
	case "1":
		DisplayScene2FightChoice()
		if combat.Combat(p, combat.MonsterDefs["Shaman"]) {
			DisplayScene2ShamanDefeated()
			p.Progression = 3
			DisplayPause()
		}
	case "2":
		if p.Money >= 30 {
			p.Money -= 30
			DisplayScene2PaidShaman()
			p.Progression = 3
			DisplayPause()
		} else {
			DisplayScene2CantPay()
			if combat.Combat(p, combat.MonsterDefs["Shaman"]) {
				DisplayScene2ShamanDefeated()
				p.Progression = 3
				DisplayPause()
			}
		}
	}
}

func sceneSora(p *character.Player) {
	DisplayScene3Intro(p.Name)
	if combat.Combat(p, combat.MonsterDefs["Sora"]) {
		DisplayScene3Victory()
		p.Progression = 4
	}
}
