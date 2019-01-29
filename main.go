package main

func newCard() string {
	return "Ace of Spades"
}

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 7)

	hand.print()
	remainingDeck.print()
}
