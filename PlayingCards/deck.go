package main

import "fmt"

// Create  a new type of 'deck'
// which is a slice of strings
type deck []string

// Create and return a list of playing cards,

func newDeck() deck {

	cards := deck{}

	cardsSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardsValues := []string{"Ace", "Two", "Three", "Four"}
	// replaceing i and j with _ as we are not using them just a declaration
	for _, suit := range cardsSuits {
		for _, value := range cardsValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// receiver on a function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}

}
