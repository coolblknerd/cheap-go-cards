package main

import "fmt"

func newCard() string {
	return "Ace of Spades"
}

func main() {
	cards := []string{newCard(), newCard()}
	cards = append(cards, "Another One")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}
