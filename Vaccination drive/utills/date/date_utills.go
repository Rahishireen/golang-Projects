package date

import (
	"fmt"
	"math"
	"time"
)
const(
	//apiTimeFormat="2021-06-13 15:04:05"
	apiTimeFormat="2006-01-02 15:04:05"
	timeFormat="01-02-2006"
)
func GetTime() time.Time {
	fmt.Println(time.Now())
	return time.Now()
}

//Get Current time in String
func GetTimeString() string{
	return GetTime().Format(apiTimeFormat)
}

//String to Time
func ConvertStringtoTime(inputDate string) time.Time{
	date,err:=time.Parse(timeFormat,inputDate)
	fmt.Println(err)
	return date
}

//Time to String
func ConvertTimetoString(inputDate time.Time)string{
	return inputDate.Format(timeFormat)
}

func CheckTimeGap(date1 time.Time,date2 time.Time)int{

	if date1.Before(date2){
		temp:=date2
		date2=date1
		date1=temp
	}
	
	duration:=date1.Sub(date2)

	gap:=int(math.Ceil(duration.Hours()/24))

	return gap

}
