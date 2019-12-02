package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
)

func fuel(module int64) int64 {
	return int64(math.Floor(float64(module)/3) - 2)
}

func main() {
	content, err := ioutil.ReadFile("cmd/day1/input")
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(bytes.NewBuffer(content))
	var total int64
	for s.Scan() {
		i, err := strconv.ParseInt(s.Text(), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		total = total + fuel(i)
	}
	fmt.Println(total)
}
