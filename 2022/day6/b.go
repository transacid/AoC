package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	c := gfi()
	a := strings.SplitAfter(string(c), "")
	var move []string
	for i, char := range a {
		if i <= 13 {
			move = append(move, char)
		} else {
			if !uniq(move) {
				move[0] = move[1]
				move[1] = move[2]
				move[2] = move[3]
				move[3] = move[4]
				move[4] = move[5]
				move[5] = move[6]
				move[6] = move[7]
				move[7] = move[8]
				move[8] = move[9]
				move[9] = move[10]
				move[10] = move[11]
				move[11] = move[12]
				move[12] = move[13]
				move[13] = char
			} else {
				fmt.Println(i)
				break
			}
		}
	}
}

func gfi() []byte {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func uniq(vals []string) bool {

	uvals := []string{}
	seen := make(map[string]bool)

	for _, val := range vals {

		if _, in := seen[val]; !in {

			seen[val] = true
			uvals = append(uvals, val)
		}
	}
	if len(vals) != len(uvals) {
		return false
	} else {
		return true
	}
}
