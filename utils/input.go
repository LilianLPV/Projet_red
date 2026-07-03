package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// Lire les inputs dit sur le clavier

func ReadInput() string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	return line
}

// Lire choix grâce au inputs

func ReadChoice(prompt string, valides []string) string {
	for {
		fmt.Print(prompt)
		entree := ReadInput()
		for _, v := range valides {
			if entree == v {
				return entree
			}
		}
		fmt.Println(Red + "Choix invalide, réessaie" + Reset)
	}
}

// Le nom du personnage

func ReadName(prompt string) string {
	for {
		fmt.Print(prompt)
		nom := ReadInput()

		if nom == "" {
			fmt.Println(Red + "Le nom ne peut pas être vide !" + Reset)
			continue
		}
		valide := true
		for _, r := range nom {
			if !unicode.IsLetter(r) {
				fmt.Println(Red + "Le nom doit contenir que des lettres !" + Reset)
				valide = false
				break
			}
		}
		if !valide {
			continue
		}
		return FormaterName(nom)
	}
}

// Formater le nom 

func FormaterName(s string) string {
	bas := strings.ToLower(s)
	r := []rune(bas)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

