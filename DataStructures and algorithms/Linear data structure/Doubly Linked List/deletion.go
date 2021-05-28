package doublylinkedlist

import "fmt"

//Delete to delete a given node
func (List *List) Delete(dataArr []interface{}) {
	for _, data := range dataArr {
		givenDataNode := List.Search(data)

		if givenDataNode != nil {

			if givenDataNode.Prev == nil {
				if givenDataNode.Next == nil {
					List.Head = nil
					List.Tail = nil

				} else {
					givenDataNode.Next.Prev = nil
					List.Head = givenDataNode.Next

				}
			} else if givenDataNode.Next == nil {
				if givenDataNode.Prev == nil {
					List.Head = nil
					List.Tail = nil
				} else {

					givenDataNode.Prev.Next = nil
					List.Tail = givenDataNode.Prev
				}
			} else {
				givenDataNode.Prev.Next = givenDataNode.Next
				givenDataNode.Next.Prev = givenDataNode.Prev
			}
		} else {
			fmt.Println("Given node is not found")
		}
	}

}
