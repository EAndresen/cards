package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck) != 52 {
		t.Errorf("Expected deck length to be 52, but got: %v", len(deck))
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got: %v", deck[0])
	}

	if deck[len(deck)-1] != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs, but got: %v", deck[len(deck)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	const testingDeck = "_testingDeck"
	_ = os.Remove(testingDeck)

	deck := newDeck()
	_ = deck.saveToFile(testingDeck)

	loadedDeck := newDeckFromFile(testingDeck)

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length to be 52, but got: %v", len(loadedDeck))
	}

	_ = os.Remove(testingDeck)
}

func TestDealHand(t *testing.T) {
	deck := newDeck()

	hand, remainingCards := dealHand(deck, 12)

	if len(hand) != 12 {
		t.Errorf("Expected hand size to be 10, but was: %v", len(hand))

	}

	if len(remainingCards) != 40 {
		t.Errorf("Expected remainging cards size to be 40, but was: %v", len(remainingCards))
	}
}
