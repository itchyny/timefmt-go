package timefmt_test

import (
	"strings"
	"testing"
	"time"

	timefmt "github.com/itchyny/timefmt-go"
)

var formatTestCases = []struct {
	format    string
	t         time.Time
	expected  string
	formatErr error
}{
	{
		format:   "%Y",
		t:        time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		expected: "2020",
	},
	{
		format:   "%Y",
		t:        time.Date(999, time.January, 1, 0, 0, 0, 0, time.UTC),
		expected: "0999",
	},
	{
		format:   "%Y",
		t:        time.Date(99, time.January, 1, 0, 0, 0, 0, time.UTC),
		expected: "0099",
	},
	{
		format:   "[%Y]",
		t:        time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		expected: "[2020]",
	},
	{
		format:   "%y",
		t:        time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		expected: "20",
	},
	{
		format:   "%y",
		t:        time.Date(2009, time.January, 1, 0, 0, 0, 0, time.UTC),
		expected: "09",
	},
	{
		format:   "%C",
		t:        time.Date(2009, time.January, 1, 0, 0, 0, 0, time.UTC),
		expected: "20",
	},
	{
		format:   "%C%y",
		t:        time.Date(1758, time.January, 1, 0, 0, 0, 0, time.UTC),
		expected: "1758",
	},
	{
		format:   "%Y-%m",
		t:        time.Date(2020, time.May, 1, 0, 0, 0, 0, time.UTC),
		expected: "2020-05",
	},
	{
		format:   "%Y-%m-%d",
		t:        time.Date(2020, time.September, 10, 0, 0, 0, 0, time.UTC),
		expected: "2020-09-10",
	},
	{
		format:   "%Y/%m/%d",
		t:        time.Date(2020, time.October, 9, 0, 0, 0, 0, time.UTC),
		expected: "2020/10/09",
	},
	{
		format:   "%e %-e",
		t:        time.Date(2020, time.January, 9, 0, 0, 0, 0, time.UTC),
		expected: " 9 9",
	},
	{
		format:   "%B",
		t:        time.Date(2020, time.October, 1, 0, 0, 0, 0, time.UTC),
		expected: "October",
	},
	{
		format:   "%b",
		t:        time.Date(2020, time.September, 1, 0, 0, 0, 0, time.UTC),
		expected: "Sep",
	},
	{
		format:   "%h",
		t:        time.Date(2020, time.September, 1, 0, 0, 0, 0, time.UTC),
		expected: "Sep",
	},
	{
		format:   "%A %a %w %u",
		t:        time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		expected: "Wednesday Wed 3 3",
	},
	{
		format:   "%A %a %w %u",
		t:        time.Date(2020, time.January, 2, 0, 0, 0, 0, time.UTC),
		expected: "Thursday Thu 4 4",
	},
	{
		format:   "%A %a %w %u",
		t:        time.Date(2020, time.January, 4, 0, 0, 0, 0, time.UTC),
		expected: "Saturday Sat 6 6",
	},
	{
		format:   "%A %a %w %u",
		t:        time.Date(2020, time.January, 5, 0, 0, 0, 0, time.UTC),
		expected: "Sunday Sun 0 7",
	},
	{
		format:   "%A %a %w %u",
		t:        time.Date(2020, time.January, 6, 0, 0, 0, 0, time.UTC),
		expected: "Monday Mon 1 1",
	},
	{
		format:   "%Y-%j-%-j",
		t:        time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		expected: "2020-001-1",
	},
	{
		format:   "%Y-%j-%-j",
		t:        time.Date(2020, time.February, 2, 0, 0, 0, 0, time.UTC),
		expected: "2020-033-33",
	},
	{
		format:   "%Y-%j-%-j",
		t:        time.Date(2020, time.August, 1, 0, 0, 0, 0, time.UTC),
		expected: "2020-214-214",
	},
	{
		format:   "%Y-%j-%-j",
		t:        time.Date(2020, time.December, 31, 0, 0, 0, 0, time.UTC),
		expected: "2020-366-366",
	},
	{
		format:   "%Y-%m-%d %H:%M:%S.%f",
		t:        time.Date(2020, time.September, 8, 7, 6, 5, 43210000, time.UTC),
		expected: "2020-09-08 07:06:05.043210",
	},
	{
		format:   "%-y-%-m-%-d %-H:%-M:%-S.%-f",
		t:        time.Date(2002, time.September, 8, 7, 6, 5, 43210000, time.UTC),
		expected: "2-9-8 7:6:5.043210",
	},
	{
		format:   "%H:%M:%S.%f",
		t:        time.Date(2020, time.January, 1, 1, 2, 3, 450000000, time.UTC),
		expected: "01:02:03.450000",
	},
	{
		format:   "%I:%M:%S %p",
		t:        time.Date(2020, time.January, 1, 1, 2, 3, 0, time.UTC),
		expected: "01:02:03 AM",
	},
	{
		format:   "%I:%M:%S %p",
		t:        time.Date(2020, time.January, 1, 12, 13, 14, 0, time.UTC),
		expected: "12:13:14 AM",
	},
	{
		format:   "%I:%M:%S %p",
		t:        time.Date(2020, time.January, 1, 13, 14, 15, 0, time.UTC),
		expected: "01:14:15 PM",
	},
	{
		format:   "%I:%M:%S %p",
		t:        time.Date(2020, time.January, 1, 23, 14, 15, 0, time.UTC),
		expected: "11:14:15 PM",
	},
	{
		format:   "%-I:%-M:%-S %p",
		t:        time.Date(2020, time.January, 1, 13, 2, 3, 0, time.UTC),
		expected: "1:2:3 PM",
	},
	{
		format:   "%k %-k",
		t:        time.Date(2020, time.January, 1, 9, 0, 0, 0, time.UTC),
		expected: " 9 9",
	},
	{
		format:   "%l %-l",
		t:        time.Date(2020, time.January, 1, 20, 0, 0, 0, time.UTC),
		expected: " 8 8",
	},
	{
		format:   "%R %r %T %D %x %X",
		t:        time.Date(2020, time.February, 9, 23, 14, 15, 0, time.UTC),
		expected: "23:14 11:14:15 PM 23:14:15 02/09/20 02/09/20 23:14:15",
	},
	{
		format:   "%F %T",
		t:        time.Date(2020, time.February, 9, 23, 14, 15, 0, time.UTC),
		expected: "2020-02-09 23:14:15",
	},
	{
		format:   "%v %X",
		t:        time.Date(2020, time.July, 9, 23, 14, 15, 0, time.UTC),
		expected: " 9-Jul-2020 23:14:15",
	},
	{
		format:   "%c",
		t:        time.Date(2020, time.February, 9, 23, 14, 15, 0, time.UTC),
		expected: "Sun Feb  9 23:14:15 2020",
	},
	{
		format:   "%F %T %z",
		t:        time.Date(2020, time.July, 24, 23, 14, 15, 0, time.UTC),
		expected: "2020-07-24 23:14:15 +0000",
	},
	{
		format:   "%F %T %z",
		t:        time.Date(2020, time.July, 24, 23, 14, 15, 0, time.FixedZone("UTC-8", -8*60*60)),
		expected: "2020-07-24 23:14:15 -0800",
	},
	{
		format:   "%F %T %z",
		t:        time.Date(2020, time.July, 24, 23, 14, 15, 0, time.FixedZone("UTC+9", 9*60*60)),
		expected: "2020-07-24 23:14:15 +0900",
	},
	{
		format:   "%F %T %z",
		t:        time.Date(2020, time.July, 24, 23, 14, 15, 0, time.FixedZone("UTC+05:30", (5*60+30)*60)),
		expected: "2020-07-24 23:14:15 +0530",
	},
	{
		format:   "%F %T %Z",
		t:        time.Date(2020, time.July, 24, 23, 14, 15, 0, time.FixedZone("JST", 9*60*60)),
		expected: "2020-07-24 23:14:15 JST",
	},
	{
		format:   "%F %T %Z",
		t:        time.Date(2020, time.July, 24, 23, 14, 15, 0, time.FixedZone("", 9*60*60)),
		expected: "2020-07-24 23:14:15 +0900",
	},
	{
		format:   "%H%%%M%t%S%n%f",
		t:        time.Date(2020, time.January, 1, 1, 2, 3, 450000000, time.UTC),
		expected: "01%02\t03\n450000",
	},
	{
		format:   "%!%.%[%]%|%+%-",
		expected: "!.[]|+-",
	},
	{
		format:   "%",
		expected: "%",
	},
}

func TestFormat(t *testing.T) {
	for _, tc := range formatTestCases {
		t.Run(tc.expected+"/"+tc.format, func(t *testing.T) {
			got, err := timefmt.Format(tc.t, tc.format)
			if tc.formatErr == nil {
				if err != nil {
					t.Fatalf("expected no error but got: %v", err)
				}
				if got != tc.expected {
					t.Errorf("expected: %v, got: %v", tc.expected, got)
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
