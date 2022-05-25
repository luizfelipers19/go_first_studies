package main

type deck []string

func newDeck() deck {
	cards := deck{}
	cardsSuits := [4]string{"Spades", "Cubs", "Hearts", "Diamonds"}
	cardsValues := [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for i := range cardsValues {
		for j := range cardsSuits {
			cards = append(cards, cardsValues[i]+" of "+cardsSuits[j])
		}
	}
	return cards
}

func dealHand(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) print() {
	for i, card := range d {
		println(i, card)
	}
}
