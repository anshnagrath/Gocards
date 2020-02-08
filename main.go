package main

import "log"

func main() {
	cards := newDeck()
	hand, _ := deal(cards, 5)
	err := hand.saveToFile("my hand")
	if err != nil {
		log.Panic(err)
	}
	newCards := newDeckFromFile("my hand")
	newCards.shuffleCards()
	newCards.print()

}
