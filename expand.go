package brainfuck

import "bytes"

func Collapse(code []byte) []byte {
	return prepareOnlyCode(code)
}

func Expand(code []byte, indent string) ([]byte, error) {
	code = prepareOnlyCode(code)
	var buffer = new(bytes.Buffer)
	in := 0
	for _, b := range code {
		switch b {
		case '[':
			{
				buffer.WriteByte('\n')
				writeIndents(buffer, indent, in)
				buffer.WriteByte('[')
				buffer.WriteByte('\n')
				in++
				writeIndents(buffer, indent, in)
			}
		case ']':
			{
				in--
				if in < 0 {
					return nil, ErrMissingOpenSquareBracket
				}
				buffer.WriteByte('\n')
				writeIndents(buffer, indent, in)
				buffer.WriteByte(']')
				buffer.WriteByte('\n')
				writeIndents(buffer, indent, in)
			}
		default:
			buffer.WriteByte(b)
		}
	}
	if in > 0 {
		return nil, ErrMissingCloseSquareBracket
	}
	return buffer.Bytes(), nil
}

func writeIndents(buffer *bytes.Buffer, indent string, n int) {
	for i := 0; i < n; i++ {
		buffer.WriteString(indent)
	}
}
