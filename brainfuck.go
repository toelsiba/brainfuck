package brainfuck

import "errors"

var (
	ErrMissingOpenSquareBracket  = errors.New("brainfuck: missing operator [")
	ErrMissingCloseSquareBracket = errors.New("brainfuck: missing operator ]")
)

const ramSize = 30000

func Run(code []byte) error {

	code = prepareOnlyCode(code)

	var (
		ram [ramSize]byte
		pos int
	)

	openJump, closeJump, err := makeJumpMaps(code)
	if err != nil {
		return err
	}

	t := newTerminalStd()

	i := 0
	for n := len(code); i < n; {
		switch b := code[i]; b {
		case '>':
			pos = rInc(pos)
		case '<':
			pos = rDec(pos)
		case '+':
			ram[pos]++
		case '-':
			ram[pos]--
		case '.':
			if err = t.putchar(ram[pos]); err != nil {
				return err
			}
		case ',':
			if ram[pos], err = t.getchar(); err != nil {
				return err
			}
		case '[':
			if ram[pos] == 0 {
				i = openJump[i]
				continue
			}
		case ']':
			if ram[pos] != 0 {
				i = closeJump[i]
				continue
			}
		}
		i++
	}

	return nil
}

func byteIsOperator(b byte) bool {
	switch b {
	case '>', '<', '+', '-', '.', ',', '[', ']':
		return true
	default:
		return false
	}
}

func prepareOnlyCode(code []byte) []byte {
	var bs []byte
	for i, b := range code {
		if byteIsOperator(b) {
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

func rInc(i int) int {
	i++
	if i == ramSize {
		i = 0
	}
	return i
}

func rDec(i int) int {
	i--
	if i < 0 {
		i = ramSize - 1
	}
	return i
}

func makeJumpMaps(code []byte) (openJump, closeJump map[int]int, err error) {
	openJump = make(map[int]int)
	closeJump = make(map[int]int)
	var is []int
	for i := range code {
		switch b := code[i]; b {
		case '[':
			is = append(is, i)
		case ']':
			{
				if len(is) == 0 {
					err = ErrMissingOpenSquareBracket
					return
				}
				k := len(is) - 1
				openJump[is[k]] = i + 1
				closeJump[i] = is[k]
				is = is[:k]
			}
		}
	}
	if len(is) > 0 {
		err = ErrMissingCloseSquareBracket
		return
	}
	return
}
