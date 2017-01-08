package brainfuck

import "errors"

var (
	ErrMissingOpenBracket  = errors.New("brainfuck: missing operator [")
	ErrMissingCloseBracket = errors.New("brainfuck: missing operator ]")
)

// Operations:

const (
	op_Right         = '>'
	op_Left          = '<'
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
		case op_Right:
			ins = appendInstruction(ins, op_Right)
		case op_Left:
			ins = appendInstruction(ins, op_Left)
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
	in := &instruction{
		Op:        op,
		Parameter: 1,
	}
	return append(ins, in)
}

func prepareJumps(ins []*instruction) error {
	var is []int
	for i, in := range ins {
		switch in.Op {
		case op_JumpIfZero:
			is = append(is, i)
		case op_JumpIfNotZero:
			if len(is) == 0 {
				return ErrMissingOpenBracket
			}
			k := len(is) - 1

			ins[is[k]].Parameter = i + 1
			ins[i].Parameter = is[k]

			is = is[:k]
		}
	}
	if len(is) > 0 {
		return ErrMissingCloseBracket
	}
	return nil
}
