package brainfuck

import (
	"errors"
	"io"
	"os"
)

var (
	ErrMissingOpenSquareBracket  = errors.New("brainfuck: missing operator [")
	ErrMissingCloseSquareBracket = errors.New("brainfuck: missing operator ]")
)

type Config struct {
	RamSize int
	Out     io.Writer
	In      io.Reader
}

var Default = Config{
	RamSize: 30000,
	Out:     os.Stdout,
	In:      os.Stdin,
}

func Run(code []byte) error {
	return RunConfig(Default, code)
}

func RunConfig(c Config, code []byte) error {

	code = prepareOnlyCode(code)

	var (
		ram = make([]byte, c.RamSize)
		pos int
	)

	openJump, closeJump, err := makeJumpMaps(code)
	if err != nil {
		return err
	}

	t := newTerminal(c.In, c.Out)

	i := 0
	for n := len(code); i < n; {
		switch b := code[i]; b {
		case '>':
			pos = rInc(pos, c.RamSize)
		case '<':
			pos = rDec(pos, c.RamSize)
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

func rInc(i int, ramSize int) int {
	i++
	if i == ramSize {
		i = 0
	}
	return i
}

func rDec(i int, ramSize int) int {
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
