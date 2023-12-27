package main

// auto imported... coool!
import "fmt"

func main() {
	// var card string = "Ace of Spades"
	cards := []string{
		newCard(),
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
