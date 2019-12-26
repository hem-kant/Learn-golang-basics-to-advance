package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	enB := englishBot{}
	esB := spanishBot{}

	printGreeting(enB)
	printGreeting(esB)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi There !!"
}

func (spanishBot) getGreeting() string {
	return "Hola!!"
}
