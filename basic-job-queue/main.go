package main

import (
	"fmt"
	"sync"
)

type jobs struct {
	work []int
	currJob int
}
var wg sync.WaitGroup
func main() {
	
	work := []int{1,2,3,4,5,6,7,8,9}
	var j jobs
	j.work = work
	j.currJob = -1

	for _, val := range work {
		val++
		wg.Add(1)
		go j.pickAndProcess()
	}
	
	wg.Wait()
	
}

func (j *jobs) pickAndProcess() {
	if j.currJob < len(j.work) {
		j.currJob++
	}
	
	val := j.work[j.currJob]
	fmt.Println("Hey, I got -> ", val)
	j.work[j.currJob] = val * val
	wg.Done()
}
