package main

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	// append: doesn't modify the existing slice, create a new slice and assign to cards
	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds" 
}