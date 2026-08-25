package character

import (
	"Projet-Red/feature/inventory"
	"Projet-Red/utils"
)

// Init crée un nouveau joueur pour la classe donnée, avec ses stats de départ.
func Init(name string, classe string) Player {
	def, ok := classDefs[classe]
	if !ok {
		return Player{}
	}
	return Player{
		Name:         name,
		Class:        classe,
		Level:        1,
		XP:           0,
		XPMax:        100,
		Money:        100,
		Stats:        def.Stats,
		BaseAttack:   def.BaseAttack,
		InventoryMax: 10,
		Inventory:    []inventory.Item{},
	}
}

// classFromChoice traduit le choix numéroté du menu en nom de classe.
func classFromChoice(choice string) string {
	switch choice {
	case "1":
		return "Soldat"
	case "2":
		return "Moine"
	case "3":
		return "Alchimiste"
	case "4":
		return "Samouraï"
	}
	return ""
}

// CharCreation demande son nom et sa classe au joueur, puis crée son personnage.
func CharCreation() Player {
	nom := utils.ReadName("Entre ton nom : ")
	displayClassChoices()
	choice := utils.ReadChoice("Ton choix : ", []string{"1", "2", "3", "4"})
	return Init(nom, classFromChoice(choice))
}

// XPWin ajoute de l'expérience au joueur et gère la montée de niveau.
func XPWin(p *Player, montant int) {
	if p.Level >= LevelMax {
		return
	}
	p.XP += montant
	DisplayXPGain(montant)
	if p.XP >= p.XPMax {
		p.XP -= p.XPMax
		p.Level++
		p.XPMax += 50
		p.HPMax += 10
		p.HP = p.HPMax
		p.Damage += 2
		if p.Level >= LevelMax {
			p.XP = 0
			DisplayLevelUp(p.Level, true)
		} else {
			p.XPMax += 50
			DisplayLevelUp(p.Level, false)
		}
	}
}

// CanAfford indique si le joueur possède assez d'or pour payer le coût donné.
func CanAfford(p *Player, cost int) bool {
	return p.Money >= cost
}

// AddItem ajoute un objet à l'inventaire du joueur (affiche un message si le sac est plein).
func AddItem(p *Player, nom string, quantite int) {
	items, ok := inventory.Add(p.Inventory, p.InventoryMax, nom, quantite)
	p.Inventory = items
	if !ok {
		inventory.DisplayFull()
	}
}

// RemoveItem retire un objet de l'inventaire du joueur.
func RemoveItem(p *Player, nom string, quantite int) {
	p.Inventory = inventory.Remove(p.Inventory, nom, quantite)
}

// QuantityOf retourne la quantité d'un objet possédé par le joueur.
func QuantityOf(p *Player, nom string) int {
	return inventory.Quantity(p.Inventory, nom)
}

// DrinkHealthPotion consomme une potion de santé de l'inventaire du joueur.
func DrinkHealthPotion(p *Player) {
	if QuantityOf(p, "Potion de santé") == 0 {
		DisplayNoItem("potion de santé")
		return
	}
	if p.HP > p.HPMax-50 {
		if !PromptWastefulHeal(p.HP, p.HPMax) {
			return
		}
	}
	p.HP = utils.Clamp(p.HP+50, 0, p.HPMax)
	RemoveItem(p, "Potion de santé", 1)
	DisplayHealthPotionUsed(p.HP, p.HPMax)
}

// DrinkManaPotion consomme une potion de mana de l'inventaire du joueur.
func DrinkManaPotion(p *Player) {
	if QuantityOf(p, "Potion de mana") == 0 {
		DisplayNoItem("potion de mana")
		return
	}
	if p.Mana > p.ManaMax-20 {
		if !PromptWastefulMana(p.Mana, p.ManaMax) {
			return
		}
	}
	p.Mana = utils.Clamp(p.Mana+20, 0, p.ManaMax)
	RemoveItem(p, "Potion de mana", 1)
	DisplayManaPotionUsed(p.Mana, p.ManaMax)
}

// AccessInventory ouvre l'écran d'inventaire du joueur.
func AccessInventory(p *Player) {
	utils.Clear()
	DisplayInventoryScreen(p)
	choice := utils.ReadChoice("Ton choix : ", []string{"1", "2", "3"})
	switch choice {
	case "1":
		DrinkHealthPotion(p)
	case "2":
		DrinkManaPotion(p)
	case "3":
		return
	}
}
