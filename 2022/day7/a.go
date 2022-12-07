package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	c := gfi()
	a := strings.SplitAfter(string(c), "\n")
	var path string
	files := make(map[string]int)
	dirs := make(map[string]int)
	for _, l := range a {
		if strings.HasPrefix(l, "$") {
			cmd, arg := eval(l)
			if cmd == "ls" {
				continue
			} else {
				path = filepath.Join(path, arg)
			}
		} else if strings.HasPrefix(l, "dir") {
			continue
		} else {
			s := strings.Split(l, " ")
			si, _ := strconv.Atoi(s[0])
			files[filepath.Join(path, strings.Trim(s[1], "\n"))] = si
		}
	}
	for d := range files {
		dir, _ := filepath.Split(d)
		p := strings.SplitAfter(dir, "/")
		var q string
		for _, k := range p[:len(p)-1] {
			q = q + k
			dirs[q] += files[d]
		}
	}
	var win int
	for _, f := range dirs {
		if f <= 100000 {
			win += f
		}
	}
	fmt.Println(win)
}

func gfi() []byte {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func eval(cmd string) (string, string) {
	re := regexp.MustCompile(`^\$ (cd|ls) ?(.+)?`)
	ins := re.FindStringSubmatch(cmd)
	return ins[1], ins[2]
}
