package inventory

// Item représente un objet possédé par le joueur (potion, matériau, équipement...).
type Item struct {
	Name        string
	Quantity    int
	Description string
}
