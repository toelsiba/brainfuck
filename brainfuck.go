package brainfuck

import (
	"errors"
	"fmt"
	"io"
	"os"
)

var ErrPointer = errors.New("cell pointer out of range")

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

	if c.RamSize <= 0 {
		return fmt.Errorf("Invalid cells size %d", c.RamSize)
	}

	var (
		cells = make([]byte, c.RamSize) // Ram

		pointer = 0 // Index of cell.
	)

	// Instructions
	instructions, err := makeInstructions(code)
	if err != nil {
		return err
	}

	var brw byteReadWriter
	brw = newTerminal(c.In, c.Out)

	for i := 0; i < len(instructions); i++ {

		var (
			instruction = instructions[i]
			param       = instruction.Parameter
		)

		switch instruction.Op {

		case opIncPointer:
			{
				pointer += param
				if pointer >= len(cells) {
					return ErrPointer
				}
			}

		case opDecPointer:
			{
				pointer -= param
				if pointer < 0 {
					return ErrPointer
				}
			}

		case opIncCell:
			cells[pointer] = byte(int(cells[pointer]) + param)

		case opDecCell:
			cells[pointer] = byte(int(cells[pointer]) - param)

		case opPutChar:
			for j := 0; j < param; j++ {
				b := cells[pointer]
				err := brw.WriteByte(b)
				if err != nil {
					return err
				}
			}

		case opGetChar:
			for j := 0; j < param; j++ {
				b, err := brw.ReadByte()
				if err != nil {
					return err
				}
				cells[pointer] = b
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
	}

	return nil
}
