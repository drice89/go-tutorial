package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 20 but got %v", len(d))
	}

	assert.Equal(t, d[0], "Ace of Hearts")
	assert.Equal(t, d[51], "King of Clubs")
}

func TestSaveDeckToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckfile")
	d := newDeck()
	d.shuffle()
	d.saveToFile("_deckfile")
	
	loadedDeck := newDeckFromFile("_deckfile")

	assert.Equal(t, len(loadedDeck), 52, "loaded deck is not correct length")
}