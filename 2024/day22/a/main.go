package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const ModuloNumber = 16777216

func main() {
	buyers := readInput("../input.txt")
	var res int
	for _, buyer := range buyers {
		for i := 0; i < 2000; i++ {
			buyer = operation(buyer)
		}
		res += buyer
	}
	fmt.Println(res)
}

func readInput(filename string) []int {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	var contents []int
	for scanner.Scan() {
		initSecretNumber, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		contents = append(contents, initSecretNumber)
	}
	return contents
}

func operation(secretNumber int) int {
	step := secretNumber * 64
	secretNumber ^= step
	secretNumber %= ModuloNumber
	step = secretNumber / 32
	secretNumber ^= step
	secretNumber %= ModuloNumber
	step = secretNumber * 2048
	secretNumber ^= step
	secretNumber %= ModuloNumber
	return secretNumber
}
