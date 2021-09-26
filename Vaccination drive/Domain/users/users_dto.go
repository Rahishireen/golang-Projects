package users

import (
	"fmt"
	"strings"
	"time"
	"vaccinationdrive/Datasources/mysql/users_db"
	apierrors "vaccinationdrive/utills/error"
)

const (
	QueryInsetUser = "INSERT INTO users(Name,DOB,Aadhaar_Number,Phone_Number,Beneficiary_Id,Date_Created) VALUES (?,?,?,?,?,?);"
	errorNoRows  ="no rows in result set"
	Duplicate="Duplicate"
	QueryGetId = "SELECT Name,DOB,Aadhaar_Number,Phone_Number,Beneficiary_Id FROM users WHERE Beneficiary_Id = ?"
)

func (user *UserRegistration) Save() *apierrors.RestErr {

	stmt,err:=users_db.Client.Prepare(QueryInsetUser)
	if err!=nil{
		return apierrors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	insertResult,err :=stmt.Exec(user.Name,user.DOB,user.AadhaarNumber,user.PhoneNumber,user.BeneficiaryID,time.Now())
	

	if err!=nil{

		if strings.Contains(err.Error(),Duplicate){			
			return apierrors.NewBadRequestError(fmt.Sprintf("Aadhaar NUmber %s already exists",user.AadhaarNumber))
	
		}
		return apierrors.NewInternalServerError(fmt.Sprintf("error when trying to save user:%s",err.Error()))
	}

	_,err= insertResult.LastInsertId()
	if err!=nil{
		return apierrors.NewInternalServerError(fmt.Sprintf("error when trying to save user:%s",err.Error()))
	}

	fmt.Println(insertResult)
	return nil
}

func (user *UserRegistration) Get() *apierrors.RestErr{
	stmt,err:=users_db.Client.Prepare(QueryGetId)
	if err!=nil{
		return apierrors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	//Query Single Row
	result:=stmt.QueryRow(user.BeneficiaryID)
	if err:=result.Scan(&user.Name,&user.DOB,&user.AadhaarNumber,&user.PhoneNumber,&user.BeneficiaryID);err!=nil{
		if strings.Contains(err.Error(),errorNoRows){
			return apierrors.NewNotFoundError(fmt.Sprintf("Beneficiary Id %s not registered",user.BeneficiaryID))

		}
		return apierrors.NewInternalServerError(
			fmt.Sprintf("error When trying to get user with Beneficiary Id %s:%s",user.BeneficiaryID,err.Error()))
		}

	return nil

}