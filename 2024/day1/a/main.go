package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
)

func main() {
	lists := readfile("../input.txt")
	l1 := lists[0]
	l2 := lists[1]
	var result int
	for i, j := range l1 {
		rs := j - l2[i]
		r := math.Abs(float64(rs))
		result += int(r)
	}
	fmt.Println(result)

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
