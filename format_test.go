package timefmt_test

import (
	"errors"
	"strings"
	"testing"
	"time"

	timefmt "github.com/itchyny/timefmt-go"
)

var formatTestCases = []struct {
	source    string
	format    string
	t         time.Time
	formatErr error
}{
	{
		source: "2020",
		format: "%Y",
		t:      time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "0999",
		format: "%Y",
		t:      time.Date(999, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "0099",
		format: "%Y",
		t:      time.Date(99, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "[2020]",
		format: "[%Y]",
		t:      time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "2020-05",
		format: "%Y-%m",
		t:      time.Date(2020, time.May, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "2020-09-10",
		format: "%Y-%m-%d",
		t:      time.Date(2020, time.September, 10, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "2020/10/09",
		format: "%Y/%m/%d",
		t:      time.Date(2020, time.October, 9, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "October",
		format: "%B",
		t:      time.Date(0, time.October, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "Sep",
		format: "%b",
		t:      time.Date(0, time.September, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "Monday Mon 1",
		format: "%A %a %w",
		t:      time.Date(0, time.January, 3, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "Tuesday Tue 2",
		format: "%A %a %w",
		t:      time.Date(0, time.January, 4, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "Saturday Sat 6",
		format: "%A %a %w",
		t:      time.Date(0, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "Sunday Sun 0",
		format: "%A %a %w",
		t:      time.Date(0, time.January, 2, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "2020-09-08 07:06:05",
		format: "%Y-%m-%d %H:%M:%S",
		t:      time.Date(2020, time.September, 8, 7, 6, 5, 4, time.UTC),
	},
	{
		format:    "%E",
		formatErr: errors.New(`unexpected format: "%E"`),
	},
}

func TestFormat(t *testing.T) {
	for _, tc := range formatTestCases {
		t.Run(tc.source+"/"+tc.format, func(t *testing.T) {
			got, err := timefmt.Format(tc.t, tc.format)
			if tc.formatErr == nil {
				if err != nil {
					t.Fatalf("expected no error but got: %v", err)
				}
				if got != tc.source {
					t.Errorf("expected: %v, got: %v", tc.source, got)
				}
			} else {
				if err == nil {
					t.Fatalf("expected error %v but got: %v", tc.formatErr, err)
				}
				if !strings.Contains(err.Error(), tc.formatErr.Error()) {
					t.Errorf("expected: %v, got: %v", tc.formatErr, err)
				}
			}
		})
	}
}
