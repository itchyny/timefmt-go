package timefmt_test

import (
	"fmt"
	"testing"
	"time"

	timefmt "github.com/itchyny/timefmt-go"
)

var formatTestCases = []struct {
	format   string
	t        time.Time
	expected string
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
		format:   "%y %_y %-y %4y %_4y",
		t:        time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		expected: "20 20 20 0020   20",
	},
	{
		format:   "%y %_y %-y %4y %_4y",
		t:        time.Date(2009, time.January, 1, 0, 0, 0, 0, time.UTC),
		expected: "09  9 9 0009    9",
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
		format:   "%-Y-%-m-%-d",
		t:        time.Date(999, time.September, 8, 0, 0, 0, 0, time.UTC),
		expected: "999-9-8",
	},
	{
		format:   "%_Y-%_m-%_d",
		t:        time.Date(999, time.September, 8, 0, 0, 0, 0, time.UTC),
		expected: " 999- 9- 8",
	},
	{
		format:   "%8Y-%12m-%16d",
		t:        time.Date(2020, time.September, 12, 0, 0, 0, 0, time.UTC),
		expected: "00002020-000000000009-0000000000000012",
	},
	{
		format:   "%Y/%m/%d",
		t:        time.Date(2020, time.October, 9, 0, 0, 0, 0, time.UTC),
		expected: "2020/10/09",
	},
	{
		format:   "%e %-e %_e %4e",
		t:        time.Date(2020, time.January, 9, 0, 0, 0, 0, time.UTC),
		expected: " 9 9  9    9",
	},
	{
		format:   "%B %_B %^B %12B %^12B",
		t:        time.Date(2020, time.October, 1, 0, 0, 0, 0, time.UTC),
		expected: "October October OCTOBER      October      OCTOBER",
	},
	{
		format:   "%b %^b %8b %^8b",
		t:        time.Date(2020, time.September, 1, 0, 0, 0, 0, time.UTC),
		expected: "Sep SEP      Sep      SEP",
	},
	{
		format:   "%h %^h %8h %^8h",
		t:        time.Date(2020, time.November, 1, 0, 0, 0, 0, time.UTC),
		expected: "Nov NOV      Nov      NOV",
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
		format:   "%^A %8A %^8A %^a %8a %^8a %4w %4u %-4w %-4u %_4w %_4u",
		t:        time.Date(2020, time.January, 6, 0, 0, 0, 0, time.UTC),
		expected: "MONDAY   Monday   MONDAY MON      Mon      MON 0001 0001    1    1    1    1",
	},
	{
		format:   "%g %G %a %V %U %W",
		t:        time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		expected: "20 2020 Wed 01 00 00",
	},
	{
		format:   "%g %G %a %V %U %W",
		t:        time.Date(2020, time.January, 4, 0, 0, 0, 0, time.UTC),
		expected: "20 2020 Sat 01 00 00",
	},
	{
		format:   "%g %G %a %V %U %W",
		t:        time.Date(2020, time.January, 5, 0, 0, 0, 0, time.UTC),
		expected: "20 2020 Sun 01 01 00",
	},
	{
		format:   "%g %G %a %V %U %W",
		t:        time.Date(2020, time.January, 6, 0, 0, 0, 0, time.UTC),
		expected: "20 2020 Mon 02 01 01",
	},
	{
		format:   "%-V %-U %-W %_V %_U %_W %4V %4U %4W %_4V %_4U %_4W",
		t:        time.Date(2020, time.January, 6, 0, 0, 0, 0, time.UTC),
		expected: "2 1 1  2  1  1 0002 0001 0001    2    1    1",
	},
	{
		format:   "%g %G %a %V %U %W",
		t:        time.Date(2009, time.December, 31, 0, 0, 0, 0, time.UTC),
		expected: "09 2009 Thu 53 52 52",
	},
	{
		format:   "%g %G %a %V %U %W",
		t:        time.Date(2010, time.January, 1, 0, 0, 0, 0, time.UTC),
		expected: "09 2009 Fri 53 00 00",
	},
	{
		format:   "%g %G %a %V %U %W",
		t:        time.Date(2010, time.January, 2, 0, 0, 0, 0, time.UTC),
		expected: "09 2009 Sat 53 00 00",
	},
	{
		format:   "%g %G %a %V %U %W",
		t:        time.Date(2010, time.January, 3, 0, 0, 0, 0, time.UTC),
		expected: "09 2009 Sun 53 01 00",
	},
	{
		format:   "%g %G %a %V %U %W",
		t:        time.Date(2010, time.January, 4, 0, 0, 0, 0, time.UTC),
		expected: "10 2010 Mon 01 01 01",
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
		expected: "2-9-8 7:6:5.43210",
	},
	{
		format:   "%_y-%_m-%_d %_H:%_M:%_S.%_f",
		t:        time.Date(2002, time.September, 8, 7, 6, 5, 43210000, time.UTC),
		expected: " 2- 9- 8  7: 6: 5. 43210",
	},
	{
		format:   "%4y-%4m-%4d %4H:%4M:%4S.%8f",
		t:        time.Date(2002, time.September, 8, 7, 6, 5, 43210000, time.UTC),
		expected: "0002-0009-0008 0007:0006:0005.00043210",
	},
	{
		format:   "%H:%M:%S.%f",
		t:        time.Date(2020, time.January, 1, 1, 2, 3, 450000000, time.UTC),
		expected: "01:02:03.450000",
	},
	{
		format:   "%I:%M:%S %p",
		t:        time.Date(2020, time.January, 1, 0, 1, 2, 0, time.UTC),
		expected: "12:01:02 AM",
	},
	{
		format:   "%I:%M:%S %p",
		t:        time.Date(2020, time.January, 1, 1, 2, 3, 0, time.UTC),
		expected: "01:02:03 AM",
	},
	{
		format:   "%I:%M:%S %P",
		t:        time.Date(2020, time.January, 1, 12, 13, 14, 0, time.UTC),
		expected: "12:13:14 am",
	},
	{
		format:   "%I:%M:%S %p",
		t:        time.Date(2020, time.January, 1, 13, 14, 15, 0, time.UTC),
		expected: "01:14:15 PM",
	},
	{
		format:   "%I:%M:%S %P",
		t:        time.Date(2020, time.January, 1, 23, 14, 15, 0, time.UTC),
		expected: "11:14:15 pm",
	},
	{
		format:   "%-I:%-M:%-S %-p %-P",
		t:        time.Date(2020, time.January, 1, 13, 2, 3, 0, time.UTC),
		expected: "1:2:3 PM pm",
	},
	{
		format:   "%4I:%4M:%4S %4p %4P",
		t:        time.Date(2020, time.January, 1, 13, 2, 3, 0, time.UTC),
		expected: "0001:0002:0003   PM   pm",
	},
	{
		format:   "%k %-k %_k",
		t:        time.Date(2020, time.January, 1, 9, 0, 0, 0, time.UTC),
		expected: " 9 9  9",
	},
	{
		format:   "%l %-l",
		t:        time.Date(2020, time.January, 1, 20, 0, 0, 0, time.UTC),
		expected: " 8 8",
	},
	{
		format:   "%s %12s",
		t:        time.Date(2020, time.August, 30, 5, 30, 32, 0, time.UTC),
		expected: "1598765432   1598765432",
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
		format:   "%9F %9T",
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
		format:   "%-c",
		t:        time.Date(2020, time.February, 9, 23, 4, 5, 0, time.UTC),
		expected: "Sun Feb  9 23:04:05 2020",
	},
	{
		format:   "%^c",
		t:        time.Date(2020, time.February, 9, 23, 4, 5, 0, time.UTC),
		expected: "SUN FEB  9 23:04:05 2020",
	},
	{
		format:   "%+",
		t:        time.Date(2020, time.February, 9, 23, 4, 5, 0, time.UTC),
		expected: "Sun Feb  9 23:04:05 UTC 2020",
	},
	{
		format:   "%^+",
		t:        time.Date(2020, time.February, 9, 23, 4, 5, 0, time.UTC),
		expected: "SUN FEB  9 23:04:05 UTC 2020",
	},
	{
		format:   "%F %T %z %-z %_4z",
		t:        time.Date(2020, time.July, 24, 23, 14, 15, 0, time.UTC),
		expected: "2020-07-24 23:14:15 +0000 +0 +   0",
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
		format:   "%8Z %8z %_8z %-z",
		t:        time.Date(2020, time.July, 24, 23, 14, 15, 0, time.FixedZone("JST", 9*60*60)),
		expected: "     JST +00000900 +     900 +900",
	},
	{
		format:   "%H%%%M%t%S%n%f",
		t:        time.Date(2020, time.January, 1, 1, 2, 3, 450000000, time.UTC),
		expected: "01%02\t03\n450000",
	},
	{
		format:   "%t,%4t,%n,%4n%4",
		t:        time.Date(2020, time.January, 1, 1, 1, 1, 0, time.UTC),
		expected: "\t,   \t,\n,   \n  %4",
	},
	{
		format:   "%1Y %1y %1C %1g %1G %1m %1V %1U %1d %1j %1H %1M %1S %1f",
		t:        time.Date(2009, time.January, 2, 3, 4, 5, 6000, time.UTC),
		expected: "2009 09 20 09 2009 01 01 00 02 002 03 04 05 6",
	},
	{
		format:   "%2Y %2y %2C %2g %2G %2m %2V %2U %2d %2j %2H %2M %2S %2f",
		t:        time.Date(2009, time.January, 2, 3, 4, 5, 6000, time.UTC),
		expected: "2009 09 20 09 2009 01 01 00 02 002 03 04 05 06",
	},
	{
		format:   "%3Y %3y %3C %3g %3G %3m %3V %3U %3d %3j %3H %3M %3S %3f",
		t:        time.Date(2009, time.January, 2, 3, 4, 5, 6000, time.UTC),
		expected: "2009 009 020 009 2009 001 001 000 002 002 003 004 005 006",
	},
	{
		format:   "%1d %2d %3d %4d %5d %6d %7d %8d %9d %10d",
		t:        time.Date(2020, time.January, 1, 1, 1, 1, 0, time.UTC),
		expected: "01 01 001 0001 00001 000001 0000001 00000001 000000001 0000000001",
	},
	{
		format:   "%!%.%[%]%|%$%-",
		expected: "!.[]|$-",
	},
	{
		format:   "%_",
		expected: "_",
	},
	{
		format:   "%^",
		expected: "^",
	},
	{
		format:   "%",
		expected: "%",
	},
}

func TestFormat(t *testing.T) {
	for _, tc := range formatTestCases {
		t.Run(tc.expected+"/"+tc.format, func(t *testing.T) {
			got := timefmt.Format(tc.t, tc.format)
			if got != tc.expected {
				t.Errorf("expected: %v, got: %v", tc.expected, got)
			}
		})
	}
}

func ExampleFormat() {
	t := time.Date(2020, time.July, 24, 9, 7, 29, 0, time.UTC)
	str := timefmt.Format(t, "%Y-%m-%d %H:%M:%S")
	fmt.Println(str)
	// Output: 2020-07-24 09:07:29
}
