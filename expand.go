package brainfuck

import "bytes"

func Collapse(code []byte) []byte {
	return prepareOnlyCode(code)
}

func Expand(code []byte, indent string) ([]byte, error) {
	code = prepareOnlyCode(code)
	var buffer = new(bytes.Buffer)
	in := 0
	var bracket bool
	for _, b := range code {
		switch b {
		case '[':
			{
				if !bracket {
					buffer.WriteByte('\n')
				}
				writeIndents(buffer, indent, in)
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
				writeIndents(buffer, indent, in)
				buffer.WriteByte(']')
				buffer.WriteByte('\n')
				bracket = true
			}
		default:
			if bracket {
				writeIndents(buffer, indent, in)
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
