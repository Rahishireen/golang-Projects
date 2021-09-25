package date

import (
	"fmt"
	"time"
)
const(
	//apiTimeFormat="2021-06-13 15:04:05"
	apiTimeFormat="2006-01-02 15:04:05"
)
func GetTime() time.Time {
	fmt.Println(time.Now())
	return time.Now()
}

func GetTimeString() string{
	return GetTime().Format(apiTimeFormat)
}