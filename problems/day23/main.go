package day23

import (
	"fmt"
	"strconv"
	"strings"
)

type helpers interface {
	debug(int)
	removeCups(int)
	initMemo() map[int]*cup
	exists(int) bool
	findDst() *cup
	insert(*cup)
}

type cup struct {
	label int
	next  *cup
}

func (c *cup) debug(n int) string {
	var bldr strings.Builder
	curr := c.next
	bldr.WriteString(fmt.Sprintf("%d ", c.label))
	if n == -1 {
		for curr != c && curr != nil {
			bldr.WriteString(fmt.Sprintf("%d ", curr.label))
			curr = curr.next
		}
		return bldr.String()
	}

	count := 0
	for count < n && curr != nil {
		count++
		bldr.WriteString(fmt.Sprintf("%d ", curr.label))
		curr = curr.next
	}
	return bldr.String()
}

func parseInput(input []string) *cup {
	line := input[0]
	var first, prev *cup
	for i, c := range line {
		n, _ := strconv.Atoi(string(c))
		if i == 0 {
			first = &cup{
				label: n,
				next:  nil,
			}
			prev = first
			continue
		}

		curr := &cup{
			label: n,
			next:  nil,
		}

		prev.next = curr
		prev = curr
	}
	prev.next = first

	return first
}

func (c *cup) removeCups() *cup {
	count := 0
	begin := c.next
	end := begin
	for count < 2 {
		count++
		end = end.next
	}
	c.next = end.next
	end.next = nil
	return begin
}

func (c *cup) initMemo() map[int]*cup {
	ret := map[int]*cup{}
	first := c
	curr := c.next
	ret[first.label] = first
	for curr != first {
		ret[curr.label] = curr
		curr = curr.next
	}
	return ret
}

func (c *cup) exists(n int) bool {
	curr := c
	if n < 1 {
		return true
	}
	for curr != nil {
		if curr.label == n {
			return true
		}
		curr = curr.next
	}
	return false
}

func (c *cup) insert(pu *cup) {
	tmpNext := c.next
	c.next = pu

	curr := pu
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = tmpNext
}

func problemOne(c *cup, n int) string {
	memo := c.initMemo() // map[int]*cup

	curr := c
	for i := 0; i < n; i++ {
		// 1 - remove 3 cards
		pu := curr.removeCups()

		// 2 - dst
		var dst *cup
		dstN := curr.label - 1
		for pu.exists(dstN) {
			dstN -= 1
			if dstN < 1 {
				dstN = 9
			}
		}
		dst = memo[dstN]

		// 3 - insert pu after dst
		dst.insert(pu)

		// 4 - next current cup
		curr = curr.next
	}

	one := strings.Split(memo[1].debug(-1), " ")[1:]
	ret := strings.Join(one, "")
	return ret
}

func problemTwo(c *cup, n int) string {
	// modify input
	end := c.next
	for end.next != c {
		end = end.next
	}
	tmpFirst := end.next
	curr := end
	for i := 10; i <= 1000000; i++ {
		curr.next = &cup{
			label: i,
			next:  nil,
		}
		curr = curr.next
	}
	curr.next = tmpFirst
	memo := c.initMemo() // map[int]*cup

	curr = c
	for i := 0; i < n; i++ {
		// 1 - remove 3 cards
		pu := curr.removeCups()

		// 2 - dst
		var dst *cup
		dstN := curr.label - 1
		for pu.exists(dstN) {
			dstN -= 1
			if dstN < 1 {
				dstN = 1000000
			}
		}
		dst = memo[dstN]

		// 3 - insert pu after dst
		dst.insert(pu)

		// 4 - next current cup
		curr = curr.next
	}

	return fmt.Sprintf("%d", memo[1].next.label*memo[1].next.next.label)
}

func Run(input []string) {
	inputParsed1 := parseInput(input)
	inputParsed2 := parseInput(input)
	fmt.Println("Answer 1: " + problemOne(inputParsed1, 100))
	fmt.Println("Answer 2: " + problemTwo(inputParsed2, 10000000))
}
