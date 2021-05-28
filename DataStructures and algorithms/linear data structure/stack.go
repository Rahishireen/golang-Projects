package stack

import (
	"errors"
	"fmt"
)


type Stack []interface{}


// IsEmpty will check for empty stack
func (s *Stack) IsEmpty() bool {
	return (len(*s) == 0)
}

func (s *Stack) Push(element interface{}) {
	*s = append(*s, element)
}

func (s *Stack) Pop() (interface{}, error) {

	if s.IsEmpty() {
		return -1, errors.New("Stack is underflow")
	}
	index := (len(*s) - 1)
	element := (*s)[index]
	*s = (*s)[:index]
	return element,nil

}

func (s *Stack)PrintStack(){
	fmt.Println(*s)
}