package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	c := gfi()
	a := strings.Split(string(c), "\n")
	stacks := make(map[int][]string)
	var ins [][]string
	for _, line := range a {
		if strings.HasPrefix(line, " 1") || line == "" {
			continue
		} else if strings.HasPrefix(line, "move") {
			ins = append(ins, buildinstruction(line))
		} else {
			buildstack(line, stacks)
		}
	}
	for _, stack := range stacks {
		revSlice(stack)
	}
	for _, instruction := range ins {
		execute(instruction, stacks)
	}
	var res []string
	for i := 1; i <= len(stacks); i++ {
		res = append(res, stacks[i][len(stacks[i])-1])
	}
	fmt.Printf("%s\n", strings.Join(res, ""))
}

func gfi() []byte {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func buildstack(line string, stacks map[int][]string) map[int][]string {
	sl := strings.Split(line, "")
	counter := 1
	for i, c := range sl {
		if i%4 == 1 {
			if c != " " {
				stacks[counter] = append(stacks[counter], c)
			}
			counter++
		}
	}
	return stacks
}

func buildinstruction(line string) []string {
	re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	ins := re.FindStringSubmatch(line)
	return ins[1:]
}

func execute(ins []string, stacks map[int][]string) {
	times := sureAtoi(ins[0])
	from := sureAtoi(ins[1])
	to := sureAtoi(ins[2])

	for i := 0; i < times; i++ {
		move := stacks[from][len(stacks[from])-1]
		stacks[from] = stacks[from][:len(stacks[from])-1]
		stacks[to] = append(stacks[to], move)
	}
}

func revSlice(slice []string) []string {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

func sureAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
