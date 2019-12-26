package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// option #1 - log the error and return a call to newDeck()
		// option #2 -  log the error and entierly quit the program
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",") // string conversion of byte
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
