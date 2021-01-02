package day25

import (
	"fmt"
	"math/big"
	"strconv"
)

type pubKeys struct {
	card, door int
}

func parseInput(input []string) pubKeys {
	card, _ := strconv.Atoi(input[0])
	door, _ := strconv.Atoi(input[1])
	ret := pubKeys{
		card: card,
		door: door,
	}

	return ret
}

func bruteForce(start int, goal int) int {
	loop := 0
	for start != goal {
		start *= 7
		start = start % 20201227
		loop++
	}
	return loop
}

func encryptKey(start int, loop int) *big.Int {
	ret := big.NewInt(int64(start))
	ret = ret.Exp(ret, big.NewInt(int64(loop)), nil)
	ret = ret.Mod(ret, big.NewInt(int64(20201227)))
	return ret
}

func problemOne(k pubKeys) string {
	// cardLoop := bruteForce(1, k.card)
	doorLoop := bruteForce(1, k.door)
	encKey := encryptKey(k.card, doorLoop)
	return fmt.Sprintf("%d", encKey)
}

func problemTwo(k pubKeys) string {
	ret := "To be implemented..."
	return ret
}

func Run(input []string) {
	inputParsed := parseInput(input)
	fmt.Println("Answer 1: " + problemOne(inputParsed))
	fmt.Println("Answer 2: " + problemTwo(inputParsed))
}
