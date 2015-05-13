package calculator

import (
	"time"
	"errors"
)

var utcLoc, err = time.LoadLocation("");

func Calculate(startTime string, leaveTime string, bedTime string) (int, error) {
	if (startTime == "") {
		return 0, errors.New("Start time is required");
	}
	if (leaveTime == "") {
		return 0, errors.New("Leave time is required");
	}
	if (bedTime == "") {
		return 0, errors.New("Bed time is required");
	}
	return 0, nil;
}