package appointment

import (
	"fmt"
	"vaccinationdrive/Datasources/mysql/users_db"
	apierrors "vaccinationdrive/utills/error"
)

func (Bapp *BookAppointment) VaccinationAvailabilityCheck(query string)*apierrors.RestErr{
	var booked_count int
	
	stmt,err:=users_db.Client.Prepare(query)
	if err!=nil{
		return apierrors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	
	SelectResult,err :=stmt.Query(Bapp.BookingDate,Bapp.UserVC)	

	if err!=nil{
		return apierrors.NewInternalServerError(fmt.Sprintf("error when trying to get Booked Vaccine count :%s",err.Error()))
	}else {
		
		for SelectResult.Next(){
			SelectResult.Scan(&booked_count)
		}
	}
	if booked_count >=30 {
		return apierrors.NewInternalServerError(fmt.Sprintf("All slots are booked for the date %s",Bapp.BookingDate))
	}
	return nil
	
}

func (Bapp *BookAppointment) DoseAvailabilityCheck(query string)*apierrors.RestErr{
	var dose_count int
	stmt,err:=users_db.Client.Prepare(query)
	if err!=nil{
		return apierrors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	SelectResult,err :=stmt.Query(Bapp.BookingDate,Bapp.DoseType,Bapp.UserVC)	

	if err!=nil{
		return apierrors.NewInternalServerError(fmt.Sprintf("error when trying to get Booked dose count:%s",err.Error()))
	}else {
		
		for SelectResult.Next(){
			SelectResult.Scan(&dose_count)
		}
	}
	if dose_count >=15 {
		return apierrors.NewInternalServerError(fmt.Sprintf("All slots are booked for the %d dose",Bapp.DoseType))
	}
	return nil
	
}

func (Bapp *BookAppointment) TimeslotAvailabilityCheck(query string)*apierrors.RestErr{
	var timeSlot_count int
	stmt,err:=users_db.Client.Prepare(query)
	if err!=nil{
		return apierrors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	SelectResult,err :=stmt.Query(Bapp.BookingStatus,Bapp.TimeSlot,Bapp.UserVC)	

	if err!=nil{
		return apierrors.NewInternalServerError(fmt.Sprintf("error when trying to get Booked TimeSlot count:%s",err.Error()))
	}else {
		
		for SelectResult.Next(){
			SelectResult.Scan(&timeSlot_count)
		}
	}
	if timeSlot_count >=10 {
		return apierrors.NewInternalServerError(fmt.Sprintf("All slots are booked for the %s Time Slot",Bapp.TimeSlot))
	}
	return nil
	
}



