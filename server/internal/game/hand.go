package game

import "sort"

func IsPair(cards []Card) bool {
	if len(cards) < 2 {
		return false
	}

	baseRank := -1

	for _, c := range cards {
		if c.Suit == Joker {
			continue
		}

		if baseRank == -1 {
			baseRank = int(c.Rank)
		} else {
			if baseRank != int(c.Rank) {
				return false
			}
		}
	}
	return true
}

func IsSequence(cards []Card) bool {
	if len(cards) < 3 {
		return false
	}

	var normalCards []Card
	var baseSuit Suit = -1
	jokerCount := 0

	for _, c := range cards {
		if c.Suit == Joker {
			jokerCount++
			continue
		}

		if baseSuit == -1 {
			baseSuit = c.Suit
		} else if baseSuit != c.Suit {
			return false
		}

		normalCards = append(normalCards, c)
	}

	sort.Slice(normalCards, func(i, j int) bool {
		return toSeqRank(normalCards[i]) < toSeqRank(normalCards[j])
	})

	for i := 0; i < len(normalCards)-1; i++ {
		current := toSeqRank(normalCards[i])
		next := toSeqRank(normalCards[i+1])

		diff := next - current

		if diff == 0 {
			return false
		}

		gap := diff - 1

		if gap > 0 {
			jokerCount -= gap // ジョーカーを消費
			if jokerCount < 0 {
				return false // ジョーカーが足りない->階段不成立
			}
		}
	}
	return true
}

func toSeqRank(c Card) int {
	if c.Rank == Ace {
		return 14
	}
	if c.Rank == Two {
		return 15
	}
	return int(c.Rank)
}
