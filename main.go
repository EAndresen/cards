package main

import "fmt"

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")

	fmt.Println("--- New Deck---")

	cardsLoaded := newDeckFromFile("my_cards")
	cardsLoaded.print()

	fmt.Println("--- New Deck Shuffled ---")

	cardsLoaded.shuffle()
	cardsLoaded.print()

	fmt.Println("--- My Hand ---")

	hand, remainingCards := dealHand(cardsLoaded, 7)
	hand.saveToFile("my_hand")
	hand.print()

	fmt.Println("--- Remaining Cards ---")

	remainingCards.shuffle()
	remainingCards.print()

	cleanUppDecks()
}
