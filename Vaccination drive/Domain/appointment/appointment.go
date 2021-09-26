package appointment

type BookAppointment struct {
	BeneficiaryID string `json:"beneficiaryid"`
	BookingDate   string `json:"bookingdate"`
	TimeSlot      string `json:"timeslot"`
	DoseType      int    `json:"dose"`
	UserVC        string `json:"vaccinationcenter"`
	BookingStatus int    `json:"action"`
}
