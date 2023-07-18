package main

import "fmt"

func main() {

	// we use := when initializing for the first time
	card := newCard()

	fmt.Println(card)
}

func newCard() string {
	return "Ace of Spades"
}
