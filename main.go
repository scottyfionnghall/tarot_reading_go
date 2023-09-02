package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
)

type Deck struct {
	Deck []Card `json:"deck"`
}

type Card struct {
	Name     string `json:"name"`
	Upright  string `json:"upright"`
	Reversed string `json:"reversed"`
}

func generateBool() bool {
	n := rand.Intn(2)
	if n == 1 {
		return true
	} else {
		return false
	}
}

func (card Card) reading() {
	fmt.Println("\tCard name:", card.Name)
	if generateBool() {
		fmt.Println("\tUpright:", card.Upright)
	} else {
		fmt.Println("\tReversed:", card.Reversed)
	}
}

func three_spread(deck Deck) {
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
		card.reading()
	}
}

func main() {
	jsonFile, err := os.Open("tarot.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	var deck Deck
	json.Unmarshal(byteValue, &deck)
	three_spread(deck)
}
