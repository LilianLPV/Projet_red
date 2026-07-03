package models

// Structure des monstres avec leurs stats perso

type Monster struct {
	Name string
	Stats
	XPReward int
}

var MonsterDefs = map[string]Monster{
	"Shaman":        {Name: "Shaman", Stats: Stats{HP: 80, HPMax: 80, StunChance: 3, Damage: 34, Speed: 8, Dodge: 6}, XPReward: 50},
	"Garde":         {Name: "Garde", Stats: Stats{HP: 140, HPMax: 140, StunChance: 2, Damage: 24, Speed: 4, Dodge: 3}, XPReward: 50},
	"Golem":         {Name: "Golem", Stats: Stats{HP: 200, HPMax: 200, StunChance: 1, Damage: 18, Speed: 8, Dodge: 6}, XPReward: 50},
	"Rat-agile":     {Name: "Rat-agile", Stats: Stats{HP: 70, HPMax: 70, StunChance: 2, Damage: 22, Speed: 9, Dodge: 25}, XPReward: 50},
	"Moine renégat": {Name: "Moine renégat", Stats: Stats{HP: 100, HPMax: 100, StunChance: 15, Damage: 26, Speed: 6, Dodge: 6}, XPReward: 50},
	"Mannequin":     {Name: "Mannequin d'entraînement", Stats: Stats{HP: 300, HPMax: 300, Damage: 0, Speed: 1, Dodge: 0, StunChance: 0}, XPReward: 75},
	"Sora":          {Name: "Sora", Stats: Stats{HP: 180, HPMax: 180, Damage: 30, Speed: 7, Dodge: 12, StunChance: 8}, XPReward: 100},
}
