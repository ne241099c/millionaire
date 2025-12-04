package game

import (
	"math/rand"
)

type Deck []Card

func NewDeck(jokerCount int) Deck {
	// make([]型, 長さ, 容量)
	cards := make(Deck, 0, 52+jokerCount)

	for suit := Spade; suit <= Club; suit++ {
		for rank := Ace; rank <= King; rank++ {
			cards = append(cards, NewCard(suit, rank))
		}
	}

	for i := 0; i < jokerCount; i++ {
		cards = append(cards, NewCard(Joker, 0))
	}

	return cards
}

func (d Deck) Shuffle() {
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}

func (d *Deck) Draw(n int) []Card {
	if n > len(*d) {
		n = len(*d)
	}

	cards := (*d)[:n]

	*d = (*d)[n:]

	return cards
}
