package main

import "fmt"

// Create new type of 'deck'
type deck []string

// Best practice to refer to receiver with single letter of type
// Same equivalent of using "this" in another language
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) shuffle() {

}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
