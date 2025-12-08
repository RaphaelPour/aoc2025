package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/RaphaelPour/stellar/input"
	smath "github.com/RaphaelPour/stellar/math"
	sstrings "github.com/RaphaelPour/stellar/strings"
)

type Point struct {
	x, y, z int
}

func ParsePoint(in string) Point {
	coords := strings.Split(strings.TrimSpace(in), ",")
	if len(coords) != 3 {
		panic(fmt.Sprintf("expected 3 elements, got %d: %s\n", len(coords), in))
	}

	return Point{
		x: sstrings.ToInt(coords[0]),
		y: sstrings.ToInt(coords[1]),
		z: sstrings.ToInt(coords[2]),
	}
}

func (p Point) Dist(other Point) float64 {
	return math.Sqrt(
		float64(smath.Pow(p.x-other.x, 2)) +
			float64(smath.Pow(p.y-other.y, 2)) +
			float64(smath.Pow(p.z-other.z, 2)),
	)
}

func (p Point) Equal(other Point) bool {
	return p.x == other.x &&
		p.y == other.y &&
		p.z == other.z
}

func (p Point) String() string {
	return fmt.Sprintf("%d,%d,%d", p.x, p.y, p.z)
}

type Circuit struct {
	boxes []Point
}

func (c *Circuit) AddOnce(box Point) bool {
	if c.Contains(box) {
		return false
	}

	c.boxes = append(c.boxes, box)
	return true
}

func (c Circuit) Contains(other Point) bool {
	for _, box := range c.boxes {
		if box.Equal(other) {
			return true
		}
	}

	return false
}

func (c Circuit) ContainsAny(others []Point) bool {
	for _, other := range others {
		if c.Contains(other) {
			return true
		}
	}
	return false
}

func (c Circuit) FindMissing(others []Point) Point {
	for _, other := range others {
		if !c.Contains(other) {
			return other
		}
	}
	panic("nothing is missing")
}

func part1(data []string) int {
	boxes := make([]Point, len(data))
	for i, line := range data {
		boxes[i] = ParsePoint(line)
	}

	distances := make(map[float64][]Point)
	distancesKeys := make([]float64, 0)

	for i := range len(boxes) {
		for j := i + 1; j < len(boxes); j++ {
			dist := boxes[i].Dist(boxes[j])
			distances[dist] = []Point{boxes[i], boxes[j]}
			distancesKeys = append(distancesKeys, dist)
		}
	}

	sort.Float64s(distancesKeys)

	// print all distances with their box pair from low to high
	// TODO: Start with the lowest distance and add the pairs into
	//       circuits. Check if one of the involved boxes are already
	//       within a circuit and just add it there.
	circuits := make([]*Circuit, 0)
	for _, key := range distancesKeys {
		fmt.Println(key, distances[key])
		var foundCircuit *Circuit
		val := distances[key]
		for _, circuit := range circuits {
			if circuit.ContainsAny(val) {
				foundCircuit = circuit
				break
			}
		}

		fmt.Print(circuits)
		// no circuit found
		if foundCircuit == nil {
			circuits = append(circuits, &Circuit{boxes: val})
		} else {
			// there is already a circuit, just add the missing one
			if !foundCircuit.AddOnce(foundCircuit.FindMissing(val)) {
				panic("moep")
			}
		}

		fmt.Println("->", circuits)
	}

	if len(distancesKeys) < 3 {
		panic(fmt.Sprintf("need at least three circuits, got %d\n", len(distancesKeys)))
	}

	maxes := make([]int, 0)
	for _, circuit := range circuits {
		maxes = append(maxes, len(circuit.boxes))
	}

	sort.Ints(maxes)
	fmt.Println(maxes)

	return smath.Product(smath.MaxN(maxes, 3))
}

func part2(data []string) int {
	return 0
}

func main() {
	data := input.LoadString("input_test")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	// fmt.Println("== [ PART 2 ] ==")
	// fmt.Println(part2(data))
}
