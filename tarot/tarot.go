package tarot

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
	d "tarot_reading/deck"
)

// Function to create Deck object
func CreateDeck() d.Deck {
	// Read json file with all the card info
	jsonFile, err := os.Open("tarot.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	// Create a variable of custom type Deck,
	// imported from module "tarot_reading/deck"
	var deck d.Deck
	json.Unmarshal(byteValue, &deck)
	return deck
}

// Function to do a three card spread.
// User chooses what type of spread would they like, then for each word we
// randomly select a card and print it out
func ThreeSpread(deck d.Deck) {
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
