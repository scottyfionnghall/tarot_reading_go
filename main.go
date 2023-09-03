package main

import (
	"fmt"
	"tarot_reading/tarot"
)

func main() {
	deck := tarot.CreateDeck()
	tarot.ThreeSpread(deck)
	var input string
	fmt.Scanln(&input)
}
