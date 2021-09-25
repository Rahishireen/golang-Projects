package app

import controllers "vaccinationdrive/Controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
	router.POST("/user_registration",controllers.CreateUser)
	//router.GET("/booking_appointment",controllers.BookAppointment)
}