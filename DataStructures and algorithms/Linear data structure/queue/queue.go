package queue

import (
	"errors"
	"fmt"
)


type Queue []interface{}


// IsEmpty will check for empty stack
func (q *Queue) IsEmpty() bool {
	return (len(*q) == 0)
}

func (q *Queue) Enqueue(element interface{}) {
	*q = append(*q, element)
}

func (q *Queue) Dequeue() (interface{}, error) {

	if q.IsEmpty() {
		return -1, errors.New("Queue is underflow")
	}
	
	element := (*q)[0]
	*q = (*q)[1:]
	return element,nil

}

func (q *Queue)PrintQueue(){
	fmt.Println(*q)
}