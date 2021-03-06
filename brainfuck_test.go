package brainfuck

import (
	"bytes"
	"errors"
	"io"
	"testing"
)

type zeroReader struct{}

var _ io.Reader = zeroReader{}

func (zeroReader) Read(p []byte) (n int, err error) {
	for i := range p {
		p[i] = 0
	}
	return len(p), nil
}

func TestHelloWorld(t *testing.T) {
	var (
		code = []byte(`++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.
+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.`)
		result = []byte("Hello World!\n")
	)
	buffer := new(bytes.Buffer)
	config := Config{
		RamSize: RamSizeDefault,
		Out:     buffer,
		In:      zeroReader{},
	}
	err := RunConfig(config, code)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(buffer.Bytes(), result) {
		t.Fatal("result is wrong:", string(buffer.Bytes()))
	}
}

func TestQuine(t *testing.T) {

	code := []byte(`>>>>>+++++>+++++>+++++>+++++>+++++>>++++++>++++>+++++++>++++>++++++>++++>+++++++
>++++>++++++>++++>+++++++>+++++>++++++>++++>++++++>++>+++++++>+>+>+>+>+>+>+>+>+>
+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>
+>+>+>+>+>+>+>+>+>+>+>+++++>++++++>++++>++>++>++>++>++>++>++>++>++>++>++>++>++>+
+>++>++>++>+++++>++++>++++>+>+++++>+++++>++>++++++>++++>+>+>+++++>++++>++++>+>++
+++>+++++>++>++++++>++++>+>+++++>++++>++++>+>+++++>+++++>++>++++++>++++>+>+>+>+>
+>+>+>+>+>+>+>+>+>+>+++++>++++>++++>+>+++++>+++++>++>++++++>++++>+>+>+++++>++++>
++++>+>+++++>+++++>++>++++++>++++>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>
+>+>+>+>+>+>+++++>++++>++++>+>+++++>+++++>++>++++++>++++>+>+>+++++>++++>++++>+>+
++++>+++++>++>++++++>++>+++++++>+++++++>+++++++>+++++++>+++++++>+++++++>+++++++>
+++++++>++++>+++>++++++>++>+++++++>+++++>+++++>+++++++>++++>++++>++++>++++++>+++
+>+++++++>+++++>++++++>++++++>++++>++++>+>+++++>+++++>++++>++++++>++>+++++++>+>+
>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+
>+>+++>++++++>++>+++++++>+++++>++>+++++++>++++++>++>+++++++>+>+>+>+>+>+>+>+>+>+>
+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>
+>+>+>+>+>+>+>+>+>+>+>+>+++>++++++>++>+++++++>+++++>+++++++>++++++>++>+++++++>+>
+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>
+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+++>++++++>++>+++++++>+++++>+++++>++++
+>++++++>++++++>++++>++++>+>+++++>+++++>++++>++++++>++>+++++++>+>+>+>+>+>+>+>+>+
>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+++>++++++>
++>+++++++>+++++>++>+++++++>++++++>++>+++++++>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>
+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>
+>+>+>+>+>+++>++++++>++>+++++++>+++++>+++++++>++++>++++>++++>++++++>++++>+++++++
>+++++>++++++>++++>++++++>++>+++++++>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+
>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+++
++>++++++>++++>++>++>++>++>++>++>++>++>++>++>++>++>++>++>++>++>++>+++++>++++>+++
+>+>+++++>+++++>++>++++++>++++>+>+>+++++>++++>++++>+>+++++>+++++>++>++++++>++++>
+>+++++>++++>++++>+>+++++>+++++>++>++++++>++++>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+++++
>++++>++++>+>+++++>+++++>++>++++++>++++>+>+>+++++>++++>++++>+>+++++>+++++>++>+++
+++>++++>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+>+++++>++++>++
++>+>+++++>+++++>++>++++++>++++>+>+>+++++>++++>++++>+>+++++>+++++>++>++++++>++>+
++++++>+++++++>+++++++>+++++++>+++++++>+++++++>+++++++>+++++++>++++>+++>++++++>+
+>+++++++>+++++>+++++>+++++++>[<]<[<]<[<]>[<[-]+++++++++++++++++++++++++++++++++
+++++++++++++++++++++++++++>[<-----------------><<+>>-[<++><<+>>-[<+><<+>>-[<+++
+++++++++++><<+>>-[<++><<+>>-[<+++++++++++++++++++++++++++++><<+>>-[<++><<+>>-[-
]]]]]]]]<.[-]>>]<<<[<]>[[<<+>><[-]+++++++++++++++++++++++++++++++++++++++++++.[-
]>-][-]++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.[-]>][-]++
++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.[-]>>>[[<<+>><[-]++
+++++++++++++++++++++++++++++++++++++++++.[-]>-][-]+++++++++++++++++++++++++++++
+++++++++++++++++++++++++++++++++.[-]>]<<<[<]>[<[-]+++++++++++++++++++++++++++++
+++++++++++++++++++++++++++++++>[<-----------------><<+>>-[<++><<+>>-[<+><<+>>-[
<++++++++++++++><<+>>-[<++><<+>>-[<+++++++++++++++++++++++++++++><<+>>-[<++><<+>
>-[-]]]]]]]]<.[-]>>]
`)
	err := testQuine(code)
	if err != nil {
		t.Fatal(err)
	}

	code = []byte(`->+>+++>>+>++>+>+++>>+>++>>>+>+>+>++>+>>>>+++>+>>++>+>+++>>++>++>>+>>+>++>++>+>>
>>+++>+>>>>++>++>>>>+>>++>+>+++>>>++>>++++++>>+>>++>+>>>>+++>>+++++>>+>+++>>>++>
>++>>+>>++>+>+++>>>++>>+++++++++++++>>+>>++>+>+++>+>+++>>>++>>++++>>+>>++>+>>>>+
++>>+++++>>>>++>>>>+>+>++>>+++>+>>>>+++>+>>>>+++>+>>>>+++>>++>++>+>+++>+>++>++>>
>>>>++>+>+++>>>>>+++>>>++>+>+++>+>+>++>>>>>>++>>>+>>>++>+>>>>+++>+>>>+>>++>+>+++
+++++++++++++++>>>>+>+>>>+>>++>+>+++>>>++>>++++++++>>+>>++>+>>>>+++>>++++++>>>+>
++>>+++>+>+>++>+>+++>>>>>+++>>>+>+>>++>+>+++>>>++>>++++++++>>+>>++>+>>>>+++>>+++
+>>+>+++>>>>>>++>+>+++>>+>++>>>>+>+>++>+>>>>+++>>+++>>>+[[->>+<<]<+]+++++[->++++
+++++<]>.[+]>>[<<+++++++[->+++++++++<]>-.------------------->-[-<.<+>>]<[+]<+>>>
]<<<[-[-[-[>>+<++++++[->+++++<]]>++++++++++++++<]>+++<]++++++[->+++++++<]>+<<<-[
->>>++<<<]>[->>.<<]<<]
`)
	err = testQuine(code)
	if err != nil {
		t.Fatal(err)
	}
}

func testQuine(code []byte) error {

	code = OnlyCode(code)

	var buf bytes.Buffer

	config := Config{
		RamSize: RamSizeDefault,
		Out:     &buf,
		In:      zeroReader{},
	}

	err := RunConfig(config, code)
	if err != nil {
		return err
	}

	if !bytes.Equal(buf.Bytes(), code) {
		return errors.New("code is not quine")
	}
	return nil
}
