package timex

import "time"

// Unix returns the local Time corresponding to the given Unix time,
// sec seconds and nsec nanoseconds since January 1, 1970 UTC.
// It is valid to pass nsec outside the range [0, 999999999].
// Not all sec values have a corresponding time value. One such
// value is 1<<63-1 (the largest int64 value).
func Unix(t int64) time.Time {
	return time.Unix(t/1e9, t%1e9)
}

// ParseString2Time serve caller to given a string to parse time
func ParseString2Time(str string) time.Time {
	loc, _ := time.LoadLocation("Asia/Taipei")
	layout := "2006-01-02 15:04:05"

	t, err := time.ParseInLocation(layout, str, loc)
	if err != nil {
		return time.Time{}
	}

	return t
}

// ParseYYYYMMddHHmmss serve caller to given a string to parse time
func ParseYYYYMMddHHmmss(str string) time.Time {
	loc, _ := time.LoadLocation("Asia/Taipei")
	layout := "20060102150405"

	t, err := time.ParseInLocation(layout, str, loc)
	if err != nil {
		return time.Time{}
	}

	return t

}
