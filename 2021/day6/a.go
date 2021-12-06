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

func input() []int {
	f, err := os.ReadFile("input.txt")
	check(err)
	s := bytes.TrimRight(f, "\n\n")
	k := bytes.Split(s, []byte(","))
	var n []int
	for i := range k {
		s := string(k[i])
		number, _ := strconv.Atoi(s)
		n = append(n, number)
	}
	return n
}

func main() {
	s := input()
	for i := 0; i < 80; i++ {
		for pos, fish := range s {
			if fish == 0 {
				s[pos] = 6
				s = append(s, 8)
			} else {
				s[pos]--
			}
		}
	}
	fmt.Printf("%d\n", len(s))
}
