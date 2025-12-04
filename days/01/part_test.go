package main

import (
    "testing"
    "fmt"

    "github.com/stretchr/testify/require"
    "github.com/RaphaelPour/stellar/input"
)

func TestExample(t *testing.T) {
    // require.Equal(t, 0, part1([]string{})	
    require.Equal(t, 6, part2(input.LoadString("input_1")))
}

func Test2(t *testing.T) {
    for i, tc := range []struct{
	in []string
	expected int
    }{
	{in: []string{"L50"}, expected: 1},
	{in: []string{"L50","L1"}, expected: 1},
	{in: []string{"L49","L1","L1"}, expected: 1},
	{in: []string{"L100"}, expected: 1},
	{in: []string{"L200"}, expected: 2},
	{in: []string{"L200","R200"}, expected: 4},
	{in: []string{"L1","L55"}, expected: 1},
	{in: []string{"L50","L200"}, expected: 3},
    }{
	t.Run(fmt.Sprintf("%d",i), func(t *testing.T){
		require.Equal(t, tc.expected, part2(tc.in))
	})
    }
}
