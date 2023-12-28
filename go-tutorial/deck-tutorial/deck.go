package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func deal(d deck, handSize int) (deck, deck) {
	// multiple return values
	// first is from beginning to handSize
	// second is from handSize to the end
	// ex [1,2,3,4], 2
	// [1,2], [3,4]
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err == nil && bs != nil {
		s := strings.Split(string(bs), ",")
		return deck(s)
	} else {
		fmt.Println("ERROR:", err)
		fmt.Println("INFO: Generating new deck due to error loading deck file")
		return newDeck()
		// alternatively we could exit
		//os.Exit(2)
	}
}

func (d deck) shuffle() deck {
	// base rand seed on system clock
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
	return d
}