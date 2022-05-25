package main

func main() {
	//cards := newDeck()
	//cards.print()
	//cards.print()

	//fmt.Println(cards.toString())
	//cards.saveToFile("myCards")
	cards := newDeck()
	cards.shuffleCards()
	cards.print()

	//cards = append(cards, "Six of Spades")
	//
	//cards.print()
	//
	//fmt.Println(cards)

	////trying string to byte conversion
	//greeting := "Hi there!"
	//fmt.Println([]byte(greeting))

}

func newCard() string {
	return "Five or diamonds"
}
