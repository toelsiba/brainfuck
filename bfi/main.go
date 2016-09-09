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
	if err != nil {
		log.Fatal(err)
	}
	if err = brainfuck.Run(data); err != nil {
		log.Fatal(err)
	}
}
