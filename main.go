package main

func newCard() string {
	return "Ace of Spades"
}

func main() {
	cards := deck{newCard(), newCard()}
	cards = append(cards, "Another One")

	cards.print()
}
