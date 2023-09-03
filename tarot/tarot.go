package tarot

import (
	"fmt"
	"math/rand"
)

// Custom type for tarot deck
type Deck struct {
	Deck []Card `json:"deck"`
}

// Custom type for every card, contains name, upright keywords, reversed keywords
// form imported json file
type Card struct {
	Name     string `json:"name"`
	Upright  string `json:"upright"`
	Reversed string `json:"reversed"`
}

// Function that used to randomly select between Upright and Reversed meanings
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

// Function to do a three card spread.
// User chooses what type of spread would they like, then for each word we
// randomly select a card and print it out
func ThreeSpread(deck Deck) {
	fmt.Println(`Which spread would you like to receive?
	1. Past, Present, Future
	2. Situation, Obstacle, Advice
	3. Mind, Body, Spirit`)
	type_spread := [][]string{
		{
			"Past", "Present", "Future",
		}, {
			"Situation", "Obstacle", "Advice",
		}, {
			"Mind", "Body", "Spirit",
		},
	}
	var choice int
	fmt.Scanf("%d", &choice)
	choice -= 1
	for i := 0; i <= 2; i++ {
		fmt.Println(type_spread[choice][i])
		card := deck.Deck[rand.Intn(len(deck.Deck))]
		card.Reading()
	}
}
