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

func input() []int {
	f, err := os.ReadFile("input.txt")
	check(err)
	s := bytes.TrimRight(f, "\n\n")
	k := bytes.Split(s, []byte(","))
	var n []int
	for i := range k {
		s := string(k[i])
		number, _ := strconv.Atoi(s)
		n = append(n, number)
	}
	return n
}

func main() {
	s := input()
	arr := make([]int, 9)
	for _, fish := range s {
		arr[fish]++
	}
	for i := 0; i < 256; i++ {
		temp0 := arr[0]
		arr[0] = arr[1]
		arr[1] = arr[2]
		arr[2] = arr[3]
		arr[3] = arr[4]
		arr[4] = arr[5]
		arr[5] = arr[6]
		arr[6] = arr[7] + temp0
		arr[7] = arr[8]
		arr[8] = temp0

	}
	fmt.Printf("%v\n", arr[0]+arr[1]+arr[2]+arr[3]+arr[4]+arr[5]+arr[6]+arr[7]+arr[8])
}
