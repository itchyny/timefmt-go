package timefmt_test

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/itchyny/timefmt-go"
)

var parseTestCases = []struct {
	source   string
	format   string
	t        time.Time
	parseErr error
}{
	{
		source: "2020",
		format: "%Y",
		t:      time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "[2020]",
		format: "[%Y]",
		t:      time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "2020",
		format: "%Y",
		t:      time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source:   "20",
		format:   "%Y",
		parseErr: errors.New("cannot parse %Y"),
	},
	{
		source:   "20xxx",
		format:   "%Y",
		parseErr: errors.New(`strconv.Atoi: parsing "20xx"`),
	},
	{
		source: "2020-05",
		format: "%Y-%m",
		t:      time.Date(2020, time.May, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "2020/1",
		format: "%Y/%m",
		t:      time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "2020/9",
		format: "%Y/%m",
		t:      time.Date(2020, time.September, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "2020/10/09",
		format: "%Y/%m/%d",
		t:      time.Date(2020, time.October, 9, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "2020/10/1",
		format: "%Y/%m/%d",
		t:      time.Date(2020, time.October, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "2020-12-12",
		format: "%Y-%m-%d",
		t:      time.Date(2020, time.December, 12, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "2020-1-1",
		format: "%Y-%m-%d",
		t:      time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "[2020-1-1]",
		format: "[%Y-%m-%d]",
		t:      time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source:   "2020-",
		format:   "%Y-%m-%d",
		parseErr: errors.New("cannot parse %m"),
	},
	{
		source:   "2020-1-",
		format:   "%Y-%m-%d",
		parseErr: errors.New("cannot parse %d"),
	},
	{
		source: "201111",
		format: "%y%m%d",
		t:      time.Date(2020, time.November, 11, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "Jan",
		format: "%b",
		t:      time.Date(0, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source:   "Ja",
		format:   "%b",
		parseErr: errors.New("cannot parse %b"),
	},
	{
		source: "Jul",
		format: "%b",
		t:      time.Date(0, time.July, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "Sep",
		format: "%b",
		t:      time.Date(0, time.September, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "MAY",
		format: "%b",
		t:      time.Date(0, time.May, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "SATURDAY",
		format: "%A",
		t:      time.Date(0, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "[sunday]",
		format: "[%A]",
		t:      time.Date(0, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "[Mon]",
		format: "[%a]",
		t:      time.Date(0, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source:   "Teu",
		format:   "%a",
		parseErr: errors.New("cannot parse %a"),
	},
	{
		source: "mondaymon1",
		format: "%A%a%w",
		t:      time.Date(0, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source:   "mooday",
		format:   "%A",
		parseErr: errors.New("cannot parse %A"),
	},
	{
		source:   "7",
		format:   "%w",
		parseErr: errors.New("cannot parse %w"),
	},
	{
		source: "2020-09-08 07:06:05",
		format: "%Y-%m-%d %H:%M:%S",
		t:      time.Date(2020, time.September, 8, 7, 6, 5, 0, time.UTC),
	},
	{
		source: "1:2:3.45",
		format: "%H:%M:%S.%f",
		t:      time.Date(0, time.January, 1, 1, 2, 3, 450000000, time.UTC),
	},
	{
		source: "1213145678912",
		format: "%H%M%S%f%d",
		t:      time.Date(0, time.January, 2, 12, 13, 14, 567891000, time.UTC),
	},
	{
		source: "12:13:14 AM",
		format: "%I:%M:%S %p",
		t:      time.Date(0, time.January, 1, 12, 13, 14, 0, time.UTC),
	},
	{
		source: "01:14:15pm",
		format: "%I:%M:%S%p",
		t:      time.Date(0, time.January, 1, 13, 14, 15, 0, time.UTC),
	},
	{
		source: "PM 11:14:15",
		format: "%p %I:%M:%S",
		t:      time.Date(0, time.January, 1, 23, 14, 15, 0, time.UTC),
	},
	{
		source:   "pp",
		format:   "%p",
		parseErr: errors.New("cannot parse %p"),
	},
	{
		format:   "%E",
		parseErr: errors.New(`unexpected format: "%E"`),
	},
}

func TestParse(t *testing.T) {
	for _, tc := range parseTestCases {
		t.Run(tc.source+"/"+tc.format, func(t *testing.T) {
			got, err := timefmt.Parse(tc.source, tc.format)
			if tc.parseErr == nil {
				if err != nil {
					t.Fatalf("expected no error but got: %v", err)
				}
				if !got.Equal(tc.t) {
					t.Errorf("expected: %v, got: %v", tc.t, got)
				}
			} else {
				if err == nil {
					t.Fatalf("expected error %v but got: %v", tc.parseErr, err)
				}
				if !strings.Contains(err.Error(), tc.parseErr.Error()) {
					t.Errorf("expected: %v, got: %v", tc.parseErr, err)
				}
			}
		})
	}
}
