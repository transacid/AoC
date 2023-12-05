package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var numbers map[string]string = map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9", "1": "1", "2": "2", "3": "3", "4": "4", "5": "5", "6": "6", "7": "7", "8": "8", "9": "9"}

func main() {
	input := readFile("../input.txt")
	si := bytes.Split(input, []byte("\n"))
	var coords int

	for _, line := range si {
		chars := make(map[int]string)
		for s, i := range numbers {
			first := bytes.Index(line, []byte(s))
			last := bytes.LastIndex(line, []byte(s))
			if first != -1 {
				chars[first] = i
			}

			if last != -1 {
				chars[last] = i
			}
		}

		keys := make([]int, 0, len(chars))
		for k := range chars {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		first := chars[keys[0]]
		last := chars[keys[len(chars)-1]]

		coord, err := strconv.Atoi(fmt.Sprintf("%s%s", first, last))
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s %d\n", string(line), coord)
		coords += coord
	}
	fmt.Println(coords)
}

func readFile(filename string) []byte {
	res, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return res
}
