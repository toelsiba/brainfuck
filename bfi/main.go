package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/toelsiba/brainfuck"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("must be present argument brainfuck file (.b, .bf)")
	}
	fileName := os.Args[1]
	data, err := ioutil.ReadFile(fileName)
	checkError(err)
	err = brainfuck.Run(data)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
