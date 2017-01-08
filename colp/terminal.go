package main

import (
	"fmt"
	"io"
)

type Terminal struct {
	w   io.Writer
	err error
}

func NewTerminal(out io.Writer) *Terminal {
	return &Terminal{w: out}
}

func (t *Terminal) Clear() {
	_, t.err = t.w.Write([]byte("\x1B[2J"))
}

func (t *Terminal) MoveTo(x, y int) {
	if x < 1 {
		x = 1
	}
	if y < 1 {
		y = 1
	}
	_, t.err = fmt.Fprintf(t.w, "\x1B[%d;%dH", y, x)
}

func (t *Terminal) WriteByte(b byte) {
	_, t.err = t.w.Write([]byte{b})
}

func (t *Terminal) WriteString(s string) {
	_, t.err = t.w.Write([]byte(s))
}

func (t *Terminal) Err() error {
	return t.err
}
