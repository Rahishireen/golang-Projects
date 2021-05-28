package doublylinkedlist

import "fmt"

//InsertBack to add the nodes at the end
func (List *List) InsertBack(dataArr []interface{}) {
	for _, data := range dataArr {
		newDataNode := NewNode(data)
		if List.Head != nil {
			TraversingPtr := List.Tail
			TraversingPtr.Next = newDataNode
			newDataNode.Prev = List.Tail
			List.Tail = newDataNode

		} else {
			List.Head = newDataNode
			List.Tail = newDataNode
		}
	}

}

//InsertFront to add the nodes at the end
func (List *List) InsertFront(dataArr []interface{}) {
	for _, data := range dataArr {
		newDataNode := NewNode(data)
		if List.Head != nil {
			TraversingPtr := List.Head
			TraversingPtr.Prev = newDataNode
			newDataNode.Next = List.Head
			List.Head = newDataNode

		} else {
			List.Head = newDataNode
			List.Tail = newDataNode
		}
	}

}

//InsertAfter to insert a node after the given node
func (List *List) InsertAfter(givendata interface{}, newdataArr []interface{}) {
	givenDataNode := List.Search(givendata)

	if givenDataNode != nil {

		for _, newdata := range newdataArr {
			newDataNode := NewNode(newdata)
			if givenDataNode.Next != nil {
				newDataNode.Next = givenDataNode.Next
				newDataNode.Prev = givenDataNode
				givenDataNode.Next = newDataNode
				givenDataNode.Next.Prev = newDataNode

			} else {
				newDataNode.Next = givenDataNode.Next
				newDataNode.Prev = givenDataNode
				givenDataNode.Next = newDataNode

			}
			givenDataNode = givenDataNode.Next

		}
	} else {
		fmt.Println("Given node is not found")
	}

}

//InsertBefore to insert a node Before the given node
func (List *List) InsertBefore(givendata interface{}, newdataArr []interface{}) {
	givenDataNode := List.Search(givendata)
	if givenDataNode != nil {

		for _, newdata := range newdataArr {
			newDataNode := NewNode(newdata)
			if givenDataNode.Prev != nil {
				newDataNode.Prev = givenDataNode.Prev
				newDataNode.Next = givenDataNode
				givenDataNode.Prev.Next = newDataNode
				givenDataNode.Prev = newDataNode

			} else {
				newDataNode.Prev = givenDataNode.Prev
				newDataNode.Next = givenDataNode
				givenDataNode.Prev = newDataNode

			}
			givenDataNode = givenDataNode.Prev

		}
	} else {
		fmt.Println("Given node is not found")
	}

}
