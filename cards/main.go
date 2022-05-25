package main

func main() {
	cards := newDeck()
	cards.print()

	//cards = append(cards, "Six of Spades")
	//
	//cards.print()
	//
	//fmt.Println(cards)

}

func newCard() string {
	return "Five or diamonds"
}
