package main

import (
	"log"
	"testing"
)

// TODO:
// - Handle RFC specific formats that are not ISO supported.
func TestDatetimeString(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"2018-03-01T22:07:04.074Z", "2018-03-01T22:07:04.074Z"}, // subseconds and original timezone preserved
		{"2006-01-02T15:04:05Z", "2006-01-02T15:04:05Z"},         // RFC but in ISO friendly format
	}
	filter := NewFilter()

	for _, tt := range tests {
		log.Printf("running test for input [%s]", tt.input)
		t.Run(tt.input, func(t *testing.T) {
			ans, err := filter.DatetimeString(tt.input)
			if err != nil {
				t.Fatal("cleaning DatetimeString fatal error", err)
			}
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
