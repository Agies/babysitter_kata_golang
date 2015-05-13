package calculator

import (
	"time"
	"errors"
)

var utcLoc, err = time.LoadLocation("");
const dateFormat = "01/02/2006 3:04pm";
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
	_, error = time.Parse(dateFormat, leaveTime);
	if (error != nil) {
		return 0, error;
	}
	if (bedTime == "") {
		return 0, errors.New("Bed time is required");
	}
	_, error = time.Parse(dateFormat, bedTime);
	if (error != nil) {
		return 0, error;
	}
	year, month, day := parsedStartTime.Date();
	if (parsedStartTime.Before(time.Date(year,month,day,17,0,0,0,utcLoc))) {
		return 0, errors.New("Start time must be no earlier than 5PM");
	}
	return 0, nil;
}