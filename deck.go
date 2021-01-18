package main

import "fmt"

type deck []string

func () newDeck() deck {
	cards := deck{"Ace of Diamonds", newCard()}
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}