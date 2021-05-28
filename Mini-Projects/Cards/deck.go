package cards

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Deck []string

//NewDeck will create a new deck of cards
func NewDeck() Deck {
	cards := Deck{}

	cardsuit := []string{"Diamonds", "Spades", "Hearts", "Clubs"}
	cardvalue := []string{"Ace", "one", "two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "King", "Queen", "Jack"}

	for _, suit := range cardsuit {
		for _, value := range cardvalue {
			cards = append(cards, suit+" of "+value)

		}

	}
	return cards

}

//Deal to create a deck of handsize and deck of remaining cards
func Deal(d Deck, handsize int) (Deck, Deck) {
	return d[:handsize], d[handsize:]
}

//Print to print the deck of cards
func (d Deck) Print() {

	for i, value := range d {
		fmt.Println(i, value)
	}

}

func toString(d Deck) string {
	return (strings.Join([]string(d), ","))
}

//SaveToFile to save the deck of cards into a file
func (d Deck) SaveToFile(filename string) error {
	a := toString(d)
	err := ioutil.WriteFile(filename, []byte(a), 0666)

	return err
}

//WriteFromFile to convert the contents of file into a Deck
func WriteFromFile(filename string) Deck {

	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s := string(bs)

	return Deck(strings.Split(s, ","))

}

//Shuffle to shuffle the given deck of cards
func (d Deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {

		newposition := r.Intn(len(d) - 1)

		d[i], d[newposition] = d[newposition], d[i]

	}
}
