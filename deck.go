package main

import (
	"fmt"
	"strings"
)

// Create a new type of deck
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// Select a specific range from a slice fruits[startIndexIncluding : upToNotIncluding]
	// can also do slice[2:] or slice[:2]

	// whenever we have a variable we dont use in loops, we just replace it with _
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of  " + suit)
		}
	}

	return cards
}

// receiver on a function
// receiver sets up methods on variables that we create
// cards is the actualy copy of the deck we are working with, similar to this keyword
// every variable of type deck has access to this method
// any variable of type deck now has access to the print method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// returning multiple values, this tells us we are turning 2 things of type deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}