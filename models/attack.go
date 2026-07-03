package models

// La structure des attaques + la map des attaquesspé

type Attack struct {
	Name     string
	Damage   int
	ManaCost int
}

var AttackDefs = map[string]Attack{
    "Sabre":                       {Name: "Sabre", Damage: 28, ManaCost: 15},
    "Décharge d'énergie":          {Name: "Décharge d'énergie", Damage: 31, ManaCost: 20},
    "Potion de dégâts instantané": {Name: "Potion de dégâts instantané", Damage: 38, ManaCost: 20},
    "Katana":                      {Name: "Katana", Damage: 38, ManaCost: 15},
}