package types_test

import (
	"blackjack/types"
	"errors"
	"slices"
	"strconv"
	"testing"
)

var DeckTests = []struct {
	deck types.DeckInterface
}{
	{deck: types.NewDeck()},
}

func TestDeckShuffle(t *testing.T) {
	for i, e := range DeckTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// deck should
			sampleDeck := types.NewDeck()
			if !slices.Equal(e.deck.GetCards(), sampleDeck.GetCards()) {
				t.Fatalf("Deck = %v, want %v", e.deck, sampleDeck)
			}
			e.deck.Shuffle()
			if slices.Equal(e.deck.GetCards(), sampleDeck.GetCards()) {
				t.Fatalf("Deck = %v, want %v", e.deck, sampleDeck)
			}
		})
	}
}

var NewMultiDeckTests = []struct {
	n int
}{
	{1},
	{2},
	{3},
	{4},
	{5},
	{8},
	{16},
	{32},
}

func TestNewMultiDeck(t *testing.T) {
	for i, e := range NewMultiDeckTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			deck := types.NewMultiDeck(e.n)
			count := len(deck.GetCards())
			expectedCount := e.n * 52
			// deck should be made up of equal number of each cards
			if count != expectedCount {
				t.Fatalf("NewMultiDeck(%v) resulted in %v cards, want %v", e.n, count, expectedCount)
			}

			freq := make(map[types.Card]int)
			for _, card := range deck.GetCards() {
				freq[card]++
			}
			if len(freq) != 52 {
				t.Fatalf("unexpected cards in %v", freq)
			}
			for key, val := range freq {
				if val != e.n {
					t.Fatalf("expected %v of each card, got %v of %v", e.n, val, key)
				}
			}
			// deck should have same make up after shuffle
			freq = make(map[types.Card]int)
			deck.Shuffle()
			for _, card := range deck.GetCards() {
				freq[card]++
			}
			if len(freq) != 52 {
				t.Fatalf("unexpected cards in %v", freq)
			}
			for key, val := range freq {
				if val != e.n {
					t.Fatalf("expected %v of each card, got %v of %v", e.n, val, key)
				}
			}
		})
	}
}

var DrawTests = []struct {
	deck types.DeckInterface
}{
	{types.NewDeck()},
	{types.NewMultiDeck(1)},
	{types.NewMultiDeck(2)},
	{types.NewMultiDeck(4)},
	{types.NewMultiDeck(8)},
	{types.NewMultiDeck(16)},
	{types.NewMultiDeck(32)},
}

func TestDraw(t *testing.T) {
	for i, e := range DrawTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e.deck.Shuffle()
			n := len(e.deck.GetCards())
			for i := 1; i <= n; i++ {
				e.deck.Draw()
				if len(e.deck.GetCards()) != n-i {
					t.Fatalf("deck.Draw() does not remove card from deck")
				}
			}
			card, err := e.deck.Draw()
			if !errors.Is(err, types.ErrEmptyDeck) {
				t.Fatalf("deck.Draw() = %v, %q, want %v, %q", card, err, types.Card{}, types.ErrEmptyDeck)
			}
		})
	}
}
