package day08_tests

import (
	"github.com/capsocrates/advent-of-code/day08"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	exampleInput = `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`
)

func TestParseInstructions(t *testing.T) {
	parsed, err := day08.ParseInstructions(exampleInput)
	require.Equal(t, nil, err)
	require.Equal(t, 9, len(parsed))
	assert.Equal(t, day08.NOP, parsed[0].Instruction)
	assert.Equal(t, 0, parsed[0].Value)
	assert.Equal(t, 0, parsed[0].RunCount)
	assert.Equal(t, day08.ACC, parsed[1].Instruction)
	assert.Equal(t, 1, parsed[1].Value)
	assert.Equal(t, 0, parsed[1].RunCount)
	assert.Equal(t, day08.JMP, parsed[2].Instruction)
	assert.Equal(t, 4, parsed[2].Value)
	assert.Equal(t, 0, parsed[2].RunCount)
	assert.Equal(t, day08.ACC, parsed[3].Instruction)
	assert.Equal(t, 3, parsed[3].Value)
	assert.Equal(t, 0, parsed[3].RunCount)
	assert.Equal(t, day08.JMP, parsed[4].Instruction)
	assert.Equal(t, -3, parsed[4].Value)
	assert.Equal(t, 0, parsed[4].RunCount)
	assert.Equal(t, day08.ACC, parsed[5].Instruction)
	assert.Equal(t, -99, parsed[5].Value)
	assert.Equal(t, 0, parsed[5].RunCount)
	assert.Equal(t, day08.ACC, parsed[6].Instruction)
	assert.Equal(t, 1, parsed[6].Value)
	assert.Equal(t, 0, parsed[6].RunCount)
	assert.Equal(t, day08.JMP, parsed[7].Instruction)
	assert.Equal(t, -4, parsed[7].Value)
	assert.Equal(t, 0, parsed[7].RunCount)
	assert.Equal(t, day08.ACC, parsed[8].Instruction)
	assert.Equal(t, 6, parsed[8].Value)
	assert.Equal(t, 0, parsed[8].RunCount)
}

func TestRunOncePart1Day08(t *testing.T) {
	parsed, err := day08.ParseInstructions(exampleInput)
	require.Equal(t, nil, err)
	runner := day08.NewRunner(parsed)
	runner.RunOnce()
	assert.Equal(t, 5, runner.Accumulator())
}
