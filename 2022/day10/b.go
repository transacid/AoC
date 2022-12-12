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
	cur    int
	cycle  int
	screen []string
	row    []string
	rowN   int
}

func main() {
	c := gfi()
	a := strings.Split(string(c), "\n")
	X := register{cur: 1}
	for _, l := range a {
		X.exe(l)
	}
	for _, i := range X.screen {
		fmt.Printf("%s", i)
	}

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
	if X.cur-1 == X.rowN || X.cur == X.rowN || X.cur+1 == X.rowN {
		X.row = append(X.row, "#")
		X.rowN++
	} else {
		X.row = append(X.row, " ")
		X.rowN++
	}
	if X.cycle%40 == 0 {
		X.row = append(X.row, "\n")
		r := strings.Join(X.row, "")
		X.screen = append(X.screen, r)
		X.row = []string{}
		X.rowN = 0
	}
}
