package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{
		"Ace",
		"Two",
		"Three",
		"Four",
	}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {

	return strings.Join(d, "\n")
}

func (d deck) saveToFile(filename string) error {

	return os.WriteFile(filename, []byte(d.toString()), 0644)
}

func newDeckFromFile(filename string) deck {
	content, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("There was an error reading the file: ", err)
		panic(err)
	}

	deck := strings.Split(string(content), "\n")
	return deck
}

func (d deck) shuffle() {
	// Randomize the order of the cards
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPos := r.Intn(len(d))
		d[i], d[newPos] = d[newPos], d[i]
	}
}
