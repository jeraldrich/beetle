package main

import (
	"fmt"
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

// NewFilter configures filter with extenable cleaner methods
// TODO: Implement filter results, cleaner formatting options
func NewFilter() *Filter {
	c := new(Cleaner)
	return &Filter{c}
}

// Scylla uses ISO-8601, but Go uses RFC-3339 time format.
// Takes a datetime string of any type and convert to ISO-8601 datetime string.
// For datetime strings that specify a location, stored timezone will be preserved.
// Otherwise, timezone will convert to UTC.
func (f *Filter) DatetimeString(datetimeInput string) (iso8601String string, err error) {
	rfcTemplate := "2006-01-02T15:04:05Z"     // RFC 3339
	isoTemplate := "2018-03-01T22:07:04.074Z" // ISO

	isoConvertedInput, err := datetime.Parse(datetimeInput, time.UTC)
	if err != nil {
		// TODO: Better exception handling for incorrectly formatted dates.
		fmt.Printf("Failed converting datetime string to iso8601: [%s]", datetimeInput)
		return "", err
	}

	// If the input matches the converted without microseconds, then return the input.
	// Using RFC 3339 to compare since datetime standard libray has trouble with some ISO.
	// Otherwise, the time.parse called at the beginning will incorrectly
	// add 1-3 microseconds based off duraction of the method.
	convertedFromIsoToRfc, err := time.Parse(rfcTemplate, isoConvertedInput.Format(rfcTemplate))
	if isoConvertedInput.Format(rfcTemplate) == convertedFromIsoToRfc.Format(rfcTemplate) {
		return datetimeInput, err
	}

	// Now that we have a valid iso object,
	// return as iso8601 string with subseconds.
	return isoConvertedInput.Format(isoTemplate), err
}

func (f *Filter) Uuid(input string) (uuid gocql.UUID, err error) {
	result, err := gocql.ParseUUID(input)
	if err != nil {
		fmt.Printf("Invalid UUID %s", input)
		var blankUuid [16]byte
		return blankUuid, err
	}

	return result, err
}
