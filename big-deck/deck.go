package main

import "fmt"

// create a new type of deck

// slice of type string. Note that append is still working on this
type deck []string

func newDeck() deck {
	// creates new variable cards of type deck (a slice of strings)
	cards := deck{}
	// create suits
	cardSuits := [4]string{"Hearts", "Diamonds", "Spades", "Clubs"}
	cardValues := [13]string{
		"Ace",
		"Two",
		"Three",
		"Four", 
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine", 
		"Ten",
		"Jack",
		"Queen",
		"King",
	}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}
	return cards
}

func (d deck) showDeck() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}