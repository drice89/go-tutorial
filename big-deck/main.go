package main

func main() {

	// deck defined in deck.go - does not need to be imported because it is in the same package
	cards := newDeck()
	cards.showDeck()
}
