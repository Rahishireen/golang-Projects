package appointment

import (
	"fmt"
	"time"
	"vaccinationdrive/Domain/users"
	"vaccinationdrive/utills/date"
	apierrors "vaccinationdrive/utills/error"
)

const (
	QuerySelectVaccination = "SELECT count(*) as No_Booked_users FROM booking_appointment WHERE booking_date in (?) AND Vaccination_Center in (?);"
	QuerySelectDose        = "SELECT count(*) as Dose_Count FROM booking_appointment WHERE booking_date in (?) AND Dose_Type in (?) AND Vaccination_Center in (?);"
	QuerySelectTimeSlot    = "SELECT count(*) as Time_Slot FROM booking_appointment WHERE booking_date in (?) AND Time_Slot in (?) AND Vaccination_Center in (?);"
	errorNoRows            = "no rows in result set"
	Duplicate              = "Duplicate"
	QueryBeneficiaryCountAdd  = "SELECT count(*) as Beneficiary_Count FROM booking_appointment WHERE Beneficiary_Id in (?)"
	QueryBeneficiaryCountUpdate  = "SELECT count(*) as Beneficiary_Count FROM booking_appointment WHERE Beneficiary_Id in (?) AND  Time_Slot in (?) AND Dose_Type in (?) AND Vaccination_Center in (?)"
	QueryBeneficiaryCountDelete  = "SELECT count(*) as Beneficiary_Count FROM booking_appointment WHERE Beneficiary_Id in (?) AND booking_date in (?) AND  Time_Slot in (?) AND Dose_Type in (?) AND Vaccination_Center in (?)"
	QueryGetId             ="SELECT Dose_Type,Booking_date FROM booking_appointment WHERE Beneficiary_Id = ?"
)


func (Bapp *BookAppointment) BasicFieldValidation() *apierrors.RestErr{
	var user users.UserRegistration
	user.BeneficiaryID=Bapp.BeneficiaryID
//Beneficiary Id Validation- Check for Beneficiary Id in Users DB,Check the pattern
	//1.Pattern Validation
	if beneficiaryErr := FieldValidation(14,`\d{14}`,Bapp.BeneficiaryID,"Beneficiary Id","14 digit"); beneficiaryErr != nil {
		return beneficiaryErr
	}
	//2.Check Entry in Users DB
	if beneficiaryEntryErr:= user.Get();beneficiaryEntryErr!=nil{
		return beneficiaryEntryErr
	}
//Vaccination Center Validation
	if vcErr:=Bapp.VaccineCenterValidation();vcErr!=nil{
		return vcErr

	}
//Booking_date Validation - Booking should be less than or equal to 90 days from Current date,Check Pattern
	//1.Pattern Validation
	if dateErr := FieldValidation(10,`\d{2}-\d{2}-\d{4}`,Bapp.BookingDate,"Booking Date","DD-MM-YYYY"); dateErr != nil {
		return dateErr
	}
	//2.No of days gap Check - Should be less than or equal to 90 days
	if Bapp.BookingStatus!=3{
	gap := date.CheckTimeGap(date.ConvertStringtoTime(Bapp.BookingDate),time.Now())
		fmt.Println("No of gap Days ",gap)
		if gap > 90 {
			return (apierrors.NewBadRequestError("Slot cant be booked before 90 days"))
		}
	}


//TimeSlot Validation
	if tsErr:=Bapp.TimeSlotValidation();tsErr!=nil{
		return tsErr
	}
//DoseType Validation
	if dtErr:=Bapp.DoseTypeValidation();dtErr!=nil{
		return dtErr
	}

	return nil
}
func (Bapp *BookAppointment) GeneralAvailabityValidation() *apierrors.RestErr {
	//Availability of Vaccination in the Booked Center on the Booked Date
	if vaccinationErr := Bapp.VaccinationAvailabilityCheck(QuerySelectVaccination); vaccinationErr != nil {
		return vaccinationErr
	}
	//Availability of Dose in the Booked Center on the Booked Date
	if doseErr:=Bapp.DoseAvailabilityCheck(QuerySelectDose);doseErr!=nil{
		return doseErr
	}	

	//Availability of Vaccination in the booked time slot on the booked date,booked vaccination center
	if timeSlotErr:=Bapp.TimeslotAvailabilityCheck(QuerySelectTimeSlot);timeSlotErr!=nil{
		return timeSlotErr
	}
	return nil
}



func (Bapp *BookAppointment) UserEligibleValidationAdd() *apierrors.RestErr{

	//Check for no of entries of Beneficiary id in Booking Appointment dataBase
	beneficiarycount,beneficiaryErr:=Bapp.CheckBeneficiaryInDBAdd(QueryBeneficiaryCountAdd)
	if beneficiaryErr!=nil{
		return beneficiaryErr
	}

	switch beneficiarycount {
	case 0:
		if Bapp.DoseType!=1{
			return apierrors.NewBadRequestError("Second Dose Should be taken after First Dose")
		}
	case 1:
		  if err:= Bapp.GetVaccineDetailsInDB(QueryGetId);err!=nil{
			  return err
		  }
	case 2:
		return apierrors.NewBadRequestError("Booking limit Exhausted,Users allowed to book upto 2 appointments")
	default:
		return apierrors.NewInternalServerError("DB contains More than 2 entries,which is Invalid")		
	}
	return nil
}

func (Bapp *BookAppointment) UserEligibleValidationUpdate() *apierrors.RestErr{

	//Check for no of entries of Beneficiary id in Booking Appointment dataBase
	beneficiarycount,beneficiaryErr:=Bapp.CheckBeneficiaryInDBUpdate(QueryBeneficiaryCountUpdate)
	if beneficiaryErr!=nil{
		return beneficiaryErr
	}

	switch beneficiarycount {
	case 0:		
			return apierrors.NewNotFoundError("No entry for the given data")	
	case 1:
		  return nil
	case 2:
		if err:= Bapp.GetVaccineDetailsInDB(QueryGetId);err!=nil{
			return err
		}
	default:
		return apierrors.NewInternalServerError("DB contains More than 2 entries,which is Invalid")		
	}
	return nil
}

func (Bapp *BookAppointment) UserEligibleValidationDelete() *apierrors.RestErr{

	//Check for no of entries of Beneficiary id in Booking Appointment dataBase
	beneficiarycount,beneficiaryErr:=Bapp.CheckBeneficiaryInDBDelete(QueryBeneficiaryCountDelete)
	if beneficiaryErr!=nil{
		return beneficiaryErr
	}
	switch beneficiarycount {
	case 0:		
			return apierrors.NewNotFoundError("No entry for the given data")		
	case 1:
		  return nil
	case 2:
			return nil
	default:
		return apierrors.NewInternalServerError("DB contains More than 2 entries,which is Invalid")		
	}
	//return nil
}