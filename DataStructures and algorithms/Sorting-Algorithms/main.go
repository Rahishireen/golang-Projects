package main

import (
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	var input []int
	n := 10
	for i := 0; i < n; i++ {
		input= append(input, r.Intn(n))
	}
	print(input)
	input = mergesort.MergeSort(input)
}