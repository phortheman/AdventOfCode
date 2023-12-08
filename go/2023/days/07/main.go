package main

import (
	file "aoc23/internal"
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"
)

var EXAMPLE string = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

func main() {
	var content [][]byte
	if len(os.Args) != 2 {
		content = file.Read_String_Into_Byte_Slice(EXAMPLE)
	} else {
		var err error
		content, err = file.Read_File_Into_Memory(os.Args[1])
		if err != nil {
			fmt.Println("Error reading file into memory: ", err)
			os.Exit(1)
		}
	}

	var hands []Hand
	for _, line := range content {
		hands = append(hands, NewHand(line, false))
	}

	slices.SortFunc(hands, func(a, b Hand) int {
		return CompareHands(a, b)
	})

	var part1Total int = Solver(hands)

	STRENGTH = JOKER_STRENGTH
	var part2hand []Hand
	for _, line := range content {
		part2hand = append(part2hand, NewHand(line, true))
	}
	slices.SortFunc(part2hand, func(a, b Hand) int {
		return CompareHands(a, b)
	})

	var part2Total int = Solver(part2hand) // 244696398 too small

	fmt.Println(part1Total)
	fmt.Println(part2Total)
}

func Solver(hands []Hand) int {
	var total int = 0
	for i, hand := range hands {
		rank := i + 1
		score := rank * hand.Bid
		total += score
		fmt.Printf("Hand: '%v' has a rank of %v with a bet of %v. Which means this hand scored %v\n", string(hand.Cards), rank, hand.Bid, score)
	}
	return total
}

func CountByte(i []byte, c byte) int {
	count := 0
	for _, v := range i {
		if v == c {
			count++
		}
	}
	return count
}

func HasByte(i []byte, c byte) bool {
	for _, v := range i {
		if v == c {
			return true
		}
	}
	return false
}

var STRENGTH map[byte]int = map[byte]int{
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
	'J': 10,
	'Q': 11,
	'K': 12,
	'A': 13,
}

var JOKER_STRENGTH map[byte]int = map[byte]int{
	'J': 0,
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}

const FIVE_OF_A_KIND int = 7  // AAAAA
const FOUR_OF_A_KIND int = 6  // AA8AA
const FULL_HOUSE int = 5      // 23332
const THREE_OF_A_KIND int = 4 // TTT98
const TWO_PAIR int = 3        // 23432
const ONE_PAIR int = 2        // A23A4
const HIGH_CARD int = 1       // 23456

type Hand struct {
	Cards []byte
	Bid   int
	Type  int
}

func NewHand(input []byte, joker bool) Hand {
	cards, bidRaw, _ := bytes.Cut(input, []byte(" "))
	bid, _ := strconv.Atoi(string(bidRaw))
	return Hand{
		Cards: cards,
		Bid:   bid,
		Type:  GetType(cards, joker),
	}
}

func GetType(hand []byte, jokerMode bool) int {
	var cardCache map[byte]int = make(map[byte]int)
	for _, card := range hand {
		cardCache[card] += 1
	}
	// If there is only one card found then they all match
	if len(cardCache) == 1 {
		return FIVE_OF_A_KIND
	} else if len(cardCache) == 5 {
		// If there are 5 cards see if we have jokers and are in jokerMode
		if jokerMode && cardCache['J'] == 0 {
			return HIGH_CARD
		} else if jokerMode {
			return ONE_PAIR
		} else {
			// All cards are unique
			return HIGH_CARD
		}
	}
	var jokers int
	if jokerMode {
		jokers = cardCache['J']
		delete(cardCache, 'J')
	}
	var cacheCount int
	for _, count := range cardCache {
		if count+jokers == 5 {
			return FIVE_OF_A_KIND
		}
		if count+jokers == 4 {
			return FOUR_OF_A_KIND
		}
		if count+jokers == 3 {
			if cacheCount == 2 {
				return FULL_HOUSE
			}
			if cacheCount == 3 {
				return FULL_HOUSE
			}
			cacheCount = count + jokers
			// Possible Full house otherwise three of a kind
		}
		if count+jokers == 2 {
			if cacheCount == 3 {
				return FULL_HOUSE
			}
			if cacheCount == 2 {
				return TWO_PAIR
			}
			cacheCount = count + jokers
			// Possible Full house or two pairs. othersie pair
		}
	}
	if cacheCount == 2 {
		return ONE_PAIR
	}
	if cacheCount == 3 {
		return THREE_OF_A_KIND
	}
	return HIGH_CARD
}

func CompareHands(hand1, hand2 Hand) int {
	if hand1.Type > hand2.Type {
		return 1
	} else if hand1.Type < hand2.Type {
		return -1
	}
	for i := range hand1.Cards {
		if STRENGTH[hand1.Cards[i]] > STRENGTH[hand2.Cards[i]] {
			return 1
		} else if STRENGTH[hand1.Cards[i]] < STRENGTH[hand2.Cards[i]] {
			return -1
		}
	}
	fmt.Println("Got an unexpected result. Both hand 1 and 2 are equal")
	return 0
}
