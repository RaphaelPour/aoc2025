package main

import (

	"strings"
	"strconv"
	"fmt"
	"iter"

	"github.com/RaphaelPour/stellar/input"
)

func Must[T any](in T, err error)T{
	if err != nil{
		panic(err)
	}
	return in
}

type Range struct{
	start,end int
}

func Check(n int)bool{
	s:= strconv.Itoa(n)
	if len(s) % 2 != 0{
		return false
	}

	return strings.Compare(s[0:len(s)/2],s[len(s)/2:]) == 0
}

func (r Range) InvalidIDs() iter.Seq[int]{
	return func(yield func(int)bool){
		for i := r.start;i <=r.end; i++{
			if Check(i){
				if !yield(i){
					return
				}
			}
		} 
	}
}

func NewRange(in string) Range{
	parts := strings.Split(in, "-")
	if len(parts) != 2{
		panic(fmt.Sprintf("invalid range %s", in))
	}
	return Range{
		start: Must(strconv.Atoi(parts[0])),
		end: Must(strconv.Atoi(parts[1])),
	}
}

func part1(data []string) int {
	if len(data) != 1{
		panic(fmt.Sprintf("expteded exactly one line, got %d",len(data)))
	}

	sum :=0
	for _, line := range strings.Split(data[0],","){
		r := NewRange(line)
		for  num := range r.InvalidIDs(){
			sum +=num
		}
	}
	return sum
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
