package main

import (
	"fmt"
	"math/rand"
)

func main() {
	suits := [4]string{"♠️", "♥️", "♦️", "♣️"}
	values := [13]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	var deck [52]string

	i := 0

	for _, suit := range suits {
		for _, value := range values {
			deck[i] = fmt.Sprintf("%s%s", value, suit)
			i++
		}
	}

	fmt.Println("Original Deck: ", deck)

	for i := len(deck) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}

	fmt.Println("Shuffled Deck: ", deck)
}
