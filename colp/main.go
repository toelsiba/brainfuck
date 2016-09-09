package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/toelsiba/brainfuck"
)

func main() {
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
