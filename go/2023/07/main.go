package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"time"
)

const relative_input string = "../../../inputs/2023/07/input.txt"
const sampleInput = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

var bTest = false
var inputFileName string

func init() {
	flag.BoolVar(&bTest, "t", false, "Run using the sample input")
	flag.StringVar(&inputFileName, "i", "",
		"Path to the puzzle input. "+
			"Default to using the internal relative path. "+
			"Pass 'stdin' to use it instead")
}

func main() {
	flag.Parse()
	var content []byte = []byte(sampleInput)
	var err error
	if !bTest {
		switch inputFileName {
		case "stdin":
			content, err = io.ReadAll(os.Stdin)
		case "":
			content, err = os.ReadFile(relative_input)
		default:
			content, err = os.ReadFile(inputFileName)
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to read file: %v", err)
			os.Exit(66)
		}
	}

	start := time.Now()
	part1, part2 := Solver(content)
	duration := time.Since(start)

	fmt.Printf("Time: %v\nPart 1: %d\nPart 2: %d\n", duration, part1, part2)
}

func Solver(input []byte) (int, int) {
	content := bytes.Split(input, []byte("\n"))
	var hands []Hand
	var part2Hands []Hand
	for _, line := range content {
		if len(line) == 0 {
			continue
		}
		hands = append(hands, NewHand(line, false))
		part2Hands = append(part2Hands, NewHand(line, true))
	}
	slices.SortFunc(hands, func(a, b Hand) int {
		return CompareHands(a, b, STRENGTH)
	})
	slices.SortFunc(part2Hands, func(a, b Hand) int {
		return CompareHands(a, b, JOKER_STRENGTH)
	})

	var part1Total int = SolveHand(hands)
	var part2Total int = SolveHand(part2Hands)

	return part1Total, part2Total
}

func SolveHand(hands []Hand) int {
	var total int = 0
	for i, hand := range hands {
		rank := i + 1
		score := rank * hand.Bid
		total += score
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
	cards := input[:5]
	bid, _ := strconv.Atoi(string(input[6:]))
	return Hand{
		Cards: cards,
		Bid:   bid,
		Type:  GetType(cards, joker),
	}
}

func GetType(hand []byte, jokerMode bool) int {
	var cards map[byte]int = make(map[byte]int)
	for _, card := range hand {
		cards[card] += 1
	}

	// If there is only one card found then they all match
	if len(cards) == 1 {
		return FIVE_OF_A_KIND
	}

	var jokers int
	if jokerMode {
		jokers = cards['J']
		delete(cards, 'J')
	}

	currentType := HIGH_CARD
	for _, count := range cards {
		if count == 5 {
			return FIVE_OF_A_KIND
		}
		if count == 4 {
			currentType = FOUR_OF_A_KIND
		}
		if count == 3 {
			if currentType == ONE_PAIR {
				currentType = FULL_HOUSE
				continue
			}
			currentType = THREE_OF_A_KIND
		}
		if count == 2 {
			if currentType == THREE_OF_A_KIND {
				currentType = FULL_HOUSE
				continue
			}
			if currentType == ONE_PAIR {
				currentType = TWO_PAIR
				continue
			}
			currentType = ONE_PAIR
		}
	}
	if jokers > 0 {
		switch currentType {
		case FOUR_OF_A_KIND:
			return FIVE_OF_A_KIND
		case FULL_HOUSE:
			if jokers == 1 {
				return FOUR_OF_A_KIND
			}
		case THREE_OF_A_KIND:
			if jokers == 2 {
				return FIVE_OF_A_KIND
			}
			if jokers == 1 {
				return FOUR_OF_A_KIND
			}
		case TWO_PAIR:
			return FULL_HOUSE

		case ONE_PAIR:
			if jokers == 3 {
				return FIVE_OF_A_KIND
			}
			if jokers == 2 {
				return FOUR_OF_A_KIND
			}
			if jokers == 1 {
				return THREE_OF_A_KIND
			}
		case HIGH_CARD:
			if jokers == 4 {
				return FIVE_OF_A_KIND
			}
			if jokers == 3 {
				return FOUR_OF_A_KIND
			}
			if jokers == 2 {
				return THREE_OF_A_KIND
			}
			return ONE_PAIR
		}
	}
	return currentType
}

func CompareHands(hand1, hand2 Hand, strength map[byte]int) int {
	if hand1.Type > hand2.Type {
		return 1
	} else if hand1.Type < hand2.Type {
		return -1
	}
	for i := range hand1.Cards {
		if strength[hand1.Cards[i]] > strength[hand2.Cards[i]] {
			return 1
		} else if strength[hand1.Cards[i]] < strength[hand2.Cards[i]] {
			return -1
		}
	}
	return 0
}
