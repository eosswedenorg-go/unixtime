package unixtime

import (
	"encoding/json"
	"strconv"
	"time"
)

// UnixTime is a simple wrapper to handle unix timestamps in json data.
// Note that the value is in milliseconds.
type Time int64

func (ts *Time) UnmarshalJSON(b []byte) error {
	var i int64

	// "borrowed" from "gopkg.in/guregu/null.v4" abit.
	if err := json.Unmarshal(b, &i); err != nil {

		// If unmarshal to int64 fails, we assume that its a numeric string.
		var str string
		if err := json.Unmarshal(b, &str); err != nil {
			return err
		}

		// Then we need to parse the string into int64
		i, err = strconv.ParseInt(str, 10, 64)
		if err != nil {
			return err
		}
	}

	*ts = Time(i)
	return nil
}

func (ts Time) Time() time.Time {
	v := int64(ts)
	return time.UnixMilli(v).UTC()
}

func (ts *Time) FromTime(t time.Time) {
	*ts = Time(t.UnixMilli())
}
