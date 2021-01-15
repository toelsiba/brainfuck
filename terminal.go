package brainfuck

import "io"

type byteReadWriter interface {
	io.ByteReader
	io.ByteWriter
}

type terminal struct {
	buf []byte

	r io.Reader
	w io.Writer
}

func newTerminal(r io.Reader, w io.Writer) *terminal {
	return &terminal{
		buf: make([]byte, 1),
		r:   r,
		w:   w,
	}
}

var _ byteReadWriter = &terminal{}

func (t *terminal) ReadByte() (byte, error) {
	_, err := t.r.Read(t.buf)
	if err != nil {
		return 0, err
	}
	return t.buf[0], nil
}

func (t *terminal) WriteByte(b byte) error {
	t.buf[0] = b
	_, err := t.w.Write(t.buf)
	return err
}
