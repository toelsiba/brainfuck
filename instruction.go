package brainfuck

import "errors"

var (
	ErrMissingOpenBracket  = errors.New("brainfuck: missing operator [")
	ErrMissingCloseBracket = errors.New("brainfuck: missing operator ]")
)

// Operations:

const (
	op_IncPointer    = '>'
	op_DecPointer    = '<'
	op_Increment     = '+'
	op_Decrement     = '-'
	op_PutChar       = '.'
	op_GetChar       = ','
	op_JumpIfZero    = '['
	op_JumpIfNotZero = ']'
)

type instruction struct {
	Op        byte
	Parameter int
}

func compileCode(code []byte) (ins []*instruction, err error) {

	for _, b := range code {
		switch b {
		case op_IncPointer:
			ins = appendInstruction(ins, op_IncPointer)
		case op_DecPointer:
			ins = appendInstruction(ins, op_DecPointer)
		case op_Increment:
			ins = appendInstruction(ins, op_Increment)
		case op_Decrement:
			ins = appendInstruction(ins, op_Decrement)
		case op_PutChar:
			ins = appendInstruction(ins, op_PutChar)
		case op_GetChar:
			ins = appendInstruction(ins, op_GetChar)
		case op_JumpIfZero:
			ins = append(ins, &instruction{Op: op_JumpIfZero})
		case op_JumpIfNotZero:
			ins = append(ins, &instruction{Op: op_JumpIfNotZero})
		}
	}

	if err = prepareJumps(ins); err != nil {
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
		case op_JumpIfZero:
			js = append(js, j)
		case op_JumpIfNotZero:
			if len(js) == 0 {
				return ErrMissingOpenBracket
			}
			k := len(js) - 1

			ins[js[k]].Parameter = j + 1
			ins[j].Parameter = js[k]

			js = js[:k]
		}
	}
	if len(js) > 0 {
		return ErrMissingCloseBracket
	}
	return nil
}
