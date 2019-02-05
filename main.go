package main

func newCard() string {
	return "Ace of Spades"
}

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
