package deckOfCards

import (
	"cassidoo/utils"
	"math/rand"
	"slices"
)

type Deck interface {
	Shuffle()
	Draw(n int) []string
}

type deckImpl struct {
	cards []string
}

var clubs = []string{"2♣️", "3♣️", "4♣️", "5♣️", "6♣️", "7♣️", "8♣️", "9♣️", "10♣️", "J♣️", "Q♣️", "K♣️", "A♣️"}
var diamonds = []string{"2♦️", "3♦️", "4♦️", "5♦️", "6♦️", "7♦️", "8♦️", "9♦️", "10♦️", "J♦️", "Q♦️", "K♦️", "A♦️"}
var hearts = []string{"2♥️", "3♥️", "4♥️", "5♥️", "6♥️", "7♥️", "8♥️", "9♥️", "10♥️", "J♥️", "Q♥️", "K♥️", "A♥️"}
var spades = []string{"2♠️", "3♠️", "4♠️", "5♠️", "6♠️", "7♠️", "8♠️", "9♠️", "10♠️", "J♠️", "Q♠️", "K♠️", "A♠️"}

func NewDeck() Deck {
	cards := slices.Concat(clubs, diamonds, hearts, spades)

	return &deckImpl{
		cards: cards,
	}
}

func (d *deckImpl) Shuffle() {
	cards := d.cards
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
}

func (d *deckImpl) Draw(n int) []string {
	drawnCards := make([]string, 0, n)
	for range n {
		first, ok := utils.Pop(&d.cards)
		if !ok {
			break
		}
		drawnCards = append(drawnCards, first)
	}
	return drawnCards
}
