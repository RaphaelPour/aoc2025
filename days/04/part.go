package main

import (
	"fmt"

	"github.com/RaphaelPour/stellar/input"
)

type Grid struct{
	fields [][]bool
}

func NewGrid(input []string)Grid{
	fields := make([][]bool, len(input))

	for y, row := range input{
	  fields[y] = make([]bool, len(row))
	  for x, field := range row{
	    fields[y][x] = field == '@'
	  }
	}

	return Grid{fields:fields}
}

func (g Grid) Neighbours(x,y int) int{
	neighbours := 0
    for xOff := -1; xOff <=1;xOff++{
	    for yOff := -1; yOff <=1;yOff++{
		
		if y + yOff <0 || y+yOff >= len(g.fields){
			continue
		}
		if x + xOff <0 || x+xOff >= len(g.fields[y]){
			continue
		}
		if xOff == 0 && yOff == 0{
			continue
		}
		if g.fields[y+yOff][x+xOff]{
			neighbours +=1
		}
	    }
    }
    return neighbours
}

func part1(data []string) int {
	grid := NewGrid(data)
	forkliftable := 0
	for y := range grid.fields{
	  for x := range grid.fields[y]{
	    if grid.Neighbours(x,y) < 4{
		    fmt.Println(x,y,grid.Neighbours(x,y))
		forkliftable +=1
	    }
	  }
	}
    return forkliftable
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

    // fmt.Println("== [ PART 2 ] ==")
    // fmt.Println(part2(data))
}
