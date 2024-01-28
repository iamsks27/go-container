package main

import "fmt"

func main() {
	fmt.Println("Making a function call")

	card := newCard()

	fmt.Println(card)
}

func newCard() string {
	return "Ace of Spades"
}
