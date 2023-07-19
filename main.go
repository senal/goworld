package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	// append to the list
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, ".", card)
	}
}

func newCard() string {
	return "Ace of Spades"
}
