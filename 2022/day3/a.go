package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	c := gfi()
	a := bytes.Split(bytes.TrimRight(c, "\n"), []byte("\n"))
	var sum int
	for _, i := range a {
		first := strings.Split(string(i[:len(i)/2]), "")
		second := strings.Split(string(i[len(i)/2:]), "")
		is := intersection(first, second)
		num := l2i(is[0])
		sum += num
		// fmt.Printf("%s %s\t%s\n", first, second, is)
	}
	fmt.Printf("%d\n", sum)

}

func gfi() []byte {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func intersection(a, b []string) (c []string) {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

func l2i(l string) int {
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	var ind int
	for i, vs := range alphabet {
		if l == vs {
			ind = i + 1
		}
	}
	return ind
}
