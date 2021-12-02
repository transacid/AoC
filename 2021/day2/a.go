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

type command struct {
	cmd    string
	amount int
}

func rf() []command {
	f, err := os.ReadFile("input.txt")
	check(err)
	s := bytes.Split(bytes.Trim(f, "\n\n"), []byte("\n"))
	var c []command
	for _, l := range s {
		b := bytes.Split(l, []byte(" "))
		com := string(b[0])
		cou, _ := strconv.Atoi(string(b[1]))
		c = append(c, command{com, cou})
	}
	return c
}

func main() {
	m := rf()
	var hpos, vpos int
	for _, cmd := range m {
		if cmd.cmd == "forward" {
			hpos += cmd.amount
		} else if cmd.cmd == "down" {
			vpos += cmd.amount
		} else if cmd.cmd == "up" {
			vpos -= cmd.amount
		}
	}
	fmt.Println(hpos * vpos)
}
