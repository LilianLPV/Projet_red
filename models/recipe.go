package models

//Strucutre des recettes des armures / équipements

type Recipe struct {
	Name      string
	Cost      int
	Materials map[string]int 
	Slot      string         
	Apply     func(p *Player) 
}

var RecipeDefs = []Recipe{
	{
		Name:      "Menpo",
		Cost:      15,
		Materials: map[string]int{"Fer": 1, "Cuir": 1},
		Slot:      "Head",
		Apply: func(p *Player) {
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
		Apply: func(p *Player) {
			p.HPMax += 12
			p.HP += 12
		},
	},
	{
		Name:      "Waraji",
		Cost:      20,
		Materials: map[string]int{"Cuir": 2, "Bois": 1},
		Slot:      "Feet",
		Apply: func(p *Player) {
			p.Dodge += 3
		},
	},
}