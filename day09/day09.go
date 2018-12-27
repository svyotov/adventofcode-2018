package day09

import (
	"container/ring"
	"fmt"
)

type player struct {
	id    int
	score int
}

func createMarbles(Nmarbles int) *ring.Ring {
	marbles := ring.New(Nmarbles)
	for i := 0; i < marbles.Len(); i++ {
		marbles.Value = i
		marbles = marbles.Next()
	}
	return marbles.Prev()
}

func createPlayers(Nplayers int) *ring.Ring {
	players := ring.New(Nplayers)
	for i := 0; i < players.Len(); i++ {
		players.Value = player{id: i}
		players = players.Next()
	}
	return players.Prev()
}

func findWinner(players *ring.Ring) player {
	maxScore := 0
	var winner player
	for i := 0; i < players.Len(); i++ {
		score := players.Value.(player).score
		if score > maxScore {
			maxScore = score
			winner = players.Value.(player)
		}
		players = players.Next()

	}
	return winner
}

// MarbleGame plays a game of marbles with N players and N marbles
// returns the score of the game
func MarbleGame(Nplayers, LastMarble int) int {
	skip := 23
	nextValue := 23
	players := createPlayers(Nplayers)
	marbles := createMarbles(LastMarble + 1)

	circle := marbles.Unlink(1)

	for i := 1; i <= LastMarble; i++ {
		if i == nextValue {
			p := players.Value.(player)
			p.score += i
			nextValue += skip
			//
			marbles.Unlink(1)
			//
			circle = circle.Move(-8)
			p.score += circle.Unlink(1).Value.(int)
			circle = circle.Move(1)
			players.Value = p
		} else {
			circle = circle.Move(1)
			circle.Link(marbles.Unlink(1))
			circle = circle.Move(1)
		}
		players = players.Move(1)
	}

	return findWinner(players).score
}

type circle struct {
	position int
}

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

// Run0 task A for current day
func Run0(players, marbles int) int {
	score := MarbleGame(players, marbles)
	fmt.Printf("Day 09 t1: '%v'\n", score)
	return score
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
