package brainfuck

import "errors"

var (
	ErrMissingOpenBracket  = errors.New("brainfuck: missing operator [")
	ErrMissingCloseBracket = errors.New("brainfuck: missing operator ]")
)

// Operations:

const (
	opIncPointer = iota
	opDecPointer
	opIncCell
	opDecCell
	opPutChar
	opGetChar
	opJumpIfZero
	opJumpIfNotZero
)

type instruction struct {
	Op        byte // Operation
	Parameter int
}

func makeInstructions(code []byte) (ins []*instruction, err error) {

	for _, b := range code {
		switch b {
		case '>':
			ins = appendInstruction(ins, opIncPointer)
		case '<':
			ins = appendInstruction(ins, opDecPointer)
		case '+':
			ins = appendInstruction(ins, opIncCell)
		case '-':
			ins = appendInstruction(ins, opDecCell)
		case '.':
			ins = appendInstruction(ins, opPutChar)
		case ',':
			ins = appendInstruction(ins, opGetChar)
		case '[':
			ins = append(ins, &instruction{Op: opJumpIfZero})
		case ']':
			ins = append(ins, &instruction{Op: opJumpIfNotZero})
		}
	}

	err = prepareJumps(ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

func appendInstruction(ins []*instruction, op byte) []*instruction {
	if n := len(ins); n > 0 {
		if last := ins[n-1]; last.Op == op {
			last.Parameter++
			return ins
		}
	}
	return append(ins, &instruction{Op: op, Parameter: 1})
}

func prepareJumps(ins []*instruction) error {
	var js []int
	for j, in := range ins {
		switch in.Op {
		case opJumpIfZero:
			js = append(js, j) // append index of '['
		case opJumpIfNotZero:
			if len(js) == 0 {
				return ErrMissingOpenBracket
			}
			k := len(js) - 1 // index of last '['

			ins[js[k]].Parameter = j // index of ']'
			ins[j].Parameter = js[k] // index of '['

			js = js[:k]
		}
	}
	if len(js) > 0 {
		return ErrMissingCloseBracket
	}
	return nil
}
