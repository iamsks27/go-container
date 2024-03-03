package main

import "fmt"

func main() {
	cards := newDeck()

	cards.saveToFile("cards.txt")

	deck := newDeckFromFile("cards.txt")

	deck.print()

	deck.shuffle()
	fmt.Println("=========================")

	deck.print()
}
