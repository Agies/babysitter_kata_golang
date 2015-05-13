package calculator

import (
	"time"
	"errors"
)

var utcLoc, err = time.LoadLocation("");

func Calculate(startTime string, leaveTime string) (int, error) {
	if (startTime == "") {
		return 0, errors.New("Start time is required");
	}
	if (leaveTime == "") {
		return 0, errors.New("Leave time is required");
	}
	return 0, nil;
}