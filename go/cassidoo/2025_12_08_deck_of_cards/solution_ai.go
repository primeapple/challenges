package deckOfCards

import (
	"math/rand"
	"time"
)

// DeckAI represents a deck of playing cards
type DeckAI struct {
	cards    []string
	position int // Current position in the deck for drawing
}

// NewDeckAI creates a new deck of 52 cards in standard order
func NewDeckAI() *DeckAI {
	suits := []string{"♣️", "♦️", "♥️", "♠️"}
	ranks := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

	cards := make([]string, 0, 52)
	for _, suit := range suits {
		for _, rank := range ranks {
			cards = append(cards, rank+suit)
		}
	}

	return &DeckAI{
		cards:    cards,
		position: 0,
	}
}

// Shuffle randomizes the order of cards in the deck and resets the position
func (d *DeckAI) Shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
	d.position = 0
}

// Draw returns n cards from the deck, or all remaining cards if fewer than n are available
func (d *DeckAI) Draw(n int) []string {
	if d.position >= len(d.cards) {
		return []string{}
	}

	remaining := len(d.cards) - d.position
	if n > remaining {
		n = remaining
	}

	drawn := make([]string, n)
	copy(drawn, d.cards[d.position:d.position+n])
	d.position += n

	return drawn
}
