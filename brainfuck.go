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

		cellPointer = 0 // Index of cell.
	)

	// Instructions
	instructions, err := makeInstructions(code)
	if err != nil {
		return err
	}

	t := newTerminal(c.In, c.Out)

	for i := 0; i < len(instructions); i++ {

		if (cellPointer < 0) || (len(cells) <= cellPointer) {
			break
		}

		var (
			instruction = instructions[i]
			param       = instruction.Parameter
		)
		switch instruction.Op {
		case opIncPointer:
			cellPointer += param
		case opDecPointer:
			cellPointer -= param
		case opIncCell:
			cells[cellPointer] = byte(int(cells[cellPointer]) + param)
		case opDecCell:
			cells[cellPointer] = byte(int(cells[cellPointer]) - param)
		case opPutChar:
			for j := 0; j < param; j++ {
				b := cells[cellPointer]
				err := t.WriteByte(b)
				if err != nil {
					return err
				}
			}
		case opGetChar:
			for j := 0; j < param; j++ {
				b, err := t.ReadByte()
				if err != nil {
					return err
				}
				cells[cellPointer] = b
			}
		case opJumpIfZero:
			if cells[cellPointer] == 0 {
				i = param
			}
		case opJumpIfNotZero:
			if cells[cellPointer] != 0 {
				i = param
			}
		}
	}

	return nil
}
