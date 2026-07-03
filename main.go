package main

import (
	"Projet-Red/game"
)

func main() {
	p := game.CharCreation()

	game.MainMenu(&p)
}
