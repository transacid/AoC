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
	var increase, tmp int
	last := s[0] + s[1] + s[2]

	for k, _ := range s {
		if k+2 < len(s) {
			tmp = s[k] + s[k+1] + s[k+2]
			if tmp > last {
				increase++
			}
			last = tmp
		}
	}
	fmt.Println(increase)
}
