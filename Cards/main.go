package main

import (
	"fmt"
)

func main() {

	fmt.Println("Original Deck")
	myCards := newDeck()
	myCards.print()

	fmt.Println("New Hand and Deck")
	hand, myCards := deal(myCards, 5)

	fmt.Println("Delt Hand:")
	hand.print()

	fmt.Println("Remaining Deck:")
	myCards.print()

	fmt.Println("Shuffled Deck:")
	myCards.shuffle()
	myCards.print()
}
