package main

import "fmt"

func main() {
	const myCards = "my_cards"
	const myHand = "my_hand"
	const handSize = 7

	cards := newDeck()
	cards.saveToFile(myCards)

	fmt.Println("--- New Deck---")

	cardsLoaded := newDeckFromFile(myCards)
	cardsLoaded.print()

	fmt.Println("--- New Deck Shuffled ---")

	cardsLoaded.shuffle()
	cardsLoaded.print()

	fmt.Println("--- My Hand ---")

	hand, remainingCards := dealHand(cardsLoaded, handSize)
	hand.saveToFile(myHand)
	hand.print()

	fmt.Println("--- Remaining Cards ---")

	remainingCards.shuffle()
	remainingCards.print()

	cleanUppDecks()
}
