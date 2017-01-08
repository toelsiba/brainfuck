package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/toelsiba/brainfuck"
)

func main() {
	//outToFile()
	exampleEscapeSeq()
	//testNumber()
}

func exampleExpand() {
	data, err := ioutil.ReadFile("../src/factorial.bf")
	if err != nil {
		log.Fatal(err)
	}
	code, err := brainfuck.Expand(data, "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(code))
}

func outToFile() {
	data, err := ioutil.ReadFile("../src/hanoi.bf")
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("out.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	defer w.Flush()

	config := brainfuck.Config{
		RamSize: brainfuck.RamSizeDefault,
		Out:     w,
		In:      nil,
	}

	err = brainfuck.RunConfig(config, data)
	if err != nil {
		log.Fatal(err)
	}
}

func exampleEscapeSeq() {

	t := NewTerminal(os.Stdout)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	const (
		maxX = 50
		maxY = 25
	)

	var (
		x = 1 + r.Intn(maxX)
		y = 1 + r.Intn(maxY)
	)

	dx := 1
	if (r.Int() & 1) == 0 {
		dx = -1
	}
	dy := 1
	if (r.Int() & 1) == 0 {
		dy = -1
	}

	for i := 0; i < 400; i++ {
		t.Clear()
		t.MoveTo(x, y)
		t.WriteString("Hello, BrainF*ck!")
		t.MoveTo(1, 1)
		time.Sleep(time.Second / 30)

		x += dx
		if x < 1 {
			x = 2
			dx = -dx
		}
		if x > maxX {
			x = maxX - 1
			dx = -dx
		}

		y += dy
		if y < 1 {
			y = 2
			dy = -dy
		}
		if y > maxY {
			y = maxY - 1
			dy = -dy
		}
	}
}

func testNumber() {
	var min byte = 120
	b := min
	for i := 0; i < 1000; i++ {

		v := min + byte(i)
		//v := byte(int(min) + i)

		if b != v {
			fmt.Println(b, v)
		}

		b++
	}
}
