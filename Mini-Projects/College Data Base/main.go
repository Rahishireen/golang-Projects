package main

import (
	"collegedatabase/doublylinkedlist"
	"fmt"
)

func main() {

	studentlist := doublylinkedlist.List{}
	stafflist := doublylinkedlist.List{}
	var choice int
	for {
		fmt.Println("Enter your Choice")
		fmt.Println("1.Insert")
		fmt.Println("2.Delete")
		fmt.Println("3.Display")
		fmt.Println("4.Strength")
		fmt.Println("5.End")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			insertchoice(&studentlist, &stafflist)
		case 2:
			deletechoice(&studentlist, &stafflist)
		case 3:
			displaychoice(&studentlist, &stafflist)
		case 4:
			lengthchoice(&studentlist, &stafflist)
		case 5:
			fmt.Println("The End")
		default:
			fmt.Println("enter valid choice")
		}
		if choice == 5 {
			break
		}
	}

}
