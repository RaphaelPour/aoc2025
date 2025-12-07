package main

import (
    "testing"

	"github.com/RaphaelPour/stellar/input"
        "github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
    require.Equal(t, 21, part1(input.LoadString("input_test")))
}
