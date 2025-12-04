package game

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
