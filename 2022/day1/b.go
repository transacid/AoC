package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	c := gfi()
	a := bytes.Split(c, []byte("\n"))
	var res []int
	var tmp int
	for _, e := range a {
		cur, _ := strconv.Atoi(string(e))
		tmp += cur
		if cur == 0 {
			res = append(res, tmp)
			tmp = 0
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(res)))
	var p int
	for _, i := range res[:3] {
		p += i
	}
	fmt.Println(p)
}

func gfi() []byte {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return f
}
