package main

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	new_deck := newDeck()

	// assuming we have 4 cardsuite and 4 cardvalues
	fmt.Println(len(new_deck))
	if len(new_deck) != 16 {
		t.Errorf("Expected deck length 16, but received %v", len(new_deck))
	}

	if new_deck[0] != "Hearts of Ace" {
		t.Errorf("Expected first deck as Hearts of Ace, but got %v", new_deck[0])
	}

	if new_deck[len(new_deck)-1] != "Clubs of Four" {
		t.Errorf("Expected last deck as Clubs of Four, but got %v", new_deck[len(new_deck)-1])
	}
}

// Next test - Delete any file in current working directory with the name "_decktesting"
// create a deck, save to file "_decktesting", load/read from file, assert deck length
// delete any files in current working directory with the name "_decktesting"

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// remove the existing file
	os.Remove("_decktesting")

	// create a new deck and save it
	d := newDeck()
	d.saveToFile("_decktesting")

	// read the deck from local file
	loadedDeck := newDeckFromFile("_decktesting")

	// assertions
	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, but got %v", len(loadedDeck))
	}

	// clean up the file
	os.Remove("_decktesting")
}
