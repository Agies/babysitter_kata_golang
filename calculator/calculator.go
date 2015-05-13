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
	_, error := time.Parse(dateFormat, startTime);
	if (error != nil) {
		return 0, error;
	}
	if (leaveTime == "") {
		return 0, errors.New("Leave time is required");
	}
	if (bedTime == "") {
		return 0, errors.New("Bed time is required");
	}
	return 0, nil;
}