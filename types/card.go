package types

import (
	"fmt"
)

type Suit rune

// Suit represent the four Suits in a standard deck
const (
	Spades   Suit = '\u2660'
	Hearts   Suit = '\u2661'
	Diamonds Suit = '\u2662'
	Clubs    Suit = '\u2663'
)

// Collection of Suits for easy iteration
// such as that done while genrating a deck
var Suits = []Suit{
	Spades,
	Hearts,
	Diamonds,
	Clubs,
}

// Rank of a card, this represent the Rank value of a card
type Rank byte

const (
	Ace Rank = 1 + iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

// Collection Rank for easy iteration
// such as that done while genrating a deck
var Ranks = []Rank{
	Ace,
	Two,
	Three,
	Four,
	Five,
	Six,
	Seven,
	Eight,
	Nine,
	Ten,
	Jack,
	Queen,
	King,
}

// Card struct
type Card struct {
	Suit Suit
	Rank Rank
}

// Simple Card constructor
func NewCard(Suit Suit, Rank Rank) Card {
	return Card{Suit, Rank}
}

// Pretty print a card
func (c Card) String() string {
	Rank := fmt.Sprint(c.Rank)
	switch c.Rank {
	case Ace:
		Rank = "ACE"
	case Jack:
		Rank = "JACK"
	case Queen:
		Rank = "QUEEN"
	case King:
		Rank = "KING"
	}
	return fmt.Sprintf("{%v, %v}", string(c.Suit), Rank)
}
