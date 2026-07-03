package models

// Initialisation des classes 

var ClassDefs = map[string]struct {
	Stats      Stats
	BaseAttack string
}{
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

func Init(nom string, classe string) Player {
	def, ok := ClassDefs[classe]
	if !ok {
		return Player{}
	}
	return Player{
		Name:         nom,
		Class:        classe,
		Level:        1,
		XP:           0,
		XPMax:        100,
		Money:        100,
		Stats:        def.Stats,
		BaseAttack:   def.BaseAttack,
		InventoryMax: 10,
		Inventory:    []Item{},
	}
}
