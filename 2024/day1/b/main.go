package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	lists := readfile("../input.txt")
	l1 := lists[0]
	l2 := lists[1]
	var counter int
	for _, li1 := range l1 {
		var icounter int
		for _, li2 := range l2 {
			if li1 == li2 {
				icounter++
			}
		}
		counter += li1 * icounter
		icounter = 0
	}
	fmt.Println(counter)
}

func readfile(filename string) [][]int {
	f, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var a, b []int
	var list [][]int
	lines := bytes.Split(f, []byte("\n"))
	for _, line := range lines {
		sLine := bytes.Split(line, []byte("   "))
		fc := sLine[0]
		sc := sLine[1]
		fci, err := strconv.Atoi(string(fc))
		if err != nil {
			panic(err)
		}
		sci, err := strconv.Atoi(string(sc))
		if err != nil {
			panic(err)
		}
		a = append(a, fci)
		b = append(b, sci)
	}
	slices.Sort(a)
	slices.Sort(b)
	list = append(list, a, b)
	return list
}
