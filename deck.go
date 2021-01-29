package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suits := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suits)
		}
	}

	return cards
}

func (deck deck) print() {
	for i, card := range deck {
		fmt.Println(i+1, card)
	}
}

func (deck deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range deck {
		newPosition := r.Intn(len(deck) - 1)

		deck[i], deck[newPosition] = deck[newPosition], deck[i]
	}
}

func dealHand(deck deck, handSize int) (deck, deck) {
	return deck[:handSize], deck[handSize:]
}

func (deck deck) toString() string {
	return strings.Join(deck, ",")
}

func (deck deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(deck.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-1)
	}

	deck := strings.Split(string(bs), ",")
	return deck
}

func cleanUppDecks() {
	os.Remove("my_cards")
	os.Remove("my_hand")
}
