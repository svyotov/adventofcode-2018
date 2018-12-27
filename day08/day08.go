package day08

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type tree struct {
	id     int
	trees  []*tree
	values []int
}

func readInputData(file string, execTime int) (data tree, err error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return data, err
	}
	lines := strings.Split(strings.Trim(strings.Split(string(b), "\n")[0], " "), " ")

	// 2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2
	intArray := make([]int, len(lines))
	for i := range lines {
		intArray[i], err = strconv.Atoi(lines[i])
		if err != nil {
			panic(err)
		}
	}

	id := -1
	data, remainded := arrayToTree(intArray, &id)
	if len(remainded) > 0 {
		panic(fmt.Sprintf("failed to process all data %v", remainded))
	}

	return data, err
}

func arrayToTree(data []int, id *int) (tree, []int) {
	(*id)++
	thisTree := tree{id: (*id)}
	nChildren := data[0]
	nValues := data[1]
	data = data[2:]
	for cid := 0; cid < nChildren; cid++ {
		var child tree
		child, data = arrayToTree(data, id)
		thisTree.trees = append(thisTree.trees, &child)
	}
	thisTree.values = data[:nValues]
	return thisTree, data[nValues:]
}

func sumLeafNodes(data tree) int {
	tsum := sum(data.values)
	for _, t := range data.trees {
		tsum += sumLeafNodes(*t)
	}
	return tsum
}

func sum(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func sumLeafNodesMetadata(data tree) int {
	Nchilds := len(data.trees)
	if Nchilds == 0 {
		return sum(data.values)
	}
	tsum := 0
	for _, indx := range data.values {
		indx--
		if indx < Nchilds {
			tsum += sumLeafNodesMetadata(*(data.trees[indx]))
		}
	}
	return tsum
}

// Run1 task A for current day
func Run1(file string) int {
	data, err := readInputData(file, 0)
	if err != nil {
		panic(err)
	}
	score := sumLeafNodes(data)
	fmt.Printf("Day 08 t1: '%v'\n", score)
	return score
}

// Run2 task B for current day
func Run2(file string, workers, execTime int) int {
	data, err := readInputData(file, execTime)
	if err != nil {
		panic(err)
	}
	score := sumLeafNodesMetadata(data)
	fmt.Printf("Day 08 t2: '%v'\n", score)
	return score
}
