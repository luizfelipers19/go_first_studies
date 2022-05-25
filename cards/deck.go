package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func (d deck) toString() string {
	return strings.Join([]string(d), ",")

}

func (d deck) saveToFile(filename string) error {

	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func deckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(byteSlice), ",")
	return deck(s)
}

func (d deck) shuffleCards() {
	sourceValue := rand.NewSource(time.Now().UnixNano())
	r := rand.New(sourceValue)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
