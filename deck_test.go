package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck.Cards) != 44 {
		t.Errorf("Expected deck length to be 52, but got: %v", len(deck.Cards))
	}

	if deck.Cards[0].Suit != "Spades" {
		t.Errorf("Expected first card suit to be Ace, but got: %v", deck.Cards[0].Suit)
	}

	if deck.Cards[len(deck.Cards)-1].Value != 10 {
		t.Errorf("Expected last card s value to be 10, but got: %v", deck.Cards[len(deck.Cards)-1].Value)
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	const testingDeck = "_testingDeck"
	_ = os.Remove(testingDeck)

	deck := newDeck()
	_ = deck.saveToFile(testingDeck)

	loadedDeck := newDeckFromFile(testingDeck)

	if len(loadedDeck.Cards) != 44 {
		t.Errorf("Expected deck length to be 52, but got: %v", len(loadedDeck.Cards))
	}

	_ = os.Remove(testingDeck)
}

func TestDealHand(t *testing.T) {
	deck := newDeck()

	hand, remainingCards := dealHand(deck, 12)

	if len(hand.Cards) != 12 {
		t.Errorf("Expected hand size to be 10, but was: %v", len(hand.Cards))

	}

	if len(remainingCards.Cards) != 32 {
		t.Errorf("Expected remainging Cards size to be 40, but was: %v", len(remainingCards.Cards))
	}
}
