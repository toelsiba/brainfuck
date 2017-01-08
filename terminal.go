package brainfuck

import "io"

type terminal struct {
	bs []byte
	r  io.Reader
	w  io.Writer
}

func newTerminal(r io.Reader, w io.Writer) *terminal {
	return &terminal{
		bs: make([]byte, 1),
		r:  r,
		w:  w,
	}
}

func (t *terminal) putChar(b byte) error {
	t.bs[0] = b
	_, err := t.w.Write(t.bs)
	return err
}

func (t *terminal) getChar() (byte, error) {
	_, err := t.r.Read(t.bs)
	if err != nil {
		return 0, err
	}
	return t.bs[0], nil
}
