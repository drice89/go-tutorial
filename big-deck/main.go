package main

func main() {

	// deck defined in deck.go - does not need to be imported because it is in the same package
	cards := newDeckFromFile("my_c")
	// cards.showDeck()

	// hand, remainingDeck := deal(cards, 5)
	cards.showDeck()
}
