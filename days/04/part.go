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

func (g Grid) Within(x,y int) bool{
	  // neighbor can't have a negative position
	  if x <0{
			return false
	  }
	  if y <0{
			return false
	  }

	  // neighbor can't be outside the range
	  if x >= len(g.fields) || y >= len(g.fields){
			return false
	  }
	return true
}

func (g Grid) String() string{
	output := ""
	for y := range g.fields{
		for x := range g.fields[y]{
			if g.fields[y][x]{
				output += "@"
			} else{
				output += "."
			}
		}
		output += "\n"
	}

	return output
}

func (g Grid) Neighbours(x,y int) int{
	neighbours := 0
    for xOff := -1; xOff <=1;xOff++{
	    for yOff := -1; yOff <=1;yOff++{
			// Don't count itself
			if xOff == 0 && yOff == 0{
				continue
			}

		  // possible neighbor position
		  ox,oy := x+xOff,y+yOff

		  if !g.Within(ox,oy){
			continue
		  }

			if g.fields[oy][ox]{
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
	    if grid.fields[y][x] && grid.Neighbours(x,y) < 4{
		forkliftable +=1
	    }
	  }
	}
    return forkliftable
}

func part2(data []string) int {
	grid := NewGrid(data)
	count := 0
	for {
		forkliftable := 0
		for y := range grid.fields{
		  for x := range grid.fields[y]{
			if grid.fields[y][x] && grid.Neighbours(x,y) < 4{
				grid.fields[y][x] = false
				forkliftable +=1
				count +=1
			}
		  }
		}
		if forkliftable == 0{
			break
		}
	}
    return count
}

func main() {
	data := input.LoadDefaultString()
    
    fmt.Println("== [ PART 1 ] ==")
    fmt.Println(part1(data))

    fmt.Println("== [ PART 2 ] ==")
    fmt.Println(part2(data))
}
