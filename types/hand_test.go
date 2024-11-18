package types_test

import (
	"blackjack/types"
	"math/rand/v2"
	"slices"
	"strconv"
	"testing"
)

var HandTests = []struct {
	cards    []types.Card
	expected int
}{
	{},
	// 1 ace
	{[]types.Card{types.NewCard(types.Clubs, types.Ace)}, 11},
	// 21 aces
	{slices.Repeat([]types.Card{types.NewCard(types.Clubs, types.Ace)}, 21), 21},
	// ACE + 10
	{makeCards([]types.Face{1, 10}), 21},
	// ACE + Jack
	{makeCards([]types.Face{1, 11}), 21},
	// ACE + Queen
	{makeCards([]types.Face{1, 12}), 21},
	// ACE + King
	{makeCards([]types.Face{1, 13}), 21},
	// ACE + 2 + 8
	{makeCards([]types.Face{1, 2, 8}), 21},
	{makeCards([]types.Face{1, 3, 7}), 21},
	{makeCards([]types.Face{1, 1, 9}), 21},
	{makeCards([]types.Face{1, 4, 6}), 21},
	{makeCards([]types.Face{1, 5, 5}), 21},
	// would break
	{makeCards([]types.Face{1, 1, 10}), 12},
	{makeCards([]types.Face{1, 2, 10}), 13},
	{makeCards([]types.Face{1, 3, 10}), 14},
	{makeCards([]types.Face{1, 4, 10}), 15},
	{makeCards([]types.Face{1, 5, 10}), 16},
	{makeCards([]types.Face{1, 6, 10}), 17},
	{makeCards([]types.Face{1, 7, 10}), 18},
	{makeCards([]types.Face{1, 8, 10}), 19},
	{makeCards([]types.Face{1, 9, 10}), 20},
	{makeCards([]types.Face{1, 10, 11}), 21},
	{makeCards([]types.Face{1, 10, 12}), 21},
	{makeCards([]types.Face{1, 10, 13}), 21},
	// no saving
	{makeCards([]types.Face{1, 2, 10, 11}), 23},
}

func TestHand(t *testing.T) {
	for i, e := range HandTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			hand := types.NewHand(e.cards)
			actual := hand.GetScore()
			if actual != e.expected {
				t.Fatalf("(%v).GetScore() = %v, want %v", hand, actual, e.expected)
			}
		})
	}
}

func makeCards(fs []types.Face) []types.Card {
	cards := make([]types.Card, len(fs))
	for i := 0; i < len(fs); i++ {
		r := rand.IntN(len(types.Suites))
		cards[i] = types.NewCard(types.Suites[r], fs[i])
	}
	return cards
}
