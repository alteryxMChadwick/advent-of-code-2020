package code_tests

import (
	"testing"

	code "github.com/capsocrates/advent-of-code/day01"
	"github.com/stretchr/testify/assert"
)

func TestDay01(t *testing.T) {
	t.Run("TestEmpty", func(t *testing.T) {
		left, right := code.GetPair([]int{})

		assert.Equal(t, 0, left)
		assert.Equal(t, 0, right)
	})

	t.Run("TestSingleMatchingElement", func(t *testing.T) {
		left, right := code.GetPair([]int{2020})

		assert.Equal(t, 0, left)
		assert.Equal(t, 0, right)
	})

	t.Run("TestTwoMatchingElements", func(t *testing.T) {
		left, right := code.GetPair([]int{1010, 1010})

		assert.Equal(t, 1010, left)
		assert.Equal(t, 1010, right)
	})

	t.Run("TestTwoMatchingElementsAgain", func(t *testing.T) {
		left, right := code.GetPair([]int{0, 2020})

		assert.Equal(t, 2020, left)
		assert.Equal(t, 0, right)
	})

	t.Run("TestNoMatchingElement", func(t *testing.T) {
		left, right := code.GetPair([]int{0})

		assert.Equal(t, 0, left)
		assert.Equal(t, 0, right)
	})

	t.Run("ExampleSet", func(t *testing.T) {
		exampleNums := []int{
			1721,
			979,
			366,
			299,
			675,
			1456,
		}

		left, right := code.GetPair(exampleNums)

		assert.Equal(t, 1721, left)
		assert.Equal(t, 299, right)

		product := code.GetProduct(left, right)
		assert.Equal(t, 514579, product)
	})
}
