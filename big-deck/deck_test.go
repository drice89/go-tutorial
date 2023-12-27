package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 20 but got %v", len(d))
	}

	assert.Equal(t, d[0], "Ace of Hearts")
	assert.Equal(t, d[51], "King of Clubs")
}