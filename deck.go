package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

// methods that can be used on every variable of type deck
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"spade", "diamond", "heart", "club"}
	cardValues := []string{"Ace", "two", "three", "four"}
	for _, suit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+suit)
		}
	}
	return cards
}
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), " ,")
}

// ioutil.WriteFile takesthree arguments filename , []bytes and permissions (to execute the file or read)
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
