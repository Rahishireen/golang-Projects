package main

import (
	"bufio"
	"collegedatabase/doublylinkedlist"
	"collegedatabase/staffdatabase"
	"collegedatabase/studentdatabase"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//InsertOption to choose user's Preferable way to insert a node to the list
func insertchoice(studentlist *(doublylinkedlist.List), stafflist *(doublylinkedlist.List)) {

	var number, choice, data int
	var filename string
	fmt.Println("Enter insert Choice")
	fmt.Println("1.get details from user")
	fmt.Println("2.get details from file")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		fmt.Println("Choose")
		fmt.Println("1.studentdata")
		fmt.Println("2.staffdata")
		fmt.Scanln(&data)
		if data == 1 {
			fmt.Println("Enter no of Records")
			fmt.Scanln(&number)
			dataArr := studentdatabase.GetDataStudent(number)
			studentlist.InsertBack(dataArr)
		} else if data == 2 {
			fmt.Println("Enter no of Records")
			fmt.Scanln(&number)
			dataArr := staffdatabase.GetDataStaff(number)
			stafflist.InsertBack(dataArr)
		} else {
			fmt.Println("enter valid Choice")
		}

	case 2:
		fmt.Println("Enter file name")
		fmt.Scanln(&filename)
		studentdataarr, staffdataarr := readCSV(filename)
		studentlist.InsertBack(studentdataarr)
		stafflist.InsertBack(staffdataarr)
	}
}

//DeleteOption to get database option from user
func deletechoice(studentlist *(doublylinkedlist.List), stafflist *(doublylinkedlist.List)) {

	var number, database int
	fmt.Println("Enter the database to deal with")
	fmt.Println("1.Student Database")
	fmt.Println("2.Staff database")
	fmt.Scanln(&database)
	switch database {
	case 1:
		fmt.Println("Enter no of elements")
		fmt.Scanln(&number)
		studentdataArr := studentdatabase.GetDataStudent(number)
		studentlist.Delete(studentdataArr)

	case 2:
		fmt.Println("Enter no of elements")
		fmt.Scanln(&number)
		staffdataArr := staffdatabase.GetDataStaff(number)
		stafflist.Delete(staffdataArr)
	default:
		fmt.Println("enter valid value")
	}
}

//ReadCSV to input from file and returns slice of interface
func readCSV(filename string) ([]interface{}, []interface{}) {
	var studentdataarr, staffdataarr []interface{}
	student := studentdatabase.StudentData{}
	staff := staffdatabase.StaffData{}

	file, ferr := os.Open(filename)
	if ferr != nil {
		fmt.Println("Cannot open")
		fmt.Println(ferr)
		panic(ferr)
	}
	fmt.Println("file open sucess")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, ",")
		key, _ := strconv.Atoi(items[0])
		if key == 1 {
			student.RNo, _ = strconv.Atoi(items[1])
			student.Name = items[2]
			studentdataarr = append(studentdataarr, student)
		}
		if key == 2 {
			staff.EmployeeID, _ = strconv.Atoi(items[1])
			staff.Name = items[2]
			staff.Designation = items[3]
			staffdataarr = append(staffdataarr, staff)
		}
	}
	return studentdataarr, staffdataarr
}

func displaychoice(studentlist *(doublylinkedlist.List), stafflist *(doublylinkedlist.List)) {
	var choice int
	fmt.Println("Enter your choice")
	fmt.Println("1.Student Database")
	fmt.Println("2.Staff Database")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		doublylinkedlist.DisplayOption(studentlist)
	case 2:
		doublylinkedlist.DisplayOption(stafflist)
	default:
		fmt.Println("enter valid Choice")
	}
}

func lengthchoice(studentlist *(doublylinkedlist.List), stafflist *(doublylinkedlist.List)) {
	var choice int
	fmt.Println("Enter your choice")
	fmt.Println("1.Student Database")
	fmt.Println("2.Staff Database")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		fmt.Println("Student Strength is ", studentlist.Length())
	case 2:
		fmt.Println("Staff Strength is ", stafflist.Length())
	default:
		fmt.Println("enter valid Choice")
	}
}
