package day22

import (
	"fmt"
	"strconv"
	"strings"
)

type Problem struct {
	input []string
}

func (p *Problem) SetInput(input []string) {
	p.input = input
}

func (p *Problem) Run() {
	p1, p2 := parseInput(p.input)
	p11, p12 := p1.copyDeck(len(*p1)), p2.copyDeck(len(*p2))
	fmt.Println("Answer 1: " + problemOne(p1, p2))
	fmt.Println("Answer 2: " + problemTwo(p11, p12))
}

type player []int

type gameMap map[string]map[string]int

type helper interface {
	draw() int
	addCards(int, int)
	copyDeck(int) *player
	toStr() string

	exists(*player, *player) bool
	add(*player, *player, int) bool
}

func (gm *gameMap) exists(p1, p2 *player) bool {
	p1Str := p1.toStr()
	p2Str := p2.toStr()
	_, ok := (*gm)[p1Str][p2Str]
	if !ok {
		return false
	}
	return true
}

func (gm *gameMap) add(p1, p2 *player, v int) {
	p1Str := p1.toStr()
	p2Str := p2.toStr()
	_, ok := (*gm)[p1Str]
	if !ok {
		(*gm)[p1Str] = map[string]int{}
	}
	_, ok = (*gm)[p1Str][p2Str]
	if !ok {
		(*gm)[p1Str][p2Str] = v
	}
}

func (p *player) draw() int {
	ret := (*p)[0]
	*p = (*p)[1:]
	return ret
}

func (p *player) addCards(first, second int) {
	*p = append(*p, []int{first, second}...)
}

func (p *player) copyDeck(next int) *player {
	ret := &player{}
	for i := 0; i < next; i++ {
		*ret = append(*ret, (*p)[i])
	}
	return ret
}

func (p *player) toStr() string {
	strSlc := []string{}
	for _, n := range *p {
		strSlc = append(strSlc, fmt.Sprintf("%d", n))
	}

	return strings.Join(strSlc, ", ")
}

func parseInput(input []string) (*player, *player) {
	p1, p2 := &player{}, &player{}
	count := 0
	for _, line := range input {
		if strings.HasPrefix(line, "Player") {
			count++
			continue
		}

		if line == "" {
			continue
		}

		n, _ := strconv.Atoi(line)
		switch count {
		case 1:
			*p1 = append(*p1, n)
		case 2:
			*p2 = append(*p2, n)
		default:
			panic("WRONG")
		}
	}

	return p1, p2
}

func problemOne(p1, p2 *player) string {
	winner := &player{}
	for {
		if len(*p1) == 0 {
			winner = p2
			break
		}

		if len(*p2) == 0 {
			winner = p1
			break
		}

		cp1 := p1.draw()
		cp2 := p2.draw()

		switch {
		case cp1 > cp2:
			p1.addCards(cp1, cp2)
		case cp2 > cp1:
			p2.addCards(cp2, cp1)
		default:
			panic("something went wrong")
		}
	}

	ret := 0
	for i := len(*winner) - 1; i >= 0; i-- {
		ret += (*winner)[i] * (len(*winner) - i)
	}

	return fmt.Sprintf("%d", ret)
}

func recCombat(p1, p2 *player, gm *gameMap, ngame int) *player {
	// fmt.Printf("=== Game %d ===\n\n", ngame)
	nround := 0
	var winner *player
	for len(*p1) != 0 && len(*p2) != 0 {
		nround++
		// fmt.Printf("-- Round %d (Game %d) --\n", nround, ngame)
		visited := gm.exists(p1, p2)
		if visited {
			return p1
		}
		gm.add(p1, p2, -1)

		// fmt.Printf("Player 1's deck: %s\n", p1.toStr())
		// fmt.Printf("Player 2's deck: %s\n", p2.toStr())

		p1d, p2d := p1.draw(), p2.draw()
		// fmt.Printf("Player 1 plays: %d\n", p1d)
		// fmt.Printf("Player 2 plays: %d\n", p2d)

		// recCom
		if p1d <= len(*p1) && p2d <= len(*p2) {
			// fmt.Println("Playing a sub-game to determine the winner...\n")
			p1c, p2c := p1.copyDeck(p1d), p2.copyDeck(p2d)
			winner = recCombat(p1c, p2c, &gameMap{}, ngame+1)
			// fmt.Printf("...anyway, back to game %d.\n", ngame)

			if winner == p1c {
				// fmt.Printf("Player 1 wins round %d of game %d!\n\n", nround, ngame)
				p1.addCards(p1d, p2d)
			} else {
				// fmt.Printf("Player 2 wins round %d of game %d!\n\n", nround, ngame)
				p2.addCards(p2d, p1d)
			}

			// fmt.Printf("=== Game %d ===\n\n", ngame)
			continue
		}

		// no cards to recurse, higher cards wins
		if len(*p1) == 0 || len(*p2) == 0 {
			if p1d > p2d {
				// fmt.Printf("Player 1 wins round %d of game %d!\n\n", nround, ngame)
				p1.addCards(p1d, p2d)
				winner = p1
				continue
			}
			// fmt.Printf("Player 2 wins round %d of game %d!\n\n", nround, ngame)
			winner = p2
			p2.addCards(p2d, p1d)
			continue
		}

		// normal game
		if p1d > p2d {
			// fmt.Printf("Player 1 wins round %d of game %d!\n\n", nround, ngame)
			p1.addCards(p1d, p2d)
			winner = p1
		} else {
			// fmt.Printf("Player 2 wins round %d of game %d!\n\n", nround, ngame)
			p2.addCards(p2d, p1d)
			winner = p2
		}
	}

	return winner
}

func problemTwo(p1, p2 *player) string {
	winner := recCombat(p1, p2, &gameMap{}, 1)
	// fmt.Println("== Post-game results ==")
	// fmt.Printf("Player 1's deck: %s\n", p1.toStr())
	// fmt.Printf("Player 2's deck: %s\n", p2.toStr())

	ret := 0
	for i := len(*winner) - 1; i >= 0; i-- {
		ret += (*winner)[i] * (len(*winner) - i)
	}

	return fmt.Sprintf("%d", ret)
}
