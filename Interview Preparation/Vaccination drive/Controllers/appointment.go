package controllers

import (
	"net/http"
	"vaccinationdrive/Domain/appointment"
	apierrors "vaccinationdrive/utills/error"

	"github.com/gin-gonic/gin"
)

func BookAppointment(c *gin.Context) {

	var beneficiary appointment.BookAppointment
	if err:=c.ShouldBindJSON(&beneficiary);err!=nil{
		resterr:=apierrors.NewBadRequestError("Invalid json Body")
		c.JSON(resterr.Code,resterr)
		return
	}
	result,saveErr:= services.check(beneficiary)
	if saveErr!=nil{
		c.JSON(saveErr.Code,saveErr)
		return
	}	

	c.JSON(http.StatusAccepted,result)

}