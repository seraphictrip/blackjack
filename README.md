# Blackjack

Build a simple CLI blackjack game where a single player can play against a dealer.

## Classes

### Player
Players come in two flavors.  The dealer, who will play as computer, and (human) Player.  


#### Player
Human player, who can interact with world.

#### Dealer
Computer player with unlimited funds, representing the house.

### Cards

#### Suite Enum and Iteratorx
* Spades
* Hearts
* Diamonds
* Clubs

#### Face Enum and Iterator
* Ace
* 2 
* 3 
* 4 
* 5 
* 6 
* 7
* 8
* 9 
* 10 
* Jack
* Queen
* King

#### Card
A card has a suite and face value.

#### Deck
A deck is a collection of cards.  A standard deck has 52 cards.

#### Hand
A hand is a collection of cards and a score.

