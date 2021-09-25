package appointment

import apierrors "vaccinationdrive/utills/error"

var VaccinationCenter = []string{"Nungambakkam", "Tambaram", "Velachery", "Shozhinganallur"}

type BookAppointment struct {
	BeneficiaryID string `json:"name"`
	SlotBooking   Slot   `json:"availability"`
	DoseType      int    `json:"dose"`
	TimeSlot      []int  `json:"aadharnumber"`
	VaccineCenter string
}

type Slot struct {
	BookingDate string `json:"bookingdate"`
	TimeSlot    string `json:"timeslot"`
}

func (a *BookAppointment) VaccineCenterCheck() *apierrors.RestErr{

	for _, value := range VaccinationCenter {
		if a.VaccineCenter == value {
			return nil
		}
	}

	return apierrors.NewBadRequestError("Vaccination is not available in this Center,Please choose between 1.Nungambakkam 2.Tambaram 3.Velachery 4.Shozhinganallur")

}
