package main

import "fmt"

func main() {
	// static type like java
	//var card string = "Ace of Spades"

	// dynamic type like javascript
	// only := when initilzing the var for the first time
	// card := "Ace of Spades"

	// card = newCard()

	// fmt.Println(card)

	cards := []string{newCard(), newCard()}

	fmt.Println(cards)
}

// must put return type like java
func newCard() string {

	return "Five of Diamonds"
}