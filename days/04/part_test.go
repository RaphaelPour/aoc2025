package main

import (
	"reflect"
	"testing"

	"github.com/RaphaelPour/stellar/input"
	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	require.Equal(t, 13, part1(input.LoadString("input_test")))
}

func TestNewGrid(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want Grid
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGrid(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrid_Within(t *testing.T) {
	type fields struct {
		fields [][]bool
	}
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Grid{
				fields: tt.fields.fields,
			}
			if got := g.Within(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Grid.Within() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrid_String(t *testing.T) {
	type fields struct {
		fields [][]bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Grid{
				fields: tt.fields.fields,
			}
			if got := g.String(); got != tt.want {
				t.Errorf("Grid.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrid_Neighbours(t *testing.T) {
	type fields struct {
		fields [][]bool
	}
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Grid{
				fields: tt.fields.fields,
			}
			if got := g.Neighbours(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Grid.Neighbours() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.data); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.data); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
