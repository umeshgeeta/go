// Finding best poker hand: https://www.888poker.com/magazine/how-to-play-poker/hands

// MIT License
// Author: Umesh Patil, Neosemantix, Inc.
package main

import "sort"

const (
	Spade   = 0
	Diamond = 1
	Club    = 2
	Heart   = 3

	HandType_RoyalFlus     = 11
	HandType_StraightFlush = 12
	HandType_FourOfAKind   = 13
	HandType_FullHouse     = 14
	HandType_Flush         = 15
	HandType_Straight      = 16
	HandType_ThreeOfAKind  = 17
	HandType_TwoPair       = 18
	HandType_Par           = 19
	HandType_HighCard      = 20
)

type Card struct {
	Symbol int
	Value  int // J: 11		Q: 12 		K: 13 		A: 14
}

type Hand []Card

func (h Hand) Len() int {
	return len(h)
}

func (h Hand) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Hand) Less(i, j int) bool {
	if h[i].Value < h[j].Value {
		return true
	}
	return false
}

// return value is HandType int and the high card when the type is HandType_HighCard; else -1
func findBestHand(hand []Card) (int, int) {
	spadeCards := make([]Card, 0)
	headCards := make([]Card, 0)
	diamondCards := make([]Card, 0)
	cludCards := make([]Card, 0)
	values := make(map[int]int)

	for _, c := range hand {
		s := c.Symbol
		ct, present := values[c.Value]
		if present {
			values[c.Value] = ct + 1
		} else {
			values[c.Value] = 1
		}
		switch s {
		case Spade:
			spadeCards = append(spadeCards, c)
		case Heart:
			headCards = append(headCards, c)
		case Diamond:
			diamondCards = append(diamondCards, c)
		case Club:
			cludCards = append(cludCards, c)
		}
	}
	sort.Sort(Hand(spadeCards))
	sort.Sort(Hand(headCards))
	sort.Sort(Hand(diamondCards))
	sort.Sort(Hand(cludCards))

	// let us first check if it is RoyalFlush or not
	flush := determineFlushType(headCards, cludCards, diamondCards, spadeCards)
	if flush != -1 {
		return flush, -1
	}

	// check for "four of a kind"
	// if there are 2 entries we know either 4 of a kind or full house
	if len(values) == 2 {
		for _, vc := range values {
			if vc == 4 {
				return HandType_FourOfAKind, -1
			}
			if vc == 3 {
				return HandType_FullHouse, -1
			}
		}
	}

	if len(values) == 3 {
		firstPairFound := false
		for _, vc := range values {
			if vc == 3 {
				return HandType_ThreeOfAKind, -1
			}
			if vc == 2 {
				if firstPairFound {
					// this is second pair
					return HandType_TwoPair, -1
				}
			}
		}
		if firstPairFound {
			return HandType_Par, -1
		}
	}

	// must be 5 different values
	// we just have to check if those ordered
	sort.Sort(Hand(hand))
	startVal := hand[0].Value
	i := 1
	for i < 5 {
		if hand[1].Value == startVal+1 {
			continue
		}
		i++
	}
	if i == 5 {
		return HandType_Straight, -1
	}
	// it is high card situation
	return HandType_HighCard, hand[4].Value
}

func determineFlushType(hearts []Card, clubs []Card, diamond []Card, spade []Card) int {
	if len(hearts) == 5 {
		return isFlush(hearts)
	}
	if len(clubs) == 5 {
		return isFlush(hearts)
	}
	if len(diamond) == 5 {
		return isFlush(hearts)
	}
	if len(spade) == 5 {
		return isFlush(hearts)
	}
	return -1
}

// we asssume that incoming cards sorted and they are of the same suite
func isFlush(cards []Card) int {
	result := HandType_Flush
	if cards[0].Value == 10 && cards[1].Value == 11 && cards[2].Value == 12 && cards[3].Value == 13 && cards[4].Value == 14 {
		result = HandType_RoyalFlus
	} else {
		// if those sorted, then we get straight flush
		diff := cards[0].Value - cards[4].Value
		if diff == 4 {
			result = HandType_StraightFlush
		}
	}
	return result
}
