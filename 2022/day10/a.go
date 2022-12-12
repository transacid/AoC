package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type register struct {
	cur   int
	cycle int
	res   []int
}

func main() {
	c := gfi()
	a := strings.Split(string(c), "\n")
	X := register{cur: 1}
	for _, l := range a {
		X.exe(l)
	}
	var f int
	for _, i := range X.res {
		f += i
	}
	fmt.Println(f)
}

func gfi() []byte {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func (X *register) exe(ins string) {
	if strings.HasPrefix(ins, "noop") {
		X.cycle++
		X.eval()
	} else {
		X.cycle++
		X.eval()
		X.cycle++
		X.eval()
		re := regexp.MustCompile(`addx (-?\d+)`)
		matches := re.FindStringSubmatch(ins)
		addx, _ := strconv.Atoi(matches[1])
		X.cur += addx
	}
}

func (X *register) eval() {
	if X.cycle == 20 || X.cycle%40 == 20 {
		X.res = append(X.res, X.cur*X.cycle)
	}
}
