package inventory

import "fmt"

func DisplayFull() {
	fmt.Println("Votre Inventaire est plein !")
}

func DisplayList(items []Item) {
	if len(items) == 0 {
		fmt.Println("Votre Inventaire est vide !")
		return
	}
	for _, i := range items {
		fmt.Println(i.Name, i.Quantity)
	}
}
