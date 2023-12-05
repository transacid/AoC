package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input := readFile("input.txt")
	si := bytes.Split(input, []byte("\n"))
	var coords int
	for _, line := range si {
		r, _ := regexp.Compile(`([0-9])`)
		res := r.FindAll(line, -1)
		first := string(res[0])
		second := string(res[len(res)-1])
		coord, _ := strconv.Atoi(fmt.Sprintf("%s%s", first, second))
		coords += coord
	}
	fmt.Println(coords)
}

func readFile(filename string) []byte {
	res, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return res
}
