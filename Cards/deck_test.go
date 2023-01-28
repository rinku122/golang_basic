package main

import (
	"os"
	"testing"
)

func TestNewPacks(t *testing.T) {
	deck := newPacks()

	if len(deck) != 16 {
		t.Errorf("Expected length is 16 but got %v", len(deck))
	}
}

func TestNewPacksElements(t *testing.T) {
	deck := newPacks()

	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected first card to Ace of Spades but got but %v", deck[0])
	}

	if deck[len(deck)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to Ace of Spades but got but %v", deck[len(deck)-1])
	}
}

func TestSaveToFile(t *testing.T) {
	os.Remove("_test")

	cards := newPacks()

	cards.saveToFile("_test")

	cards = readFromDrive("_test")

	if len(cards) != 16 {
		t.Errorf("Error in Reading File")
	}

	os.Remove("_test")

}
