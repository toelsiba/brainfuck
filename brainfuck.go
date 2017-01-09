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
		ramSize = c.RamSize
		ram     = make([]byte, ramSize)
		pos     int
	)

	ins, err := compileCode(code)
	if err != nil {
		return err
	}

	t := newTerminal(c.In, c.Out)

	i := 0
	for i < len(ins) {
		var (
			in    = ins[i]
			param = in.Parameter
		)
		switch in.Op {
		case op_IncPointer:
			pos = mod(pos+param, ramSize)
		case op_DecPointer:
			pos = mod(pos-param, ramSize)
		case op_Increment:
			ram[pos] = byte(int(ram[pos]) + param)
		case op_Decrement:
			ram[pos] = byte(int(ram[pos]) - param)
		case op_PutChar:
			for j := 0; j < param; j++ {
				if err = t.putChar(ram[pos]); err != nil {
					return err
				}
			}
		case op_GetChar:
			for j := 0; j < param; j++ {
				if ram[pos], err = t.getChar(); err != nil {
					return err
				}
			}
		case op_JumpIfZero:
			if ram[pos] == 0 {
				i = param
				continue
			}
		case op_JumpIfNotZero:
			if ram[pos] != 0 {
				i = param
				continue
			}
		}
		i++
	}

	return nil
}

func mod(x, y int) int {
	res := x % y
	if res < 0 {
		res += y
	}
	return res
}
