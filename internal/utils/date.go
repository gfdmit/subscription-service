package utils

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type CustomDate struct {
	time.Time
}

func (c *CustomDate) UnmarshalJSON(b []byte) (err error) {
	c.Time, err = CustomDateToTime(string(b))
	if err != nil {
		return err
	}
	return nil
}

func CustomDateToTime(customDate string) (time.Time, error) {
	s := strings.Trim(customDate, "\"")
	if s == "null" {
		return time.Time{}, nil
	}

	tokens := strings.Split(s, "-")
	month, err := strconv.Atoi(tokens[0])
	if err != nil {
		return time.Time{}, err
	}

	year, err := strconv.Atoi(tokens[1])
	if err != nil {
		return time.Time{}, err
	}

	date := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	return date, nil
}

func (c CustomDate) MarshalJSON() ([]byte, error) {
	if c.Time.IsZero() {
		return nil, nil
	}
	return []byte(fmt.Sprintf(`"%02d-%d"`, c.Time.Month(), c.Time.Year())), nil
}

func (c *CustomDate) Scan(value interface{}) error {
	if value == nil {
		c.Time = time.Time{}
		return nil
	}
	t, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("cannot scan %T into CustomDate", value)
	}
	c.Time = t
	return nil
}

func (c CustomDate) Value() (driver.Value, error) {
	if c.IsZero() {
		return nil, nil
	}
	return c.Time, nil
}
