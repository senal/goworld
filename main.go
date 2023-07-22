package main

func main() {

	HAND_SIZE := 5
	cards := newDeck()
	hand, remainingDeck := deal(cards, HAND_SIZE)

	hand.print()
	remainingDeck.print()
}
