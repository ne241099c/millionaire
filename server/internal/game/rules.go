package game

func GetStrength(c Card, isRev bool) int {
	if c.Suit == Joker {
		return 99
	}

	strength := 0
	switch c.Rank {
	case Ace:
		strength = 13
	case Two:
		strength = 14
	default:
		strength = int(c.Rank) - 2
	}

	if isRev {
		strength *= -1
	}

	return strength
}

func IsStronger(candidate, target Card, isRev bool) bool {
	myStrength := GetStrength(candidate, isRev)
	targetStrength := GetStrength(target, isRev)

	return myStrength > targetStrength
}
