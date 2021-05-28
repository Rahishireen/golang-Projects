package doublylinkedlist

import "fmt"

//Node structure in DLL
type Node struct {
	Prev *Node
	Data interface{}
	Next *Node
}

//List DLL
type List struct {
	Head *Node
	Tail *Node
}

//NewNode to create a node and returns the node address
func NewNode(data interface{}) *Node {
	NewNodeAdd := &Node{
		Prev: nil,
		Data: data,
		Next: nil,
	}
	return NewNodeAdd
}

//Sample to check sample import
func Sample() {
	fmt.Println("HELLO WORLD")
}
