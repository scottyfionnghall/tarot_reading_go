package deck

import (
	"fmt"
	"math/rand"
)

// Custom type for tarot deck
type Deck struct {
	Deck []Card `json:"deck"`
}

// Custom type for every card, contains name, upright keywords,
// reversed keywords form imported json file
type Card struct {
	Name     string `json:"name"`
	Upright  string `json:"upright"`
	Reversed string `json:"reversed"`
}

// Function that used by Reading() method to randomly select
// between Upright and Reversed meanings
func generateBool() bool {
	n := rand.Intn(2)
	if n == 1 {
		return true
	} else {
		return false
	}
}

// Card type method used to print out a formatted infromation about a card
// in a form of a reading
func (card Card) Reading() {
	fmt.Println("\tCard name:", card.Name)
	if generateBool() {
		fmt.Println("\tUpright:", card.Upright)
	} else {
		fmt.Println("\tReversed:", card.Reversed)
	}
}
