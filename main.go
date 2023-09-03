package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"tarot_reading/tarot"
)

func main() {
	// Read json file with all the card info
	jsonFile, err := os.Open("tarot.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	// Create a variable of custom type Deck,
	// imported from module "tarot/reading/taort"
	var deck tarot.Deck
	json.Unmarshal(byteValue, &deck)
	// Do a three card spread
	tarot.ThreeSpread(deck)
	// Wait for user input to close the program
	var input string
	fmt.Scanln(&input)
}
