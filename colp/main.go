package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/toelsiba/brainfuck"
)

func main() {
	exampleExpand()
}

func exampleExpand() {

	var (
		//filename = "../src/factorial.bf"
		//filename = "../src/mandelbrot.bf"
		//filename = "../src/life.bf"
		//filename = "../src/hanoi.bf"
		filename = "../src/wiki_rot13.bf"
	)

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	//data = brainfuck.OnlyCode(data)
	code, err := brainfuck.Expand(data, "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(code))

	var (
		c1 = brainfuck.OnlyCode(data)
		c2 = brainfuck.OnlyCode(code)
	)
	if !bytes.Equal(c1, c2) {
		log.Fatal("Collapse-Expand error")
	}
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
