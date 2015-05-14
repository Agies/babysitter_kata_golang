package calculator

import (
	"math"
	"time"
	"errors"
)

var utcLoc, err = time.LoadLocation("");
const dateFormat = "01/02/2006 3:04pm";
var MinStartTime = 17;
var MaxLeaveTime = 4;
var BeforeBedtimePayment float64 = 12;
var BedtimeToMidnightPayment float64 = 8;
var AfterMidnightPayment float64 = 16;

func Calculate(startTime string, leaveTime string, bedTime string) (int, error) {
	if (startTime == "") {
		return 0, errors.New("Start time is required");
	}
	parsedStartTime, error := time.ParseInLocation(dateFormat, startTime, utcLoc);
	if (error != nil) {
		return 0, error;
	}
	if (leaveTime == "") {
		return 0, errors.New("Leave time is required");
	}
	parsedLeaveTime, error := time.Parse(dateFormat, leaveTime);
	if (error != nil) {
		return 0, error;
	}
	if (bedTime == "") {
		return 0, errors.New("Bed time is required");
	}
	parsedBedTime, error := time.Parse(dateFormat, bedTime);
	if (error != nil) {
		return 0, error;
	}
	year, month, day := parsedStartTime.Date();
	if (parsedStartTime.Before(time.Date(year,month,day,MinStartTime,0,0,0,utcLoc))) {
		return 0, errors.New("Start time must be no earlier than 5PM");
	}
	if (parsedLeaveTime.After(time.Date(year,month,day,MaxLeaveTime,0,0,0,utcLoc).AddDate(0,0,1))) {
		return 0, errors.New("Leave time must be no later than 4AM");
	}
	
	var midnight = time.Date(year,month,day,0,0,0,0,utcLoc).AddDate(0,0,1);
	
	beforeBedTime := math.Floor(parsedBedTime.Sub(parsedStartTime).Hours());
	var beforeBedTimeAmount float64 = 0.0;
	if beforeBedTime > 0 { 
		beforeBedTimeAmount = beforeBedTime * BeforeBedtimePayment; 
	} 
	
	var timeToUse = parsedLeaveTime;
	if midnight.Before(parsedLeaveTime) {
		timeToUse = midnight;
	}
	afterBedTimeBeforeMidnight := math.Floor(timeToUse.Sub(parsedBedTime).Hours());
	var afterBedTimeBeforeMidnightAmount float64 = 0.0;
	if afterBedTimeBeforeMidnight > 0 { 
		afterBedTimeBeforeMidnightAmount = afterBedTimeBeforeMidnight * BedtimeToMidnightPayment; 
	} 
	
	afterMidnight := math.Floor(parsedLeaveTime.Sub(midnight).Hours());
	var afterMidnightAmount float64 = 0.0;
	if afterMidnight > 0 { 
		afterMidnightAmount = afterMidnight * AfterMidnightPayment; 
	} 
	
	return int(beforeBedTimeAmount + afterMidnightAmount + afterBedTimeBeforeMidnightAmount), nil;
}