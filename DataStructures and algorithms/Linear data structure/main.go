package main

import (
	"fmt"
	"math/rand"
	"queue/queue"
	"time"
)

func main() {
	n := 10

	queue := insert(n)
	(&queue).PrintQueue()
	for i:=0;i<n;i++ {
		element,err:=(&queue).Dequeue()
		if err!=nil{
			fmt.Println(err)
		}else{
			fmt.Println("Dequeued element is",element)
		}

	}
	(&queue).PrintQueue()
}

func insert(n int) queue.Queue {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	var queue queue.Queue
	for i := 0; i < n; i++ {
		element:=r.Intn(n)
		(&queue).Enqueue(element)
		fmt.Println("Enqueued element is",element)
	}
	return queue

}