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
	for _, match := range sets {
		re, _ := regexp.Compile(` ([0-9]+ [a-z]+),?`)
		cubes := re.FindAllSubmatch(match, -1)
		for _, i := range cubes {
			for _, j := range i[1:] {
				bla := bytes.Split(j, []byte(" "))
				eyes, _ := strconv.Atoi(string(bla[0]))
				color := string(bla[1])
				switch color {
				case "red":
					if eyes > game.red {
						game.red = eyes
					}
				case "green":
					if eyes > game.green {
						game.green = eyes
					}
				case "blue":
					if eyes > game.blue {
						game.blue = eyes
					}
				}
			}
		}
	}
	counter += game.red * game.green * game.blue
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
