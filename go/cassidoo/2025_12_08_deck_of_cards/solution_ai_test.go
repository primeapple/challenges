package deckOfCards

import (
	"cassidoo/utils"
	"testing"
)

func TestDeckOfCardsAI(t *testing.T) {
	t.Run("should give expected cards without shuffling", func(t *testing.T) {
		deck := NewDeckAI()
		got := deck.Draw(3)
		want := []string{"2♣️", "3♣️", "4♣️"}
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should give max all cards", func(t *testing.T) {
		deck := NewDeckAI()
		got := len(deck.Draw(1000))
		want := 52
		utils.AssertEquals(t, got, want)
	})

	t.Run("should shuffle cards", func(t *testing.T) {
		deck := NewDeckAI()
		unshuffledCards := deck.Draw(52)

		deck = NewDeckAI()
		deck.Shuffle()
		shuffledCards := deck.Draw(52)
		utils.AssertNotDeepEquals(t, shuffledCards, unshuffledCards)
	})
}
