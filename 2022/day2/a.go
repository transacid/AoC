package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

var mapping = map[string]string{"A": "r", "X": "r", "B": "p", "Y": "p", "C": "s", "Z": "s"}

func main() {
	c := gfi()
	a := bytes.Split(bytes.TrimRight(c, "\n"), []byte("\n"))
	var res int
	for _, games := range a {
		game := strings.Split(string(games), " ")
		res += win(game[0], game[1])
	}
	fmt.Println(res)
}

func gfi() []byte {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func win(o, m string) int {
	op := mapping[o]
	me := mapping[m]
	var res int
	switch mapping[m] {
	case "r":
		res += 1
	case "p":
		res += 2
	case "s":
		res += 3
	}
	if op == me {
		res += 3
	} else if (op == "r" && me == "p") || (op == "p" && me == "s") || (op == "s" && me == "r") {
		res += 6
	}
	return res
}
