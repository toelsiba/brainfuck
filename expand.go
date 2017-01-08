package brainfuck

import "bytes"

var tableByteIsOp = makeOpTable()

func Collapse(code []byte) []byte {
	return prepareOnlyCode(code)
}

func Expand(code []byte, indent string) ([]byte, error) {
	code = prepareOnlyCode(code)
	var buffer bytes.Buffer
	in := 0
	var bracket bool
	for _, b := range code {
		switch b {
		case op_JumpIfZero:
			{
				if !bracket {
					buffer.WriteByte('\n')
				}
				writeIndents(&buffer, indent, in)
				buffer.WriteByte(op_JumpIfZero)
				buffer.WriteByte('\n')
				in++
				bracket = true
			}
		case ']':
			{
				in--
				if in < 0 {
					return nil, ErrMissingOpenBracket
				}
				if !bracket {
					buffer.WriteByte('\n')
				}
				writeIndents(&buffer, indent, in)
				buffer.WriteByte(']')
				buffer.WriteByte('\n')
				bracket = true
			}
		default:
			if bracket {
				writeIndents(&buffer, indent, in)
			}
			buffer.WriteByte(b)
			bracket = false
		}
	}
	if in > 0 {
		return nil, ErrMissingCloseBracket
	}
	return buffer.Bytes(), nil
}

func writeIndents(buffer *bytes.Buffer, indent string, n int) {
	for i := 0; i < n; i++ {
		buffer.WriteString(indent)
	}
}

func prepareOnlyCode(code []byte) []byte {
	var bs []byte
	for i, b := range code {
		if tableByteIsOp[b] {
			if bs != nil {
				bs = append(bs, b)
			}
		} else {
			if bs == nil {
				bs = make([]byte, i)
				if len(bs) > 0 {
					copy(bs, code)
				}
			}
		}
	}
	if bs != nil {
		return bs
	}
	return code
}

func makeOpTable() (table [256]bool) {
	for i := range table {
		switch i {
		case op_Right:
			fallthrough
		case op_Left:
			fallthrough
		case op_Increment:
			fallthrough
		case op_Decrement:
			fallthrough
		case op_PutChar:
			fallthrough
		case op_GetChar:
			fallthrough
		case op_JumpIfZero:
			fallthrough
		case op_JumpIfNotZero:
			table[i] = true
		}
	}
	return
}
