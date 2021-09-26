package controllers

import (
	"net/http"
	"vaccinationdrive/Domain/users"
	services "vaccinationdrive/Services"
	apierrors "vaccinationdrive/utills/error"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var beneficiary users.UserRegistration
	if err:=c.ShouldBindJSON(&beneficiary);err!=nil{
		resterr:=apierrors.NewBadRequestError("Invalid json Body")
		c.JSON(resterr.Code,resterr)
		return
	}
	result,saveErr:= services.CreateUser(beneficiary)
	if saveErr!=nil{
		c.JSON(saveErr.Code,saveErr)
		return
	}	

	c.JSON(http.StatusAccepted,result)

}