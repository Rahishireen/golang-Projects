package doublylinkedlist

import "fmt"

//DisplayOption to get user Preference
func DisplayOption(List *List) {
	var choice int
	fmt.Println("Enter your choice")
	fmt.Println("1.Display Data")
	fmt.Println("2.Display Data with Node details")
	fmt.Println("3.Display Data in Reverse Order")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		List.DisplayData()
	case 2:
		List.DisplayNode()
	case 3:
		List.DisplayDataReverse()
	default:
		fmt.Println("enter Valid Choice")
	}

}

//DisplayData to display the data
func (List *List) DisplayData() {
	len := List.Length()
	if len == 0 {
		fmt.Println("Empty List")
	} else {
		traversingPtr := List.Head
		for traversingPtr.Next != nil {
			fmt.Printf("%v -->", traversingPtr.Data)
			traversingPtr = traversingPtr.Next
		}
		fmt.Printf("%v\n", traversingPtr.Data)
	}
}

//DisplayNode to display the data with node details
func (List *List) DisplayNode() {
	len := List.Length()
	if len == 0 {
		fmt.Println("Empty List")
	} else {
		traversingPtr := List.Head
		for traversingPtr.Next != nil {
			fmt.Printf("%v-->", traversingPtr.Prev)
			fmt.Printf("%v -->", traversingPtr.Data)
			fmt.Printf("%v===>", traversingPtr.Next)
			traversingPtr = traversingPtr.Next
		}
		fmt.Printf("%v-->", traversingPtr.Prev)
		fmt.Printf("%v -->", traversingPtr.Data)
		fmt.Printf("%v\n", traversingPtr.Next)
	}
}

//DisplayDataReverse to display the data
func (List *List) DisplayDataReverse() {
	len := List.Length()
	if len == 0 {
		fmt.Println("Empty List")
	} else {
		traversingPtr := List.Tail
		for traversingPtr.Prev != nil {
			fmt.Printf("%v -->", traversingPtr.Data)
			traversingPtr = traversingPtr.Prev
		}
		fmt.Printf("%v\n", traversingPtr.Data)
	}
}
