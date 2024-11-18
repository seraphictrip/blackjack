package types

import (
	"errors"
	"fmt"
)

// GameRoundInterface
type GameRoundInterface interface {
	Play()
	GetDealer() PlayerInterface
	GetPlayer() Player
}

// Concrete implementation of a game round
type GameRound struct {
	dealer DealerInterface
	player PlayerInterface
	deck   *Deck
}

// Constructor to generate a new GameRound
func NewGameRound(dealer DealerInterface, player PlayerInterface, deck *Deck) *GameRound {
	return &GameRound{
		dealer,
		player,
		deck,
	}
}

// Interactive Driver code to play a game round
func (r *GameRound) Play() {
	// shuffle the deck
	r.deck.Shuffle()
	r.player.ClearHand()
	r.dealer.ClearHand()

	if r.player.GetBalance() == 0 {
		fmt.Printf("%q does not have money to bet\n", r.player.GetName())
		return
	}

	bet := r.getPlayerBet()
	// house meets bet
	pot := bet * 2
	playerHand := r.player.GetHand()
	dealerHand := r.dealer.GetHand()

	r.dealInitialHand()
	if playerHand.GetScore() == 21 {
		fmt.Printf("%v (%v): %v", r.player.GetName(), playerHand.GetScore(), playerHand)
		r.player.RecieveWinnings(pot)
		return
	} else if r.dealer.GetHand().GetScore() == 21 {
		fmt.Printf("%v (%v): %v", r.dealer.GetName(), dealerHand.GetScore(), dealerHand)
		return
	}

	r.dealer.SetTarget(playerHand.GetScore())

	for {
		fmt.Printf("%v hand (%v): %v\n", r.player.GetName(), playerHand.GetScore(), playerHand)
		// competitors decide if they will hit
		playerHit := r.player.Hit()
		dealerHit := r.dealer.Hit()
		if !playerHit && !dealerHit {
			// if no one is hitting the round is over, calculate winner
			pScore, dScore := playerHand.GetScore(), dealerHand.GetScore()
			if pScore <= 21 && pScore > dScore {
				//
				r.player.RecieveWinnings(pot)
			} else if pScore == dScore {
				// draw, get money back
				r.player.RecieveWinnings(bet)
			}
			fmt.Printf("%v (%v): %v vs %v (%v): %v\n", r.dealer.GetName(), dealerHand.GetScore(), dealerHand, r.player.GetName(), playerHand.GetScore(), playerHand)
			return
		}
		if playerHit {
			card, _ := r.deck.Draw()
			playerHand.AddCard(card)
			if playerHand.GetScore() > 21 {
				fmt.Printf("%v busts: %v\n", r.player.GetName(), playerHand)
				return
			}
			r.dealer.SetTarget(playerHand.GetScore())
		}
		if dealerHit {
			card, _ := r.deck.Draw()
			dealerHand.AddCard(card)
			if dealerHand.GetScore() > 21 {
				fmt.Printf("%v busts: %v\n", r.dealer.GetName(), dealerHand)
				r.player.RecieveWinnings(pot)
				return
			}
		}
	}

}

// Get player input on bet
func (r GameRound) getPlayerBet() int {

	var bet int
	fmt.Printf("%v place bet: ", r.player.GetName())
	for {
		_, err := fmt.Scanln(&bet)
		if err != nil {
			continue
		}
		err = r.player.Bet(bet)
		if err != nil {
			fmt.Println(err)
			// cheating is not tolerated, take the players money and kick them out
			if errors.Is(err, ErrAttemptedCheat) {
				bet = 0
				r.player.RecieveWinnings(-r.player.GetBalance())
			} else {
				continue
			}
		}
		break
	}
	return bet
}

// Deal an intial hand to Player and Dealer
// player, dealer, player, dealer
func (r GameRound) dealInitialHand() {
	deck := r.deck
	dealersHand := r.dealer.GetHand()
	playersHand := r.player.GetHand()
	for i := 0; i < 2; i++ {
		card, _ := deck.Draw()
		playersHand.AddCard(card)
		card, _ = deck.Draw()
		dealersHand.AddCard(card)

	}
}
