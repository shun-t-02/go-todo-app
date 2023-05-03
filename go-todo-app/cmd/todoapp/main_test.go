package main

import (
	"reflect"
	"testing"
	"time"
)

func TestFormatDateTime(t *testing.T) {
	tests := []struct {
		name string
		time time.Time
		want string
	}{
		{name: "test1", time: time.Date(2023, 5, 1, 9, 0, 0, 0, time.Local), want: "2023-05-01 09:00"},
		{name: "test2", time: time.Time{}, want: ""},
	}

	for _, tc := range tests {
		got := formatDateTime(tc.time)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("%s: expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}
