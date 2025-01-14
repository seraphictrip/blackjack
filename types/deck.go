package types

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

// Errors Raised by deck
var (
	// Error raised when trying to Draw from an empty deck
	ErrEmptyDeck = errors.New("empty deck")
)

type DeckInterface interface {
	// Shuffle deck
	Shuffle()
	// Draw a  card from the deck
	// will raise ErrEmptyDeck if called on an empty deck
	Draw() (Card, error)
	// Return Cards in deck
	GetCards() []Card
}

// Struct representing a deck of cards
// NOTE: a deck can be made up of multi 52 card decks
type Deck struct {
	cards []Card
}

// String representation fo deck
// Used to ensure Suite shown while printing
func (d Deck) String() string {
	return fmt.Sprint(d.cards)
}

// Constructor to geenrate a standard 52 Card deck
func NewDeck() *Deck {
	deck := make([]Card, 0, 52)
	for _, suite := range Suits {
		for _, rank := range Ranks {
			deck = append(deck, NewCard(suite, rank))
		}
	}
	return &Deck{
		cards: deck,
	}
}

// Create a Deck made up of n standard decks
func NewMultiDeck(n int) *Deck {
	deck := make([]Card, 0, 52*n)
	for i := 0; i < n; i++ {
		for _, suite := range Suits {
			for _, rank := range Ranks {
				deck = append(deck, Card{suite, rank})
			}
		}
	}
	return &Deck{
		cards: deck,
	}
}

// Shuffle deck
func (d Deck) Shuffle() {
	for i := 0; i < len(d.cards); i++ {
		r := rand.IntN(len(d.cards))
		swap(d.cards, i, r)
	}
}

// Draw a  card from the deck
// will raise ErrEmptyDeck if called on an empty deck
func (d *Deck) Draw() (card Card, err error) {
	if len(d.cards) == 0 {
		return card, ErrEmptyDeck
	}
	topIndex := len(d.cards) - 1
	card = d.cards[topIndex]
	d.cards = d.cards[:topIndex]
	return card, nil
}

// Return Cards in deck
func (d Deck) GetCards() []Card {
	return d.cards
}

// helper function to swap cards
func swap[T any](arr []T, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
