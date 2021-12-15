package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type poly struct {
	template string
	pairs    map[string]string
}

func getInputsByLine() poly {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("could not find file")
		os.Exit(1)
	}

	defer inputFile.Close()

	var polymers poly

	scanner := bufio.NewScanner(inputFile)
	scanner.Scan()
	polymers.template = scanner.Text()
	scanner.Scan()
	polymers.pairs = make(map[string]string)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " -> ")
		polymers.pairs[s[0]] = s[1]
	}

	return polymers
}

func main() {
	p := getInputsByLine()

	alphabet := make(map[string]int)
	for _, c := range p.template {
		alphabet[string(c)]++
	}

	for i := 0; i < 10; i++ {
		var l []string
		for i := range p.template {
			char := p.template[i]
			l = append(l, string(char))
			if i+2 <= len(p.template) {
				pairs := p.pairs[p.template[i:i+2]]
				alphabet[string(pairs)]++
				l = append(l, string(pairs))
			}
		}
		p.template = strings.Join(l, "")
	}

	keys := make([]string, 0, len(alphabet))
	for key := range alphabet {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool { return alphabet[keys[i]] > alphabet[keys[j]] })
	
	fmt.Printf("%d\n", alphabet[keys[0]] - alphabet[keys[len(keys)-1]])
}
