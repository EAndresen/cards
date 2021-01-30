package main

import (
	"fmt"
	"strings"
)

func main() {
	const firstDeckName = "first_deck"
	const secondDeckName = "second_deck"
	const handSize = 7

	firstDeck := newDeck()
	firstDeck.saveToFile(firstDeckName)
	secondDeck := newDeck()
	secondDeck.saveToFile(secondDeckName)

	fmt.Println("--- New Player - Paquito ---")
	paquito := createPlayer("Paquito", "Navarro")
	paquito.deck = firstDeck
	paquito.deck.shuffle()
	paquito.hand, paquito.deck = dealHand(paquito.deck, handSize)
	fmt.Printf("%+v \n", paquito)

	fmt.Println("--- New Player - Gaben ---")
	gaben := createPlayer("Gaben", "Newell")
	gaben.deck = secondDeck
	gaben.deck.shuffle()
	gaben.hand, paquito.deck = dealHand(paquito.deck, handSize)
	fmt.Printf("%+v \n", gaben)

	fmt.Println("--- AND THE BIG WINNER IS!")
	winner := playCards(paquito, gaben)
	fmt.Print(strings.ToUpper(winner.firstName))

	cleanUppDecks()
}
