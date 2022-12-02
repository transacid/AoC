package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	c := gfi()
	a := bytes.Split(c, []byte("\n"))
	var res int
	var tmp int
	for _, e := range a {
		cur, _ := strconv.Atoi(string(e))
		tmp += cur
		if cur == 0 {
			if tmp > res {
				res = tmp
			}
			tmp = 0
		}
	}
	fmt.Println(res)
}

func gfi() []byte {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return f
}
