package types

import (
	"errors"
	"fmt"
	"math"
)

var (
	// A soft error raised when a player tries to bet more than they have
	ErrInsufficentFunds = errors.New("insufficent funds")
	// an error raised when a player attempt to bet negative amounts
	ErrAttemptedCheat = errors.New("cheating")
)

type PlayerInterface interface {
	// apply a bet and adust players balance
	Bet(int) error
	// Getter for Balance
	GetBalance() int
	// Method used by House to payout
	RecieveWinnings(int)
	// Getter for hand
	GetHand() *Hand
	// Clear the players hand
	ClearHand()
	// Get players name
	GetName() string
	// Indicator as to if player should hit or not
	Hit() bool
}

// Human player
type Player struct {
	name    string
	balance int
	hand    *Hand
}

// Constructor for a human player
func NewPlayer(name string, balance int) *Player {
	return &Player{name: name, balance: balance, hand: NewHand(nil)}
}

// Place a bet
func (p *Player) Bet(bet int) error {
	if p.balance-bet < 0 {
		return ErrInsufficentFunds
	}
	if bet < 0 {
		return ErrAttemptedCheat
	}
	p.balance -= bet
	return nil
}

// Getter for balance
func (p *Player) GetBalance() int {
	return p.balance
}

// Method so house can payout
func (p *Player) RecieveWinnings(winnings int) {
	p.balance = p.balance + winnings
}

// Getter for hand
func (p *Player) GetHand() *Hand {
	return p.hand
}

// Clear hand, used at end of each round
func (p *Player) ClearHand() {
	p.hand = NewHand(nil)
}

// Interactive method, used to determine if player should hit
func (p *Player) Hit() bool {
	if p.GetHand().GetScore() >= 21 {
		return false
	}
	hit := "n"
	fmt.Printf("%v: hit? [y/n]", p.name)
	fmt.Scanln(&hit)
	return hit == "y"
}

// Get for Player name
func (p *Player) GetName() string {
	return p.name
}

// Dealer Interface
// Dealers are players with a traget
type DealerInterface interface {
	PlayerInterface
	SetTarget(int)
}

// Dealers are players with an unlimited bank
// and non-interactive Hit functionality
type Dealer struct {
	*Player
	target int
}

func NewDealer() *Dealer {
	p := NewPlayer("Dealer", 0)
	return &Dealer{p, 21}
}

func (d *Dealer) GetName() string {
	return "Dealer"
}

func (d *Dealer) GetBalance() int {
	return math.MaxInt
}

func (d *Dealer) Bet(int) error {
	return nil
}

func (d *Dealer) Hit() bool {
	return d.GetHand().GetScore() < d.target
}

func (d *Dealer) SetTarget(target int) {
	d.target = target
}
