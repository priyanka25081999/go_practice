package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of 'deck'
// which is a slice of string
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuite := []string{"Hearts", "Diamonds", "Shades", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suite := range cardSuite {
		for _, values := range cardValues {
			cards = append(cards, suite+" of "+values)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, size int) (deck, deck) {
	// return start-to-size and size-to-end (2 values)
	return d[:size], d[size:]
}

// helper function
// convert the deck into string type, it takes receiver and returns string
func (d deck) deckToString() string {

	// ["hearts of ace", "clubes of ace"...] - slice of string
	// ["hearts of ace, clubes of ace, ..."] - string type (single string)
	// Join method takes the slice of strings and then we pass separator
	// the join function will create a single string separated by comma
	return strings.Join([]string(d), ", ")
}

func (d deck) saveToFile(file_name string) error {
	// 0666 -> read and write permission
	return ioutil.WriteFile(file_name, []byte(d.deckToString()), 0666)
}

func newDeckFromFile(file_name string) deck {
	bs, err := ioutil.ReadFile(file_name) // returns the []byte and error

	// error occurs
	if err != nil {
		fmt.Println("Error : ", err)
		// 1 - not success
		os.Exit(1)
	}

	// convert byte slice into the string and then convert it into string slice using split method
	str := strings.Split(string(bs), ",")
	// convert str (slice of strings) into deck
	return deck(str)
}

func (d deck) shuffle() {
	// here, we use current time as the value of seed to generate random number
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for index := range d {
		// generate a random number
		new_index := r.Intn(len(d) - 1)

		// swap the current position with new position
		d[index], d[new_index] = d[new_index], d[index]
	}
}
