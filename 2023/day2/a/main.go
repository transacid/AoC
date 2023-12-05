package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Game struct {
	id    int
	red   int
	blue  int
	green int
}

func (game Game) parseGame(counter int, line []byte) int {
	re, _ := regexp.Compile(`Game ([0-9]+):(.+)`)
	games := re.FindSubmatch(line)
	game.id, _ = strconv.Atoi(string(games[1]))
	sets := bytes.Split(games[2], []byte(";"))
	fmt.Println(game.id)
	for _, match := range sets {
		re, _ := regexp.Compile(` ([0-9]+ [a-z]+),?`)
		cubes := re.FindAllSubmatch(match, -1)
		// Outerloop:
		for _, i := range cubes {
			for _, j := range i[1:] {
				bla := bytes.Split(j, []byte(" "))
				eyes, _ := strconv.Atoi(string(bla[0]))
				color := string(bla[1])
				switch color {
				case "red":
					if eyes > 12 {
						fmt.Println(color, eyes, game.id)
						// break Outerloop
						return counter
					} else {
						game.red += eyes
					}
				case "blue":
					if eyes > 14 {
						fmt.Println(color, eyes, game.id)
						return counter
						// break Outerloop
					} else {
						game.blue += eyes
					}
				case "green":
					if eyes > 13 {
						fmt.Println(color, eyes, game.id)
						return counter
						// break Outerloop
					} else {
						game.green += eyes
					}
				}
			}
		}
	}
	counter += game.id
	return counter
}

func main() {
	input := readFile("../input.txt")
	si := bytes.Split(input, []byte("\n"))
	var counter int
	var game Game
	for _, l := range si {
		counter = game.parseGame(counter, l)
	}
	fmt.Println(counter)

}

func readFile(filename string) []byte {
	res, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return res
}
