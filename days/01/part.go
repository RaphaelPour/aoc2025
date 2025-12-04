package main

import (
	"fmt"

	"github.com/RaphaelPour/stellar/input"
	"github.com/RaphaelPour/stellar/strings"
	"github.com/RaphaelPour/stellar/math"
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
		if dial == 0 {
			zeroHits += 1
		}
	}

	return zeroHits
}

func part2(data []string) int {
	dial := 50
	zeroHits := 0

	fmt.Println("dial -> 50")

	for i, line := range data {
		if len(line) < 2 {
			fmt.Printf("bad line %d: %s\n", i, line)
			continue
		}

		lastDialZero := dial == 0
		lastDial := dial
		rawVal := strings.ToInt(line[1:])
		val := rawVal
		if line[0] == 'L' {
			val = math.Abs(100 - val) % 100
		}
		dial = (dial + val) % 100
		revolutions := 0
		if line[0] == 'R'{
			revolutions = (lastDial+rawVal)/100
		} else if line[0] == 'L'{
			revolutions = (dial+rawVal)/100
		}

		fmt.Println(revolutions)

		correctHit := 0
		if lastDialZero && line[0] == 'L' {
			correctHit +=1
		}

		zeroHits += revolutions - correctHit
		if dial == 0 && revolutions == 0 {
			zeroHits += 1
		}
		fmt.Printf("%s r=%d: %d |0|=%d\n", line,revolutions, dial, zeroHits)
	}

	return zeroHits
}

func main() {
	data := input.LoadString("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println("bad: 31")
	fmt.Println(part1(data))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(data))
	fmt.Println("bad: 1889,2296,2347,2357,3492,5722,6016,6252")
}
