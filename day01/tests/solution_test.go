package code_tests

import (
	"testing"

	code "github.com/capsocrates/advent-of-code/day01"
	"github.com/stretchr/testify/require"
)

func TestDay01(t *testing.T) {
	exampleNums := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}

	left, right := code.GetPair(exampleNums)

	require.Equal(t, 1721, left)
	require.Equal(t, 299, right)

	product := code.GetProduct(left, right)
	require.Equal(t, 514579, product)
}
