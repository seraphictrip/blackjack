package types

import (
	"fmt"
)

type Suite rune

// Suite represent the four suites in a standard deck
const (
	Spades   Suite = '\u2660'
	Hearts   Suite = '\u2661'
	Diamonds Suite = '\u2662'
	Clubs    Suite = '\u2663'
)

// Array version of suits for easy iteration
// such as that done while genrating a deck
var Suites = []Suite{
	Spades,
	Hearts,
	Diamonds,
	Clubs,
}

// Face of a card, this represent the face value of a card
// NOTE: represents all cards, not just "face" cards, i.e. Two, Three... included
type Face byte

const (
	Ace Face = 1 + iota
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

// Array version of face for easy iteration
// such as that done while genrating a deck
var Faces = []Face{
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
	suite Suite
	face  Face
}

// Simple Card constructor
func NewCard(suite Suite, face Face) Card {
	return Card{suite, face}
}

// Pretty print a card
func (c Card) String() string {
	face := fmt.Sprint(c.face)
	switch c.face {
	case Ace:
		face = "ACE"
	case Jack:
		face = "JACK"
	case Queen:
		face = "QUEEN"
	case King:
		face = "KING"
	}
	return fmt.Sprintf("{%v, %v}", string(c.suite), face)
}
