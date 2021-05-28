package doublylinkedlist

import "fmt"

//Search to search the given data and returns the node details
func (List *List) Search(data interface{}) *Node {

	traversingPtr := List.Head
	len := List.Length()
	fmt.Println(len)
	for i := 0; i < len; i++ {
		if traversingPtr.Data == data {
			return traversingPtr
		}
		traversingPtr = traversingPtr.Next
	}
	return nil

}

//Length to fnd no of elements in the list
func (List *List) Length() int {
	len := 0
	traversingPtr := List.Head

	for traversingPtr != nil {
		len = len + 1
		traversingPtr = traversingPtr.Next
	}
	return len
}
