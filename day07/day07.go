package day07

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type node struct {
	id       string
	duration int
	previous []string
}

type nodes []node

func (e nodes) Len() int {
	return len(e)
}

func (e nodes) Less(i, j int) bool {
	return e[i].id < e[j].id
}

func (e nodes) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func readInputData(file string, execTime int) (data map[string]node, err error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return data, err
	}
	lines := strings.Split(strings.Trim(string(b), "\n"), "\n")

	data = make(map[string]node, len(lines))

	for _, line := range lines {
		first, second := "", ""
		_, err := fmt.Sscanf(line, "Step %v must be finished before step %v can begin.", &first, &second)
		if err != nil {
			return data, err
		}
		data[second] = node{id: second, duration: duration(second, execTime), previous: append(data[second].previous, first)}
		data[first] = node{id: first, duration: duration(first, execTime), previous: data[first].previous}
	}

	return data, err
}

func duration(id string, execTime int) int {
	return execTime + int(id[0]) - 'A' + 1
}

func sortData(in map[string]node) nodes {
	data := make(nodes, 0, len(in))
	for _, value := range in {
		data = append(data, value)
	}
	sort.Sort(data)
	return data
}

func getExecutionOrder(data *nodes) string {
	executionOrder := ""
	for {
		nextKey := ""
		nextIndex := -1
		for k, v := range *data {
			if len(v.previous) == 0 {
				nextKey = v.id
				nextIndex = k
				break
			}
		}
		if nextIndex == -1 {
			panic(fmt.Sprintf("no next move available %v", data))
		}
		(*data) = remove(*data, nextIndex)
		popFromQueue(data, nextKey)
		executionOrder += nextKey
		if len(*data) == 0 {
			break
		}
	}
	return executionOrder
}

func popFromQueue(data *nodes, id string) {
	for k := range *data {
		(*data)[k].previous = removeString((*data)[k].previous, id)
	}
}

func removeString(data []string, key string) []string {
	for i, v := range data {
		if v == key {
			data = append(data[:i], removeString(data[i+1:], key)...)
			return data
		}
	}
	return data
}

func remove(slice nodes, s int) nodes {
	return append(slice[:s], slice[s+1:]...)
}

type queue struct {
	freeWorkers int
	inProgress  nodes
}

func (e queue) HasFreeWorkers() bool {
	return e.freeWorkers > 0
}

func (e queue) Done() bool {
	return len(e.inProgress) == 0
}

func (e *queue) Tick() []string {
	done := []string{}
	toremove := []int{}
	for k := range e.inProgress {
		e.inProgress[k].duration--
		if e.inProgress[k].duration <= 0 {
			done = append(done, e.inProgress[k].id)
			toremove = append(toremove, k)
		}
	}
	sort.Ints(toremove)
	for idx := len(toremove) - 1; idx >= 0; idx-- {
		e.inProgress = remove(e.inProgress, toremove[idx])
		e.freeWorkers++
	}
	return done
}

func (e *queue) Push(job node) {
	if !e.HasFreeWorkers() {
		panic("trying to add to queue, no workers available")
	}
	e.inProgress = append(e.inProgress, job)
	e.freeWorkers--
}

func getExecutionTimeParallel(data nodes, workers int) int {
	execTime := 0
	inProgress := queue{freeWorkers: workers}
	for l := 0; l < 10000; l++ {
		singleStep(&data, &inProgress)
		// no more data on queue
		if len(data) == 0 && inProgress.Done() {
			break
		}
		execTime++
	}
	return execTime
}

func singleStep(data *nodes, workers *queue) string {
	done := workers.Tick()
	for _, taskID := range done {
		popFromQueue(data, taskID)
	}

	for l := 0; l < 10000; l++ {
		if !workers.HasFreeWorkers() || len(*data) == 0 {
			break
		}
		nextIndex := -1
		for k, v := range *data {
			if len(v.previous) == 0 {
				nextIndex = k
				break
			}
		}
		if nextIndex == -1 {
			break
		}
		workers.Push((*data)[nextIndex])
		*data = remove(*data, nextIndex)
		if len(*data) > 0 && workers.Done() {
			panic(fmt.Sprintf("no next move available %v", data))
		}
	}
	return fmt.Sprintf("%v", done)
}

// Run1 task A for current day
func Run1(file string) string {
	data, err := readInputData(file, 0)
	if err != nil {
		panic(err)
	}
	sdata := sortData(data)
	score := getExecutionOrder(&sdata)
	fmt.Printf("Day 07 t1: '%v'\n", score)
	return score
}

// Run2 task B for current day
func Run2(file string, workers, execTime int) int {
	data, err := readInputData(file, execTime)
	if err != nil {
		panic(err)
	}
	sdata := sortData(data)
	score := getExecutionTimeParallel(sdata, workers)
	fmt.Printf("Day 07 t2: '%v'\n", score)
	return score
}
