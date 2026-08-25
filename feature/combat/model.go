package combat

import "Projet-Red/feature/character"

// CombatState garde les effets en cours (étourdissement, poison) et le numéro du tour.
type CombatState struct {
	PlayerStun    bool
	EnemyStun    bool
	EnemyPoisons int
	Tour             int
}

// Monster représente un ennemi affrontable en combat.
type Monster struct {
	Name string
	character.Stats
	XPReward int
}

// MonsterDefs contient tous les monstres du jeu.
var MonsterDefs = map[string]Monster{
	"Shaman":        {Name: "Shaman", Stats: character.Stats{HP: 80, HPMax: 80, StunChance: 3, Damage: 34, Speed: 8, Dodge: 6}, XPReward: 50},
	"Garde":         {Name: "Garde", Stats: character.Stats{HP: 140, HPMax: 140, StunChance: 2, Damage: 24, Speed: 4, Dodge: 3}, XPReward: 50},
	"Golem":         {Name: "Golem", Stats: character.Stats{HP: 200, HPMax: 200, StunChance: 1, Damage: 18, Speed: 8, Dodge: 6}, XPReward: 50},
	"Rat-agile":     {Name: "Rat-agile", Stats: character.Stats{HP: 70, HPMax: 70, StunChance: 2, Damage: 22, Speed: 9, Dodge: 25}, XPReward: 50},
	"Moine renégat": {Name: "Moine renégat", Stats: character.Stats{HP: 100, HPMax: 100, StunChance: 15, Damage: 26, Speed: 6, Dodge: 6}, XPReward: 50},
	"Mannequin":     {Name: "Mannequin d'entraînement", Stats: character.Stats{HP: 300, HPMax: 300, Damage: 0, Speed: 1, Dodge: 0, StunChance: 0}, XPReward: 75},
	"Sora":          {Name: "Sora", Stats: character.Stats{HP: 180, HPMax: 180, Damage: 30, Speed: 7, Dodge: 12, StunChance: 8}, XPReward: 100},
}

// Attack représente une attaque spéciale (compétence de classe).
type Attack struct {
	Name     string
	Damage   int
	ManaCost int
}

// AttackDefs contient toutes les attaques spéciales du jeu.
var AttackDefs = map[string]Attack{
	"Sabre":                       {Name: "Sabre", Damage: 28, ManaCost: 15},
	"Décharge d'énergie":          {Name: "Décharge d'énergie", Damage: 31, ManaCost: 20},
	"Potion de dégâts instantané": {Name: "Potion de dégâts instantané", Damage: 38, ManaCost: 20},
	"Katana":                      {Name: "Katana", Damage: 38, ManaCost: 15},
}
