package main

import (
	"fmt"
	"github.com/RaphaelPour/stellar/input"
	"github.com/RaphaelPour/stellar/math"
)

type Manifold struct{
	tachyons []math.Point
	diagram [][]bool
}

func (m *Manifold) AddTachyonOnce(p math.Point) bool{
	if len(m.tachyons) == 0 {
		m.tachyons = []math.Point{p}
		return true
	}

	if m.IsOccupied(p){
		return false
	}

	m.tachyons = append(m.tachyons, p)
	return true
}

func (m Manifold) IsOccupied(p math.Point) bool{
	for _, tachyon := range m.tachyons{
		if tachyon.Equal(p){
			return true
		}
	}
	return false
}

func part1(data []string) int {
	m := Manifold{}
	m.tachyons = make([]math.Point, 1)
	splits := 0
	output := make([][]rune,len(data))

	for y := range data{
		output[y] = make([]rune, len(data[y]))
		for x := range data[y]{
			output[y][x] = []rune(data[y])[x]
			if data[y][x] == 'S'{
				fmt.Println("S:", x,y)
				m.tachyons =  []math.Point{math.Point{x,y}}
			}
		}
	}

	// skip first row as we start there
	for yOff := 0; yOff < len(data); yOff+=1{
		origlen := len(m.tachyons)
		for i := 0; i<origlen;i++{
			tachyon := m.tachyons[i]

			switch data[yOff][tachyon.X] {
			case 'S':
				m.tachyons = append(m.tachyons, tachyon.Add(math.Point{0,1}))
			case '.':
				m.tachyons = append(m.tachyons, tachyon.Add(math.Point{0,1}))
				output[yOff][tachyon.X] = 'p'
			}
		}
		for i := 0; i<origlen;i++{
			tachyon := m.tachyons[i]

			switch data[yOff][tachyon.X] {
			case '^': 
				splits += 1
				m.AddTachyonOnce(tachyon.Add(math.Point{X:1,Y:1}))
				m.AddTachyonOnce(tachyon.Add(math.Point{X:-1,Y:1}))
				output[yOff][tachyon.X-1] = '|'
				output[yOff][tachyon.X] = output[yOff][tachyon.X]+1 
				output[yOff][tachyon.X+1] = '|'
			}
		}

		m.tachyons = m.tachyons[origlen:]
	}

	for _,row := range output{
		for _,cell := range row{
			fmt.Printf("%c",cell)
		}
		fmt.Println("")
	}

	return splits
}

func part2(data []string) int {
    return 0
}

func main() {
	// data := input.LoadString("input")
	// data := input.LoadDefaultInt()
	// data := input.LoadInt("input")
	data := input.LoadDefaultString()
    
    fmt.Println("== [ PART 1 ] ==")
    fmt.Println(part1(data))
	fmt.Println("too high: 1761")

    // fmt.Println("== [ PART 2 ] ==")
    // fmt.Println(part2(data))
}
