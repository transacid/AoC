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
	var res int
	switch m {
	case "X":
		if op == "r" {
			res += 3
		} else if op == "p" {
			res += 1
		} else {
			res += 2
		}
	case "Y":
		res += 3
		if op == "r" {
			res += 1
		} else if op == "p" {
			res += 2
		} else {
			res += 3
		}
	case "Z":
		res += 6
		if op == "r" {
			res += 2
		} else if op == "p" {
			res += 3
		} else {
			res += 1
		}
	}
	return res
}
