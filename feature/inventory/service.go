package inventory

// Add ajoute une quantité d'un item à la liste. Si l'item existe déjà, sa quantité
// est simplement augmentée. Retourne la liste mise à jour et false si le sac est plein.
func Add(items []Item, max int, name string, quantity int) ([]Item, bool) {
	for i := range items {
		if items[i].Name == name {
			items[i].Quantity += quantity
			return items, true
		}
	}
	if len(items) >= max {
		return items, false
	}
	return append(items, Item{Name: name, Quantity: quantity}), true
}

// Remove retire une quantité d'un item de la liste, et le supprime totalement
// si sa quantité tombe à 0 ou moins.
func Remove(items []Item, name string, quantity int) []Item {
	for i := range items {
		if items[i].Name == name {
			items[i].Quantity -= quantity
			if items[i].Quantity <= 0 {
				return append(items[:i], items[i+1:]...)
			}
			return items
		}
	}
	return items
}

// Quantity retourne la quantité possédée d'un item donné (0 si absent).
func Quantity(items []Item, name string) int {
	for _, item := range items {
		if item.Name == name {
			return item.Quantity
		}
	}
	return 0
}

// Has indique si le joueur possède au moins une unité de l'item.
func Has(items []Item, name string) bool {
	return Quantity(items, name) > 0
}
