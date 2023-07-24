package main

func main() {

	HAND_SIZE := 5
	cards := newDeck()
	hand, remainingDeck := deal(cards, HAND_SIZE)

	hand.print()
	remainingDeck.print()

<<<<<<< HEAD
	cards.print()
	cards.shuffle()
	cards.print()
=======
	cards.saveToFile("my_deck")
	d := newDeckFromFile("my_deckx")
	d.print()
>>>>>>> chapter/convert-string-slice-to-byte-slice
}
