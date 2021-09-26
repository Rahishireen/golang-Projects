package services

import (
	"vaccinationdrive/Domain/appointment"
	apierrors "vaccinationdrive/utills/error"
)

func Appointment(Bapp appointment.BookAppointment) (*appointment.BookAppointment,*apierrors.RestErr){
	if FieldErr:=Bapp.BasicFieldValidation();FieldErr!=nil{
		return nil,FieldErr
	}
	switch Bapp.BookingStatus{
	case 1:
		if generalErr:=Bapp.GeneralAvailabityValidation();generalErr!=nil{
		return nil,generalErr
		}
		if beneficiaryErr:=Bapp.UserEligibleValidationAdd();beneficiaryErr!=nil{
		return nil,beneficiaryErr
		}

		if DBAdderr := Bapp.AddBeneficiary();DBAdderr!=nil{
			return nil,DBAdderr
		}
	    return &Bapp, nil
	case 2:
		if generalErr:=Bapp.GeneralAvailabityValidation();generalErr!=nil{
			return nil,generalErr
			}
		if beneficiaryErr:=Bapp.UserEligibleValidationUpdate();beneficiaryErr!=nil{
				return nil,beneficiaryErr
		                                                                                                                              	}
		if DBUpdateerr := Bapp.UpdateBeneficiary();DBUpdateerr!=nil{
				return nil,DBUpdateerr
			}
			return &Bapp, nil
		
	case 3:
		if beneficiaryErr:=Bapp.UserEligibleValidationDelete();beneficiaryErr!=nil{
			return nil,beneficiaryErr
			}
		if DBUpdateerr := Bapp.DeleteBeneficiary();DBUpdateerr!=nil{
			return nil,DBUpdateerr
		}
		return &Bapp, nil

	}

	return nil,apierrors.NewInternalServerError("Please provide valid Action,It Should be 1 (Add), 2 (Update), 3 (Delete)")
}