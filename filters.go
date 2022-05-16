package main

import (
	"errors"
	"time"

	"github.com/gocql/gocql"
	"github.com/nav-inc/datetime"
)

type Filter struct {
	cleaner *Cleaner
}

type Cleaner interface {
	DatetimeString(datetimeString string) (iso8601String string, err error)
	Uuid(input string) (uuid gocql.UUID, err error)
}

func NewFilter() *Filter {
	c := new(Cleaner)
	return &Filter{c}
}

// Scylla uses ISO-8601, but Go uses RFC-3339 time format.
// Takes a datetime string of any type and convert to ISO-8601 time.Time
// For datetime strings that specify a location, stored timezone will be preserved.
// Otherwise, timezone will convert to UTC.
func (f *Filter) DatetimeString(datetimeInput string) (iso8601 time.Time, err error) {
	rfcTemplate := "2006-01-02T15:04:05Z" // RFC 3339
	// For returning a zero value time.
	blankTime := time.Time{}

	if len(datetimeInput) == 0 {
		return blankTime, nil
	}

	isoConvertedInput, err := datetime.Parse(datetimeInput, time.UTC)
	if err != nil {
		// Track 2022/05/16 09:37:27 as field value error.
		err = errors.New("Invalid datetime. Not Iso OR rfc")
		return blankTime, err
	}

	// If the input matches the converted without microseconds, then return the input.
	// Using RFC 3339 to compare since datetime standard library has trouble with some ISO.
	// Otherwise, the time.parse called at the beginning will incorrectly
	// add 1-3 microseconds based off duraction of the method.
	convertedFromIsoToRfc, err := time.Parse(rfcTemplate, isoConvertedInput.Format(rfcTemplate))
	if isoConvertedInput.Format(rfcTemplate) == convertedFromIsoToRfc.Format(rfcTemplate) {
		return isoConvertedInput, err
	}

	// Now that we have a valid iso object,
	// return as iso8601 with subseconds.
	return isoConvertedInput, err
}

func (f *Filter) Uuid(input string) (uuid gocql.UUID, err error) {
	result, err := gocql.ParseUUID(input)

	if err != nil {
		// fmt.Printf("Invalid UUID %s", input)
		var blankUuid [16]byte
		return blankUuid, err
	}

	return result, err
}
