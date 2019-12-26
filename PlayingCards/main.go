package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()
	fmt.Println("*************************")
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()

	//byte slice example
	greetings := "Hello, how are you ?"
	fmt.Println([]byte(greetings))
	// end of example

	// using of strings join
	fmt.Println(cards.toString())
	// end

	// save file to disk
	cards.saveToFile("my_file")
	// end

	// read file from disk
	fmt.Println("********* Load file from disk*******")
	cardsFromFile := newDeckFromFile("my_file")
	cardsFromFile.print()
	//end

	// shuffle cards
	fmt.Println("********* shuffle cards*******")
	cards.shuffle()
	cards.print()
	// end
}
