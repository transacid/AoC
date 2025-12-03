package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readInput("../input.txt")
	r := ring.New(100)
	for i := 0; i < 100; i++ {
		r.Value = i
		r = r.Next()
	}
	var counter int
	r = r.Move(50)
	for _, i := range input {
		value, err := strconv.Atoi(i[1])
		if err != nil {
			panic(err)
		}
		if i[0] == "R" {
			for range value {
				r = r.Move(1)
				if r.Value == 0 {
					counter++
				}
			}
		} else {
			for range value {
				r = r.Move(-1)
				if r.Value == 0 {
					counter++
				}
			}
		}

	}
	fmt.Println(counter)

}

func readInput(filename string) [][]string {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	var list [][]string
	for scanner.Scan() {
		sList := strings.SplitN(scanner.Text(), "", 2)
		list = append(list, sList)
	}
	return list
}
