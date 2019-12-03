package aoc19

import (
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

func input() []byte {
	data, err := ioutil.ReadFile(path.Dir(os.Args[0]) + "/../input")
	if err != nil {
		panic(err)
	}
	return data
}

func intcode(input []byte) []int {
	data := strings.Split(string(input), ",")
	code := make([]int, len(data))
	for i, x := range data {
		var err error
		code[i], err = strconv.Atoi(strings.TrimSpace(x))
		if err != nil {
			panic(err)
		}
	}
	return code
}
