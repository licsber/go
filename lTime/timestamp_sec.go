package lTime

import (
	"strconv"
	"time"
)

// TimeStampSec aka Unix Time
type TimeStampSec time.Time

func (ts *TimeStampSec) MarshalJSON() ([]byte, error) {
	t := time.Time(*ts)
	return []byte(strconv.FormatInt(t.Unix(), 10)), nil
}

func (ts *TimeStampSec) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	seconds, err := strconv.ParseInt(string(data), 10, 64)
	*ts = TimeStampSec(time.Unix(seconds, 0))
	return err
}
