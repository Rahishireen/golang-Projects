package appointment

import (
	"fmt"
	"vaccinationdrive/Datasources/mysql/users_db"
	"vaccinationdrive/utills/date"
	apierrors "vaccinationdrive/utills/error"
)

func (Bapp *BookAppointment) CheckBeneficiaryInDBAdd(query string) (int,*apierrors.RestErr) {
	var beneficiary_Entry int
	stmt, err := users_db.Client.Prepare(query)
	if err != nil {
		return 0,apierrors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	SelectResult, err := stmt.Query(Bapp.BeneficiaryID)
	
	if err != nil {
		return 0,apierrors.NewInternalServerError(fmt.Sprintf("error when trying to get Beneficiary Count in DB for Add:%s", err.Error()))
	} else {

	for SelectResult.Next() {
		SelectResult.Scan(&beneficiary_Entry)
		
	}
	fmt.Println("Beneficiary Count",beneficiary_Entry)
	return beneficiary_Entry,nil

}
}

func (Bapp *BookAppointment) CheckBeneficiaryInDBUpdate(query string) (int,*apierrors.RestErr) {
	var beneficiary_Entry int
	stmt, err := users_db.Client.Prepare(query)
	if err != nil {
		return 0,apierrors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	SelectResult, err := stmt.Query(Bapp.BeneficiaryID,Bapp.TimeSlot,Bapp.DoseType,Bapp.UserVC)
	
	if err != nil {
		return 0,apierrors.NewInternalServerError(fmt.Sprintf("error when trying to get Beneficiary Count in DB for Update:%s", err.Error()))
	} else {

	for SelectResult.Next() {
		SelectResult.Scan(&beneficiary_Entry)
		
	}
	fmt.Println("Beneficiary Count",beneficiary_Entry)
	return beneficiary_Entry,nil

}
}
func (Bapp *BookAppointment) CheckBeneficiaryInDBDelete(query string) (int,*apierrors.RestErr) {
	var beneficiary_Entry int
	stmt, err := users_db.Client.Prepare(query)
	if err != nil {
		return 0,apierrors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	SelectResult, err := stmt.Query(Bapp.BeneficiaryID,Bapp.BookingDate,Bapp.TimeSlot,Bapp.DoseType,Bapp.UserVC)
	
	if err != nil {
		return 0,apierrors.NewInternalServerError(fmt.Sprintf("error when trying to get Beneficiary Count in DB for Delete:%s", err.Error()))
	} else {

	for SelectResult.Next() {
		SelectResult.Scan(&beneficiary_Entry)
		
	}
	fmt.Println("Beneficiary Count",beneficiary_Entry)
	return beneficiary_Entry,nil

}
}
func (Bapp *BookAppointment) GetVaccineDetailsInDB(query string) (*apierrors.RestErr) {
	var beneficiary_DB_Data BookAppointment
	stmt, err := users_db.Client.Prepare(query)
	if err != nil {
		return apierrors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	SelectResult:= stmt.QueryRow(Bapp.BeneficiaryID)

	if err:=SelectResult.Scan(&beneficiary_DB_Data.DoseType,&beneficiary_DB_Data.BookingDate);err!=nil{		
		return apierrors.NewInternalServerError(
			fmt.Sprintf("error When trying to get user with Beneficiary Id %s:%s",Bapp.BeneficiaryID,err.Error()))
		}

	if doseErr:=Bapp.ValidateDoseCheck(beneficiary_DB_Data.DoseType,beneficiary_DB_Data.BookingDate);doseErr!=nil{
		return doseErr
	}	
	return nil
}

func (b *BookAppointment) ValidateDoseCheck(dosetype int, dbDate string) *apierrors.RestErr {
	if b.DoseType == dosetype {
		return apierrors.NewInternalServerError("Vaccination already taken for the given dose type")

	} else if b.DoseType > dosetype {
		gap := date.CheckTimeGap(date.ConvertStringtoTime(b.BookingDate),date.ConvertStringtoTime(dbDate))
		fmt.Println("No of gap Days ",gap)
		if gap < 15 {
			return (apierrors.NewBadRequestError("Second Dose Should be taken after 15 days"))
		}
	}

	return nil
}