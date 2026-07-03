package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// Diminuer la vitesse du texte

func SlowPrint(text string, delay time.Duration) {
	delay = 0
	for _, c := range text {
		fmt.Printf("%c", c)
		time.Sleep(delay)
	}
	fmt.Println()
}
// Clear le terminal pour que ce soit propre

func Clear() {
	fmt.Print("\033[H\033[2J")
}

// 

func Clamp(valeur, min, max int) int {
	if valeur < min {
		return min
	}
	if valeur > max {
		return max
	}
	return valeur
}

// Le % de chance 

func Chance(pourcentage int) bool {
	return rand.Intn(100) < pourcentage
}
