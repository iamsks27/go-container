package main

func main() {
	cards := newDeck()

	cards.saveToFile("cards.txt")

	deck := newDeckFromFile("cards.txt")

	deck.print()
}
