package studentdatabase

import (
	"fmt"
)

//StudentData contains student name and roll_no
type StudentData struct {
	RNo  int
	Name string
}

//GetDataStudent to get input from user
func GetDataStudent(num int) []interface{} {
	var dataArr []interface{}
	student := StudentData{}
	for i := 0; i < num; i++ {
		fmt.Println("Enter Student Name & No")
		fmt.Scanln(&student.Name)
		fmt.Scanln(&student.RNo)
		dataArr = append(dataArr, student)
	}
	return dataArr
}
