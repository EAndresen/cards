package main

type player struct {
	firstName string
	lastName  string
	rank      int
	hand      deck
	deck      deck
}

func createPlayer(firstName string, lastName string) player {
	return player{
		firstName: firstName,
		lastName:  lastName,
	}
}
