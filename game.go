package main

func playCards(player1 player, player2 player) player {
	var winner player
	if summeryCardsInHand(player1.hand) > summeryCardsInHand(player2.hand) {
		player1.rank = 1
		player2.rank = 2
		winner = player1
	} else if summeryCardsInHand(player1.hand) < summeryCardsInHand(player2.hand) {
		player1.rank = 2
		player2.rank = 1
		winner = player2
	}
	return winner
}

func summeryCardsInHand(hand deck) int {
	var summery int
	for _, card := range hand.Cards {
		summery += card.Value
	}

	return summery
}
