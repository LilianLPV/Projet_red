package main

import (
	"Projet-Red/game"
)

func main() {
	p := game.CharCreation()
	game.IntroLore(&p)
	game.StartMenu(&p)
	game.MainMenu(&p)
}
