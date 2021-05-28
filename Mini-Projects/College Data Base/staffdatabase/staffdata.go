package staffdatabase

import "fmt"

//StaffData contains emp id,name,designation of staff
type StaffData struct {
	EmployeeID  int
	Name        string
	Designation string
}

//GetDataStaff to get input from user for staff
func GetDataStaff(num int) []interface{} {
	var dataArr []interface{}
	staff := StaffData{}
	for i := 0; i < num; i++ {
		fmt.Println("Enter emp id,Name,Designation")
		fmt.Scanln(&staff.EmployeeID)
		fmt.Scanln(&staff.Name)
		fmt.Scanln(&staff.Designation)
		dataArr = append(dataArr, staff)
	}
	return dataArr
}
