package brainfuck

import (
	"bytes"
	"testing"
)

type zeroReader struct{}

func (zeroReader) Read(p []byte) (n int, err error) {
	for i := range p {
		p[i] = 0
	}
	return len(p), nil
}

func TestHelloWorld(t *testing.T) {
	var (
		code   = []byte("--[+++++++<---->>-->+>+>+<<<<]<.>++++[-<++++<++>>>->--<<]>>-.>--..>+.<<<.<<-.>>+>->>.+++[.<]<<++.")
		result = []byte("Hello World!\n")
	)
	buffer := new(bytes.Buffer)
	config := Config{
		RamSize: RAM_SIZE_DEFAULT,
		Out:     buffer,
		In:      zeroReader{},
	}
	err := RunConfig(config, code)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(buffer.Bytes(), result) {
		t.Fatal("result is wrong")
	}
}

func TestQuine(t *testing.T) {
	code := []byte(`
->+>+++>>+>++>+>+++>>+>++>>>+>+>+>++>+>>>>+++>+>>++>+>+++>>++>++
>>+>>+>++>++>+>>>>+++>+>>>>++>++>>>>+>>++>+>+++>>>++>>++++++>>+>
>++>+>>>>+++>>+++++>>+>+++>>>++>>++>>+>>++>+>+++>>>++>>+++++++++
++++>>+>>++>+>+++>+>+++>>>++>>++++>>+>>++>+>>>>+++>>+++++>>>>++>
>>>+>+>++>>+++>+>>>>+++>+>>>>+++>+>>>>+++>>++>++>+>+++>+>++>++>>
>>>>++>+>+++>>>>>+++>>>++>+>+++>+>+>++>>>>>>++>>>+>>>++>+>>>>+++
>+>>>+>>++>+>++++++++++++++++++>>>>+>+>>>+>>++>+>+++>>>++>>+++++
+++>>+>>++>+>>>>+++>>++++++>>>+>++>>+++>+>+>++>+>+++>>>>>+++>>>+
>+>>++>+>+++>>>++>>++++++++>>+>>++>+>>>>+++>>++++>>+>+++>>>>>>++
>+>+++>>+>++>>>>+>+>++>+>>>>+++>>+++>>>+[[->>+<<]<+]+++++[->++++
+++++<]>.[+]>>[<<+++++++[->+++++++++<]>-.------------------->-[-
<.<+>>]<[+]<+>>>]<<<[-[-[-[>>+<++++++[->+++++<]]>++++++++++++++<
]>+++<]++++++[->+++++++<]>+<<<-[->>>++<<<]>[->>.<<]<<]
`)
	code = Collapse(code)

	buffer := new(bytes.Buffer)

	config := Config{
		RamSize: RAM_SIZE_DEFAULT,
		Out:     buffer,
		In:      zeroReader{},
	}

	err := RunConfig(config, code)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(buffer.Bytes(), code) {
		t.Fatal("result is wrong")
	}
}
