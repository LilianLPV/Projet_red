package models

// Structure du joueur 

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
	Inventory    []Item
	InventoryMax int
	Equipment    Equipment
}
