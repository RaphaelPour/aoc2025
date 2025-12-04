package main

import (
    "testing"

        "github.com/stretchr/testify/require"

	"github.com/RaphaelPour/stellar/input"
)

func TestExample(t *testing.T) {
    require.Equal(t, 13, part1(input.LoadString("input_test")))
}
