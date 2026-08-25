package main

import (
	"Projet-Red/feature/character"
	"Projet-Red/feature/story"
	"Projet-Red/menu"
)

func main() {
	p := character.CharCreation()
	story.IntroLore(&p)
	if story.StartMenu(&p) {
		menu.MainMenu(&p)
	}
}
