package deckOfCards

import (
	"cassidoo/utils"
	"testing"
)

func TestDeckOfCards(t *testing.T) {
	t.Run("should give expected cards without shuffling", func(t *testing.T) {
		deck := NewDeck()
		got := deck.Draw(3)
		want := []string{"2♣️", "3♣️", "4♣️"}
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should give max all cards", func(t *testing.T) {
		deck := NewDeck()
		got := len(deck.Draw(1000))
		want := 52
		utils.AssertEquals(t, got, want)
	})

	t.Run("should shuffle cards", func(t *testing.T) {
		deck := NewDeck()
		unshuffledCards := deck.Draw(52)

		deck = NewDeck()
		deck.Shuffle()
		shuffledCards := deck.Draw(52)
		utils.AssertNotDeepEquals(t, shuffledCards, unshuffledCards)
	})

	t.Run("drawing multiple times works", func(t *testing.T) {
		deck := NewDeck()
		deck.Shuffle()
		first := deck.Draw(1)
		second := deck.Draw(1)

		utils.AssertNotDeepEquals(t, first, second)
	})
}
