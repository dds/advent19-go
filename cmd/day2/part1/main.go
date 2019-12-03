package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("cmd/day2/input")
	if err != nil {
		log.Fatal(err)
	}
	data := strings.Split(string(content), ",")
	program := make([]int, len(data))
	for i, x := range data {
		var err error
		program[i], err = strconv.Atoi(strings.TrimSpace(x))
		if err != nil {
			panic(err)
		}
	}
	program[1] = 12
	program[2] = 2

	program = run(program)
	fmt.Println(program[0])
}

func run(data []int) []int {
	for i := 0; ; {
		switch data[i] {
		case 99:
			// halt
			return data
		case 1:
			// add
			data[data[i+3]] = data[data[i+1]] + data[data[i+2]]
			i += 4
		case 2:
			// multiply
			data[data[i+3]] = data[data[i+1]] * data[data[i+2]]
			i += 4
		}
	}
}
