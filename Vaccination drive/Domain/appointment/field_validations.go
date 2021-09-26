package appointment

import (
	"fmt"
	"regexp"
	apierrors "vaccinationdrive/utills/error"
)
var(

AvailableVaccinationCenter = []string{"Nungambakkam", "Tambaram", "Velachery", "Shozhinganallur"}
AvailableTimeSlot          = []string{"9.30 am to 11.30 am","2pm to 4pm","6pm to 8pm"}
AvailableDoseType          =[]int {1,2}

)

func (a *BookAppointment) VaccineCenterValidation() *apierrors.RestErr{

	for _, value := range AvailableVaccinationCenter {
		if a.UserVC == value {
			return nil
		}
	}

	return apierrors.NewBadRequestError(
		fmt.Sprintf("Vaccination is not available %s,Please choose between 1.Nungambakkam 2.Tambaram 3.Velachery 4.Shozhinganallur",a.UserVC))

}
func (a *BookAppointment) TimeSlotValidation() *apierrors.RestErr{

	for _, value := range AvailableTimeSlot {
		if a.TimeSlot == value {
			return nil
		}
	}

	return apierrors.NewBadRequestError(
		fmt.Sprintf("Selected Slot is not available for the time %s,Please choose between 9.30 am to 11.30 am /2pm to 4pm /6pm to 8pm",a.TimeSlot))

}

func (a *BookAppointment) DoseTypeValidation() *apierrors.RestErr{

	for _, value := range AvailableDoseType {
		if a.DoseType == value {
			return nil
		}
	}

	return apierrors.NewBadRequestError(
		fmt.Sprintf ("Selected % d Dose is not applicable,Please choose between 1 (first Dose) / 2 (second Dose)",a.DoseType))

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