package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func rf() []int {
	f, err := os.ReadFile("input.txt")
	check(err)
	s := bytes.Split(bytes.TrimRight(f, "\n\n"), []byte("\n"))
	var n []int
	for i := range s {
		s := string(s[i])
		number, _ := strconv.Atoi(s)
		n = append(n, number)
	}
	return n
}

func main() {
	s := rf()
	var decrease int
	last := s[0]
	for _, measure := range s {
		if measure > last {
			decrease++
		}
		last = measure
	}
	fmt.Println(decrease)
}
