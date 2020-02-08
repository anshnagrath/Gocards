package main

import "log"

func main() {
	cards := newDeck()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
	err := hand.saveToFile("my hand")
	if err != nil {
		log.Panic(err)
	}

}
