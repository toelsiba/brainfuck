package brainfuck

import (
	"io"
	"os"
)

type Config struct {
	RamSize int
	Out     io.Writer
	In      io.Reader
}

const RamSizeDefault = 30000

var Default = Config{
	RamSize: RamSizeDefault,
	Out:     os.Stdout,
	In:      os.Stdin,
}

func Run(code []byte) error {
	return RunConfig(Default, code)
}

func RunConfig(c Config, code []byte) error {

	var (
		cells = make([]byte, c.RamSize) // Ram

		pointer = 0 // Index of cell.
	)

	// Instructions
	ins, err := makeInstructions(code)
	if err != nil {
		return err
	}

	t := newTerminal(c.In, c.Out)

	i := 0 // Index of instruction.
	for i < len(ins) {

		if (pointer < 0) || (len(cells) <= pointer) {
			break
		}

		var (
			in    = ins[i]
			param = in.Parameter
		)
		switch in.Op {
		case opIncPointer:
			pointer += param
		case opDecPointer:
			pointer -= param
		case opIncCell:
			cells[pointer] = byte(int(cells[pointer]) + param)
		case opDecCell:
			cells[pointer] = byte(int(cells[pointer]) - param)
		case opPutChar:
			for j := 0; j < param; j++ {
				if err = t.putChar(cells[pointer]); err != nil {
					return err
				}
			}
		case opGetChar:
			for j := 0; j < param; j++ {
				if cells[pointer], err = t.getChar(); err != nil {
					return err
				}
			}
		case opJumpIfZero:
			if cells[pointer] == 0 {
				i = param
			}
		case opJumpIfNotZero:
			if cells[pointer] != 0 {
				i = param
			}
		}
		i++
	}

	return nil
}
