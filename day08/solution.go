package day08

import (
	"errors"
	"strconv"
	"strings"
)

type Command int

const (
	NOP Command = iota
	JMP
	ACC
)

type Instruction struct {
	RunCount int
	Instruction Command
	Value int
}

type Runner struct {
	accumulator int
	instructions []Instruction
}

func ParseCommand(cmd string) (Command, error) {
	switch cmd {
	case "nop":
		return NOP, nil
	case "jmp":
		return JMP, nil
	case "acc":
		return ACC, nil
	default:
		return -1, errors.New("invalid command")
	}
}

func ParseInstructions(i string) ([]Instruction, error) {
	ins := make([]Instruction, 0, len(i)/7)
	for _, instruction := range strings.Split(i, "\n") {
		parts := strings.Split(instruction, " ")
		if 2 != len(parts) {
			return nil, errors.New("invalid instruction")
		}
		x, err := strconv.ParseInt(parts[1], 10, 32)
		if nil != err {
			return nil, err
		}
		cmd, err := ParseCommand(parts[0])
		if nil != err {
			return nil, err
		}
		ins = append(ins, Instruction{
			RunCount: 0,
			Instruction: cmd,
			Value: int(x),
		})
	}
	return ins, nil
}

func NewRunner(i []Instruction) Runner {
	return Runner{0, i}
}

func (r *Runner) RunOnce() error {
	idx := 0
	instructionCount := len(r.instructions)
	if instructionCount < 1 {
		return errors.New("no instructions")
	}
	for {
		instruction := &r.instructions[idx]
		if 0 < instruction.RunCount {
			break
		}
		instruction.RunCount++
		switch instruction.Instruction {
		case NOP:
			println(idx, "-", r.accumulator, ":", "NOP", instruction.Value, instruction.RunCount)
			idx++
		case JMP:
			println(idx, "-", r.accumulator, ":", "JMP", instruction.Value, instruction.RunCount)
			idx += instruction.Value
		case ACC:
			println(idx, "-", r.accumulator, ":", "ACC", instruction.Value, instruction.RunCount)
			idx++
			r.accumulator += instruction.Value
		default:
			return errors.New("invalid command")
		}

		if instructionCount <= idx || idx < 0 {
			return errors.New("instruction out of bounds")
		}
	}
	return nil
}

func (r *Runner) Accumulator() int {
	return r.accumulator
}
