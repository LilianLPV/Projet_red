package blacksmith

import "Projet-Red/feature/character"


type Recipe struct {
	Name      string
	Cost      int
	Materials map[string]int
	Slot      string
	Apply     func(p *character.Player)
}

// RecipeDefs contient toutes les recettes forgeables chez Thorin.
var RecipeDefs = []Recipe{
	{
		Name:      "Menpo",
		Cost:      15,
		Materials: map[string]int{"Fer": 1, "Cuir": 1},
		Slot:      "Head",
		Apply: func(p *character.Player) {
			p.HPMax += 3
			p.HP += 3
			p.StunChance += 2
		},
	},
	{
		Name:      "Tatami-do",
		Cost:      30,
		Materials: map[string]int{"Acier": 2, "Bois": 1, "Cuir": 1},
		Slot:      "Torso",
		Apply: func(p *character.Player) {
			p.HPMax += 12
			p.HP += 12
		},
	},
	{
		Name:      "Waraji",
		Cost:      20,
		Materials: map[string]int{"Cuir": 2, "Bois": 1},
		Slot:      "Feet",
		Apply: func(p *character.Player) {
			p.Dodge += 3
		},
	},
}
