package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type rate struct {
	one  int
	zero int
}

func rf() [][]int {
	f, err := os.ReadFile("input.txt")
	check(err)
	s := bytes.Split(bytes.Trim(f, "\n\n"), []byte("\n"))
	var l [][]int
	for _, c := range s {
		s := bytes.Split(c, []byte(""))
		var r []int
		for _, k := range s {
			b, _ := strconv.Atoi(string(k))
			r = append(r, b)
		}
		l = append(l, r)
		// fmt.Printf("%d\n", r)
	}
	return l
}

func main() {
	m := rf()
	var gamma, epsilon []string
	for r := 0; r < len(m[0]); r++ {
		var rates rate
		for c := 0; c < len(m); c++ {
			if m[c][r] == 0 {
				rates.zero++
			} else {
				rates.one++
			}
		}
		if rates.one > rates.zero {
			gamma = append(gamma, "1")
			epsilon = append(epsilon, "0")
		} else {
			gamma = append(gamma, "0")
			epsilon = append(epsilon, "1")
		}
	}
	e := strings.Join(epsilon, "")
	g := strings.Join(gamma, "")
	ga, _ := strconv.ParseInt(e, 2, 64)
	ep, _ := strconv.ParseInt(g, 2, 64)
	fmt.Println(ga * ep)
}
