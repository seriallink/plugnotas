package plugnotas

import (
	"time"
)

const DateFormat = `"02/01/2006 15:04:05"` // dd/mm/yyyy hh:mm:ss

type DateTime time.Time

func NewDateTime(t time.Time) DateTime {
	return DateTime(time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location()))
}

func (d DateTime) Pointer() *DateTime {
	return &d
}

func (d DateTime) Time() time.Time {
	return time.Time(d)
}

func (d DateTime) String() string {
	return d.Time().String()
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(d).Format(DateFormat)), nil
}

func (d *DateTime) UnmarshalJSON(data []byte) error {
	t, err := time.Parse(DateFormat, string(data))
	if err == nil {
		*d = DateTime(t)
	}
	return err
}
