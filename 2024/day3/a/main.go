package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	contents := readInput("input.txt")
	re := regexp.MustCompile(`(?m)(mul\(([0-9]{1,3},[0-9]{1,3})\))`)
	var res int
	for _, l := range contents {
		for _, match := range re.FindAllStringSubmatch(l, -1) {
			r := evaluate(match[len(match)-1])
			res += r
		}
	}
	fmt.Println(res)
}

func readInput(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	var contents []string
	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}
	return contents
}

func evaluate(input string) int {
	si := strings.Split(input, ",")
	d1, _ := strconv.Atoi(si[0])
	d2, _ := strconv.Atoi(si[1])
	return d1 * d2
}
