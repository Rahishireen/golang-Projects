package appointment

import (
	"fmt"
	"time"
	"vaccinationdrive/Datasources/mysql/users_db"
	apierrors "vaccinationdrive/utills/error"
)
const (
	QueryInsertUser = "INSERT INTO booking_appointment(Beneficiary_Id,Booking_Date,Time_Slot,Dose_type,Vaccination_Center,Date_Created) VALUES (?,?,?,?,?,?);"
	QueryUpdateUser ="UPDATE booking_appointment set Booking_Date = ? WHERE Beneficiary_Id = ? AND Time_Slot = ? AND Dose_Type = ? AND Vaccination_Center =? ;"
	QueryDeleteUser ="DELETE FROM booking_appointment WHERE Beneficiary_Id = ?  AND Booking_Date = ? and Time_Slot = ? AND Dose_Type = ? AND Vaccination_Center =? "
)

func (Bapp *BookAppointment) AddBeneficiary() *apierrors.RestErr {

	stmt, err := users_db.Client.Prepare(QueryInsertUser)
	if err != nil {
		return apierrors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(Bapp.BeneficiaryID,Bapp.BookingDate,Bapp.TimeSlot,Bapp.DoseType,Bapp.UserVC,time.Now())

	if err != nil {
		return apierrors.NewInternalServerError(fmt.Sprintf("error when trying to Add user in Booking_Appointment DB:%s", err.Error()))
	}

	_, err = insertResult.LastInsertId()
	if err != nil {
		return apierrors.NewInternalServerError(fmt.Sprintf("error when trying to Add user in Booking_Appointment DB:%s", err.Error()))
	}

	return nil
}

func (Bapp *BookAppointment) UpdateBeneficiary() *apierrors.RestErr {

	stmt, err := users_db.Client.Prepare(QueryUpdateUser)
	if err != nil {
		return apierrors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(Bapp.BookingDate,Bapp.BeneficiaryID,
		Bapp.TimeSlot,Bapp.DoseType,Bapp.UserVC)

	if err != nil {
		return apierrors.NewInternalServerError(fmt.Sprintf("error when trying to update user in Booking_Appointment DB:%s", err.Error()))
	}

	_, err = insertResult.LastInsertId()
	if err != nil {
		return apierrors.NewInternalServerError(fmt.Sprintf("error when trying to update user in Booking_Appointment DB:%s", err.Error()))
	}

	return nil
}

func (Bapp *BookAppointment) DeleteBeneficiary() *apierrors.RestErr {
	fmt.Println(QueryDeleteUser)

	stmt, err := users_db.Client.Prepare(QueryDeleteUser)
	if err != nil {
		return apierrors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	deleteResult, err := stmt.Exec(Bapp.BeneficiaryID,Bapp.BookingDate,Bapp.TimeSlot,Bapp.DoseType,Bapp.UserVC)
	
	if err != nil {
		return apierrors.NewInternalServerError(fmt.Sprintf("error when trying to delete user in Booking_Appointment DB:%s", err.Error()))
	}
	_, err = deleteResult.LastInsertId()
	if err != nil {
		return apierrors.NewInternalServerError(fmt.Sprintf("error when trying to delete user in Booking_Appointment DB:%s", err.Error()))
	}
	return nil
}