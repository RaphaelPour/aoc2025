package main

import (
	"fmt"
	"math"
	"slices"
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
	return fmt.Sprintf("(%d %d %d)", p.x, p.y, p.z)
}

type Circuit struct {
	boxes []Point
}

func NewCircuit(boxes []Point) Circuit {
	c := Circuit{}
	c.boxes = make([]Point, len(boxes))
	copy(c.boxes, boxes)
	return c
}

func (c Circuit) String() string {
	output := ""
	for _, box := range c.boxes {
		output += box.String()
	}
	return output
}

func (c *Circuit) AddOnce(box Point) bool {
	if c.Contains(box) {
		return false
	}

	c.boxes = append(c.boxes, box)
	return true
}

func (c Circuit) Contains(that Point) bool {
	return slices.ContainsFunc(c.boxes, func(this Point) bool {
		return this.Equal(that)
	})
}

func (c Circuit) ContainsAny(others []Point) bool {
	return slices.ContainsFunc(others, c.Contains)
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
	fmt.Println(distancesKeys)

	// we only need the lowest 1000
	distancesKeys = distancesKeys[:1001]

	// print all distances with their box pair from low to high
	// TODO: Start with the lowest distance and add the pairs into
	//       circuits. Check if one of the involved boxes are already
	//       within a circuit and just add it there.
	circuits := make([]Circuit, 0)
	for _, key := range distancesKeys {
		fmt.Println("processing", key, distances[key])
		var foundIndex = -1
		pair := distances[key]
		for i, circuit := range circuits {
			if circuit.ContainsAny(pair) {
				fmt.Printf("found %s in %s\n", pair, circuit)
				foundIndex = i
				break
			}
		}

		// no circuit found
		if foundIndex == -1 {
			circuits = append(circuits, NewCircuit(pair))
		} else {
			circuits[foundIndex].AddOnce(pair[0])
			circuits[foundIndex].AddOnce(pair[1])
		}

		// TODO: don't consider just the pair but also the circuit it may be added to
		//       or consolidate pair and circuit
	}

	fmt.Printf("found %d circuits\n", len(circuits))

	if len(distancesKeys) < 3 {
		panic(fmt.Sprintf("need at least three circuits, got %d\n", len(distancesKeys)))
	}

	maxes := make([]int, 0)
	for _, circuit := range circuits {
		maxes = append(maxes, len(circuit.boxes))
	}

	sort.Ints(maxes)
	fmt.Println("max:", maxes)

	return smath.Product(smath.MaxN(maxes, 3))
}

func part2(data []string) int {
	return 0
}

func main() {
	data := input.LoadString("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))
	fmt.Println("too low: 2548")

	// fmt.Println("== [ PART 2 ] ==")
	// fmt.Println(part2(data))
}
