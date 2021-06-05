package main

import (
	"fmt"
	"math/rand"
	"quicksort/mergesort"
	"quicksort/quicksort"
	"time"
)

func main() {

	input:=slice_RandNUmbers()
	fmt.Println("Before Sorting")	
	fmt.Println(input)
	choice:=2
	switch(choice){
	case 1:
		input=mergesort.Mergesort(input)
	case 2:
		inputlen:=len(input)
	    quicksort.QuickSort(input,0,(inputlen-1))
	default:
		fmt.Println("Enter Valid Choice")

	
	}
	
	fmt.Println("After Sorting")
	fmt.Println(input)
}

func slice_RandNUmbers()[]int{
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	var input []int
	n := 10
	for i := 0; i < n; i++ {
		input= append(input, r.Intn(n))
	}
	return input
}