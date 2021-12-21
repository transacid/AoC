package main

import (
	"container/ring"
	"fmt"
)

type player struct {
	pos   *ring.Ring
	score int
}

func main() {
	p1_init := 5
	p2_init := 10
	r1 := ring.New(10)
	n1 := r1.Len() + 1
	for i := 1; i < n1; i++ {
		r1.Value = i
		r1 = r1.Next()
	}

	r2 := ring.New(10)
	n2 := r1.Len() + 1
	for i := 1; i < n2; i++ {
		r2.Value = i
		r2 = r2.Next()
	}

	p1 := player{r1, 0}
	p2 := player{r2, 0}
	r1 = r1.Move(p1_init - 1)
	r2 = r2.Move(p2_init - 1)
	ita := 0
	i := 1
	for {
		r1 = r1.Move(i)
		i++
		ita++
		r1 = r1.Move(i)
		i++
		ita++
		r1 = r1.Move(i)
		i++
		ita++
		p1.score += r1.Value.(int)
		if p1.score >= 1000 {
			fmt.Printf("Rolls: %d\n", ita)
			fmt.Printf("P1 score: %d\nP2: %d\n", p1.score, p2.score)
			fmt.Printf("Solve: %d\n", ita*p2.score)
			break
		}

		r2 = r2.Move(i)
		i++
		ita++
		r2 = r2.Move(i)
		i++
		ita++
		r2 = r2.Move(i)
		i++
		ita++
		p2.score += r2.Value.(int)
		if p2.score >= 1000 {
			fmt.Printf("Rolls: %d\n", ita)
			fmt.Printf("P1 score: %d\nP2: %d\n", p1.score, p2.score)
			fmt.Printf("Solve: %d\n", ita*p1.score)
			break
		}

		if i == 100 {
			i = 1
		}
	}
}
