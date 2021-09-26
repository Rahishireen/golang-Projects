package users

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
	"vaccinationdrive/utills/age"
	"vaccinationdrive/utills/date"
	apierrors "vaccinationdrive/utills/error"
)

type UserRegistration struct {
	Name          string `json:"name"`
	DOB           string `json:"dob"`
	AadhaarNumber string  `json:"aadharnumber"`
	PhoneNumber   string  `json:"phonenumber"`
	BeneficiaryID string  `json:"beneficiaryid"`
	DateCreated    string `json:"datecreated"`
}


func (u *UserRegistration) CalculateAge() int {

	day,month,year:=splitDate(u.DOB)
	dob:=getDOB(year,month,day)
	return age.Age(dob)

}

func getDOB(year,month,day int) time.Time{
	dob:=time.Date(year,time.Month(month),day,0,0,0,0,time.UTC)
	return dob
}

func splitDate(date string)(int,int,int){
	

	datestr:=strings.Split(date,"-")

	day,_:=strconv.Atoi(datestr[0])
	month,_:=strconv.Atoi(datestr[1])
	year,_:=strconv.Atoi(datestr[2])

	return day,month,year

	
}

func FieldValidation(length int,RegEx string,field string,fieldName string,Pattern string) *apierrors.RestErr{
	fmt.Println(len(field))
	if len(field)!=length{
		return apierrors.NewBadRequestError(fmt.Sprintf("Invalid %s - Length should be %d",fieldName,length))

	}
	inputPattern:=regexp.MustCompile(RegEx)
	if !inputPattern.MatchString(field){
		return apierrors.NewBadRequestError(fmt.Sprintf("Invalid %s - Format should be %s",fieldName,Pattern))

	}
	return nil
}

func (u *UserRegistration) AgeValidation() *apierrors.RestErr{
	age:=u.CalculateAge()
	if age<=45{
		return apierrors.NewBadRequestError(fmt.Sprintf("Member (Age - %d) not eligible - Age Should be greater than 45",age))

	}
	return nil
}

func CreateBeneficiaryId() string{
	var beneficiaryid string
	
	currenttime:=date.GetTimeString()

	str1 := strings.ReplaceAll(currenttime,"-","")
	str2 := strings.ReplaceAll(str1," ","")
	beneficiaryid= strings.ReplaceAll(str2,":","")

	fmt.Println(beneficiaryid)
	return beneficiaryid
}

