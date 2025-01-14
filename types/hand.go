package types

import "fmt"

type HandInterface interface {
	// Get a reference to slice of cards in hand
	GetCards() []Card
	// Get score of currently held hand
	// the score will be recalculated each time a card is added
	// via AddCard
	GetScore() int
	// Add a new card to hand, and return current score
	AddCard(card Card) int
}

// Hand represents a blackjack hand
type Hand struct {
	score int
	cards []Card
}

// Printing a hand should show cards
func (h Hand) String() string {
	return fmt.Sprint(h.cards)
}

// Constructor for a new Hand given an intial set of cards
// and Empty hand can be genrated with NewHand(nil)
func NewHand(inital []Card) *Hand {
	hand := &Hand{}
	for _, card := range inital {
		hand.AddCard(card)
	}
	return hand
}

// Getter for score
// NOTE: Game round is responsible for interpreting score
func (h Hand) GetScore() int {
	return h.score
}

// Getter for Cards
func (h Hand) GetCards() []Card {
	return h.cards
}

// Add a card to hand and trigger recalculating the score
func (h *Hand) AddCard(card Card) int {
	h.cards = append(h.cards, card)
	h.calculateScore()
	return h.GetScore()
}

// Internal method to recalculate score
// score is recalculated after each card is added
// Aces are applied last
func (h *Hand) calculateScore() {
	score := 0
	// iterate over cards and calc score
	aces := 0
	for _, card := range h.cards {
		switch card.Rank {
		case Ace:
			aces++
		case Jack, Queen, King, Ten:
			score += 10
		default:
			score += int(card.Rank)
		}
	}
	// apply aces
	// if all of our aces together max us out
	// we can short-circuit
	if (score + aces) >= 21 {
		// apply them all as ones and move on
		h.score = score + aces
		return
	}
	for i := 0; i < aces; i++ {
		if score+11 <= 21 && score+11+(aces-1) <= 21 {
			score += 11
		} else {
			score += 1
		}
	}

	h.score = score
}
