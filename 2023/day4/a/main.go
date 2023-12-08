package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func parseGame(line [][]byte) int {
	var counter int
	for _, l := range line {
		gameSplit := bytes.Split(l, []byte(":"))
		winngingHave := bytes.Split(gameSplit[1], []byte("|"))
		winning := bytes.Split(winngingHave[0], []byte(" "))
		have := bytes.Split(winngingHave[1], []byte(" "))
		var winningInts []int
		var haveInts []int
		var winCount int
		for _, i := range winning {
			winningNummber, err := strconv.Atoi(string(i))
			if err == nil {
				winningInts = append(winningInts, winningNummber)
			}
		}
		for _, j := range have {
			haveNummber, err := strconv.Atoi(string(j))
			if err == nil {
				haveInts = append(haveInts, haveNummber)
			}
		}
		for _, match := range haveInts {
			if slices.Contains(winningInts, match) {
				if winCount == 0 {
					winCount++
				} else {
					winCount = winCount * 2
				}
			}
		}
		counter += winCount
	}
	return counter
}

func main() {
	input := readFile("../input.txt")
	sum := parseGame(input)
	fmt.Println(sum)
}

func readFile(filename string) [][]byte {
	res, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return bytes.Split(res, []byte("\n"))

}
