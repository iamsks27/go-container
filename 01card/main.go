package main

import "fmt"

func main() {
	card := newCard()

	fmt.Println(card)
}

func newCard() string {
	return "Ace of Spades"
}
