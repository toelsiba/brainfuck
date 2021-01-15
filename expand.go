package brainfuck

import "bytes"

var tableByteIsOperation = [256]bool{
	'>': true,
	'<': true,
	'+': true,
	'-': true,
	'.': true,
	',': true,
	'[': true,
	']': true,
}

func OnlyCode(code []byte) []byte {
	bs := make([]byte, 0, len(code))
	for _, b := range code {
		if tableByteIsOperation[b] {
			bs = append(bs, b)
		}
	}
	return bs
}

func Expand(code []byte, indent string) ([]byte, error) {
	var buffer bytes.Buffer
	in := 0
	var bracket bool
	for _, b := range code {
		switch b {
		case '[':
			{
				if !bracket {
					buffer.WriteByte('\n')
				}
				writeIndents(&buffer, indent, in)
				buffer.WriteByte('[')
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

func writeIndents(b *bytes.Buffer, indent string, n int) {
	for i := 0; i < n; i++ {
		b.WriteString(indent)
	}
}
