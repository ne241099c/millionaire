package game

import "fmt"

type Suit int

const (
	Spade Suit = iota
	Heart
	Diamond
	Club
	Joker
)

type Rank int

const (
	_ Rank = iota
	Ace
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

type Card struct {
	Suit Suit
	Rank Rank
}

func NewCard(suit Suit, rank Rank) Card {
	return Card{
		Suit: suit,
		Rank: rank,
	}
}

func (c Card) String() string {
	if c.Suit == Joker {
		return "Joker"
	}

	suits := []string{"♠", "♥", "♦", "♣"}
	ranks := []string{"", "A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	if int(c.Rank) >= len(ranks) {
		return "?"
	}

	return fmt.Sprintf("%s%s", suits[c.Suit], ranks[c.Rank])
}
