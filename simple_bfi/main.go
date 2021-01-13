package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
)

const defaultCellSize = 30000

func Interpreter(code []byte, cellSize int) {

	var (
		cells   = make([]byte, cellSize)
		pointer = 0 // Index of cell.
	)

	var (
		commands     = []byte(code)
		commandIndex = 0
	)

	var (
		br = bufio.NewReader(os.Stdin)
		bw = bufio.NewWriter(os.Stdout)
	)
	defer bw.Flush()

	m := map[byte]func(){
		'>': func() { pointer++ },
		'<': func() { pointer-- },
		'+': func() { cells[pointer]++ },
		'-': func() { cells[pointer]-- },
		'.': func() {
			b := cells[pointer]
			err := bw.WriteByte(b)
			checkError(err)

			bw.Flush()
		},
		',': func() {
			b, err := br.ReadByte()
			checkError(err)
			cells[pointer] = b
		},
		'[': func() {
			if cells[pointer] == 0 {
				commandIndex = jumpIfZero(commands, commandIndex)
			}
		},
		']': func() {
			if cells[pointer] != 0 {
				commandIndex = jumpIfNotZero(commands, commandIndex)
			}
		},
	}

	for commandIndex < len(commands) {
		f, ok := m[commands[commandIndex]]
		if ok {
			f()
		}
		commandIndex++
	}
}

func jumpIfZero(commands []byte, commandIndex int) int {
	k := 1
	for k > 0 {
		commandIndex++
		switch commands[commandIndex] {
		case '[':
			k++
		case ']':
			k--
		}
	}
	return commandIndex
}

func jumpIfNotZero(commands []byte, commandIndex int) int {
	k := 1
	for k > 0 {
		commandIndex--
		switch commands[commandIndex] {
		case ']':
			k++
		case '[':
			k--
		}
	}
	return commandIndex
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	if len(os.Args) < 2 {
		log.Fatal("number of arguments < 2")
	}

	filename := os.Args[1]
	code, err := ioutil.ReadFile(filename)
	checkError(err)

	Interpreter(code, defaultCellSize)
}
