package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	c := gfi()
	a := strings.Split(strings.TrimRight(string(c), "\n"), "\n")
	var res int
	for _, p := range a {
		pairs := strings.Split(p, ",")
		pair1 := strings.Split(pairs[0], "-")
		pair2 := strings.Split(pairs[1], "-")
		p11, _ := strconv.Atoi(pair1[0])
		p12, _ := strconv.Atoi(pair1[1])
		p21, _ := strconv.Atoi(pair2[0])
		p22, _ := strconv.Atoi(pair2[1])
		if p12 >= p21 && p22 >= p11 {
			res++
		}
	}
	fmt.Printf("%v\n", res)
}

func gfi() []byte {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return f
}
