package main

import (
	"blackjack/types"
	"fmt"
	"os"
)

func main() {
	// ask player for name and instantiate player
	playerName := getPlayerName()
	player := types.NewPlayer(playerName, 1000)
	// instantiate dealer
	dealer := types.NewDealer()
	// instantiate deck
	deck := types.NewDeck()

	// play as long as player has money
	// the house always wins
	for player.GetBalance() > 0 {
		fmt.Printf("%v starts round with $%d\n", player.GetName(), player.GetBalance())
		types.NewGameRound(dealer, player, deck).Play()
	}
}

func getPlayerName() string {
	var playerName string
	// instantiate componenets
	if len(os.Args) != 2 {
		fmt.Print("What is your name player? ")
		fmt.Scanln(&playerName)
	} else {
		playerName = os.Args[1]
	}
	return playerName
}
