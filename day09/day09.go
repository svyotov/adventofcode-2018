package day09

import (
	"container/ring"
	"fmt"
)

// MarbleGameV2 plays a game of marbles with N players and N marbles
// returns the score of the game faster vestion of MarbleGame
func MarbleGameV2(Nplayers, LastMarble int) int {
	skip := 23
	nextValue := 23
	players := make([]int, Nplayers)
	circle := ring.New(1)
	circle.Value = 0
	pid := 0

	for i := 1; i <= LastMarble; i++ {
		if i == nextValue {
			nextValue += skip
			circle = circle.Move(-8)
			players[pid] = players[pid] + i + circle.Unlink(1).Value.(int)
			circle = circle.Move(1)
		} else {
			circle = circle.Move(1)
			circle.Link(&ring.Ring{Value: i})
			circle = circle.Move(1)
		}
		pid = (pid + 1) % Nplayers
	}
	return max(players)
}

func max(s []int) int {
	m := s[0]
	for i := range s {
		if m < s[i] {
			m = s[i]
		}
	}
	return m
}

// Run1 task A for current day
func Run1(players, marbles int) int {
	score := MarbleGameV2(players, marbles)
	fmt.Printf("Day 09 t1: '%v'\n", score)
	return score
}

// Run2 task B for current day
func Run2(players, marbles int) int {
	score := MarbleGameV2(players, marbles*100)
	fmt.Printf("Day 09 t2: '%v'\n", score)
	return score
}
