package main

func main() {
	// static type like java
	// var card string = "Ace of Spades"

	// dynamic type like javascript
	// only := when initilzing the var for the first time
	// card := "Ace of Spades"

	// card = newCard()

	// fmt.Println(card)

	cards := deck{newCard(), newCard()}
	cards = append(cards, "Six of Spades")

	// range is a keyword when we want to iterate through the slice
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	cards.print()
}

// must put return type like java
func newCard() string {

	return "Five of Diamonds"
}