package main

func main() {
	// static type like java
	// var card string = "Ace of Spades"

	// dynamic type like javascript
	// only := when initilzing the var for the first time
	// card := "Ace of Spades"

	// card = newCard()

	// fmt.Println(card)

	//cards := deck{newCard(), newCard()}
	//cards = append(cards, "Six of Spades")

	// range is a keyword when we want to iterate through the slice
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// Type coversion in go, type we want(value we have)
	// []byte("Hello World")
	// greeting := "hello"
	// fmt.Println([]byte(greeting))

	cards := newDeck()

	cards.saveToFile("my_cards")

	// cards.print()

	// we are initizling and assigning at the sametime from a func return val
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()
}

// must put return type like java
// func newCard() string {

// 	return "Five of Diamonds"
// }