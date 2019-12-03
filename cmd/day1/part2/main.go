package main

import (
	"io/ioutil"
	"log"
	"math"
	"strings"
)

func fuel(module int64) int64 {
	return int64(math.Floor(float64(module)/3) - 2)
}

func main() {
	content, err := ioutil.ReadFile("cmd/day1/input")
	if err != nil {
		log.Fatal(err)
	}
	data := strings.Split(string(content), ",")
	data[1] = "12"
	data[2] = "2"
}
