package services

import (
	users "vaccinationdrive/Domain/users"
	"vaccinationdrive/utills/date"
	apierrors "vaccinationdrive/utills/error"
)

func CreateUser(user users.UserRegistration) (*users.UserRegistration,*apierrors.RestErr) {

	if dateerr := users.FieldValidation(10,`\d{2}-\d{2}-\d{4}`,user.DOB,"Date","DD-MM-YYYY"); dateerr != nil {
		return nil, dateerr
	}

	if aadharerr:=users.FieldValidation(12,`\d{12}`,user.AadhaarNumber,"Aadhaar Number","12 Digit");aadharerr!=nil{
		return nil,aadharerr
	}
	if phoneNumErr:=users.FieldValidation(10,`\d{10}`,user.PhoneNumber,"Phone Number","10 Digit");phoneNumErr!=nil{
		return nil,phoneNumErr
	}

	if ageErr:= user.AgeValidation();ageErr!=nil{
		return nil,ageErr
	}

	user.BeneficiaryID=users.CreateBeneficiaryId()
	user.DateCreated=date.GetTimeString()

	if DBerr := user.Save();DBerr!=nil{
		return nil,DBerr
	}
	
	return &user, nil
}