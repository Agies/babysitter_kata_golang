package calculator

import (
	"time"
	"errors"
)

var utcLoc, err = time.LoadLocation("");

func Calculate(startTime string) (int, error) {
	if (startTime == "") {
		return 0, errors.New("Start time is required");
	}
	return 0, nil;
}