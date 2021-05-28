package main

import (
	"fmt"
	"math/rand"
	"mergesort/mergesort"
	"time"
)

func main() {

	input:=slice_RandNUmbers()

	fmt.Println("Before Sorting")	
	fmt.Println(input)
	input = mergesort.Mergesort(input)
	fmt.Println("After Sorting")
	fmt.Println(input)
}

func slice_RandNUmbers()[]int{
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	var input []int
	n := 20
	for i := 0; i < n; i++ {
		input= append(input, r.Intn(n))
	}
	return input
}