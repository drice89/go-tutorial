package main

// auto imported... coool!
import "fmt"

func main() {

	// deck defined in deck.go - does not need to be imported because it is in the same package
	cards := deck{
		newCard(),
		"Ace of Diamonds",
	}
	cards = append(cards, "Six of spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
