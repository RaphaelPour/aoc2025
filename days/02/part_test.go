package main

import (
    "testing"

    "github.com/stretchr/testify/require"
	
	"github.com/RaphaelPour/stellar/input"
)

func TestExample(t *testing.T) {
 	require.Equal(
		t, 
		1227775554, 
		part1(input.LoadString("input_test")),
	)
}

func TestExampleLines(t *testing.T){
	require.Equal(
		t,
		33,
		part1([]string{"11-22"}),
	)
}
