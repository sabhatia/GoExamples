package main

import (
	"fmt"
	"math/rand"
	"time"
)

type deck []string

func newDeck() deck {

	var cardType = []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	var cardValue = []string{"Ace", "King", "One", "Two"}

	cardDeck := deck{}

	for _, cardVal := range cardValue {
		for _, cardSuite := range cardType {
			cardDeck = append(cardDeck, cardVal+" of "+cardSuite)
		}
	}

	return cardDeck
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

/*
* Given a deck, shuffle it and add some random order to it
 */
func (d deck) shuffle() {
	// Get a new Source
	src := rand.NewSource(time.Now().UTC().UnixNano())

	// Get a new Rand
	randomizer := rand.New(src)

	// shuffle the cards in the deck
	for i := range d {
		newIndx := randomizer.Intn(len(d))
		d[i], d[newIndx] = d[newIndx], d[i]
	}

}
