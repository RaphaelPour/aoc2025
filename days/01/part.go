package main

import (
	"fmt"

	"github.com/RaphaelPour/stellar/input"
	"github.com/RaphaelPour/stellar/strings"
)

func part1(data []string) int {
	dial := 50
	zeroHits := 0

	fmt.Println("dial -> 50")

	for i, line := range data {
		if len(line) < 2 {
			fmt.Printf("bad line %d: %s\n", i, line)
			continue
		}

		val := strings.ToInt(line[1:])
		if line[0] == 'L' {
			val = 100 - val
		}

		dial = (dial + val) % 100

		fmt.Printf("%s: %d\n", line, dial)

		if dial == 0 {
			zeroHits += 1
		}
	}

	return zeroHits
}

func part2(data []string) int {
	return 0
}

func main() {
	data := input.LoadString("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println("bad: 31")
	fmt.Println(part1(data))

	// fmt.Println("== [ PART 2 ] ==")
	// fmt.Println(part2(data))
}
