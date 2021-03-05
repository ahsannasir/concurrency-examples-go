// Go Channels
// -> chan means reading into channel
// cham <- writing into channel

package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to job distributor, please type finish to get the final output")
	i := 0
	jobs := []int{2,4,6,8,10,12,15,16,18,22,34}
	c := make(chan int)
	// one go routine for every element of the array
	for j:=0;j<len(jobs);j++ {
		go exec(c,jobs)
	}
	// initial write to start work in go routines
	c <- i
	// just wait a while and print out final work
	var halter string
	fmt.Scanf("%s", &halter)
	if halter == "finish" {
		fmt.Println(jobs)
	}
}

func exec(c chan int, jobs []int) {
	// read from channel
	i := <- c
	//work -> square
	jobs[i] = jobs[i] * jobs[i]
	//make sure not write into channel after last element to avoid deadlock
	if i < len(jobs) {
		i = i+1
		c <- i
	}
}