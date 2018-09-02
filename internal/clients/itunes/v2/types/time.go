package types

import (
	"strconv"
	"time"
)

type Time struct {
	Value time.Time
}

func (t *Time) MarshalJSON() ([]byte, error) {
	return []byte(t.Value.Format(strconv.Quote("2006-01-02"))), nil
}

func (t *Time) UnmarshalJSON(data []byte) error {
	var err error
	t.Value, err = time.ParseInLocation(strconv.Quote("2006-01-02"), string(data), time.FixedZone("UTC", 0))
	return err
}