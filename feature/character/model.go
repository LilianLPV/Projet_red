package character

import "Projet-Red/feature/inventory"

// Stats regroupe les caractéristiques vitales communes au joueur et aux monstres.
type Stats struct {
	HPMax      int
	HP         int
	ManaMax    int
	Mana       int
	Damage     int
	Speed      int
	StunChance int
	Dodge      int
}

// Equipment représente les emplacements d'armure du joueur.
type Equipment struct {
	Head  *inventory.Item
	Torso *inventory.Item
	Feet  *inventory.Item
}

// Player représente le personnage incarné par le joueur.
type Player struct {
	Name                   string
	Class                  string
	Level                  int
	XP                     int
	XPMax                  int
	Money                  int
	DejaRessuscite         bool
	PotionGratuiteUtilisee bool
	Progression            int
	Stats
	BaseAttack   string
	Skill        []string
	Inventory    []inventory.Item
	InventoryMax int
	Equipment    Equipment
}

// classDef décrit les statistiques et l'attaque de base de départ d'une classe.
type classDef struct {
	Stats      Stats
	BaseAttack string
}

// classDefs définit les 4 classes jouables.
var classDefs = map[string]classDef{
	"Soldat": {
		Stats: Stats{
			HPMax: 110, HP: 110,
			ManaMax: 30, Mana: 30,
			Damage: 22, Speed: 7,
			StunChance: 3, Dodge: 8,
		},
		BaseAttack: "Poing",
	},
	"Moine": {
		Stats: Stats{
			HPMax: 100, HP: 100,
			ManaMax: 50, Mana: 50,
			Damage: 25, Speed: 6,
			StunChance: 20, Dodge: 8,
		},
		BaseAttack: "Poing",
	},
	"Alchimiste": {
		Stats: Stats{
			HPMax: 70, HP: 70,
			ManaMax: 40, Mana: 40,
			Damage: 30, Speed: 8,
			StunChance: 3, Dodge: 18,
		},
		BaseAttack: "Caillou",
	},
	"Samouraï": {
		Stats: Stats{
			HPMax: 150, HP: 150,
			ManaMax: 30, Mana: 30,
			Damage: 18, Speed: 4,
			StunChance: 5, Dodge: 3,
		},
		BaseAttack: "Caillou",
	},
}

// LevelMax est le niveau maximum atteignable.
const LevelMax = 5
