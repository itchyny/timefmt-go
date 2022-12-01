package timefmt_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/itchyny/timefmt-go"
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
		format:   "%Y %1Y %2Y %3Y %4Y %5Y %-Y",
		t:        time.Date(1000, time.January, 1, 0, 0, 0, 0, time.UTC),
		expected: "1000 1000 1000 1000 1000 01000 1000",
	},
	{
		format:   "%Y %1Y %2Y %3Y %4Y %5Y %-Y",
		t:        time.Date(999, time.January, 1, 0, 0, 0, 0, time.UTC),
		expected: "0999 999 999 999 0999 00999 999",
	},
	{
		format:   "%Y %1Y %2Y %3Y %4Y %5Y %-Y",
		t:        time.Date(100, time.January, 1, 0, 0, 0, 0, time.UTC),
		expected: "0100 100 100 100 0100 00100 100",
	},
	{
		format:   "%Y %1Y %2Y %3Y %4Y %5Y %-Y",
		t:        time.Date(99, time.January, 1, 0, 0, 0, 0, time.UTC),
		expected: "0099 99 99 099 0099 00099 99",
	},
	{
		format:   "%Y %1Y %2Y %3Y %4Y %5Y %-Y",
		t:        time.Date(9999, time.January, 1, 0, 0, 0, 0, time.UTC),
		expected: "9999 9999 9999 9999 9999 09999 9999",
	},
	{
		format:   "%Y %1Y %2Y %3Y %4Y %5Y %-Y",
		t:        time.Date(10000, time.January, 1, 0, 0, 0, 0, time.UTC),
		expected: "10000 10000 10000 10000 10000 10000 10000",
	},
	{
		format:   "[%Y]",
		t:        time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		expected: "[2020]",
	},
	{
		format:   "%y %_y %-y %4y %_4y %04y %_04y %0_4y",
		t:        time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		expected: "20 20 20 0020   20 0020 0020   20",
	},
	{
		format:   "%y %_y %-y %4y %_4y %04y %_04y %0_4y",
		t:        time.Date(2009, time.January, 1, 0, 0, 0, 0, time.UTC),
		expected: "09  9 9 0009    9 0009 0009    9",
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
		format:   "%_8Y-%_12m-%_16d",
		t:        time.Date(2020, time.September, 12, 0, 0, 0, 0, time.UTC),
		expected: "    2020-           9-              12",
	},
	{
		format:   "%Y/%m/%d",
		t:        time.Date(2020, time.October, 9, 0, 0, 0, 0, time.UTC),
		expected: "2020/10/09",
	},
	{
		format:   "%e %-e %_e %4e %04e",
		t:        time.Date(2020, time.January, 9, 0, 0, 0, 0, time.UTC),
		expected: " 9 9  9    9 0009",
	},
	{
		format:   "%B %_B %^B %#B %12B %^12B %012B %0^12B",
		t:        time.Date(2020, time.October, 1, 0, 0, 0, 0, time.UTC),
		expected: "October October OCTOBER OCTOBER      October      OCTOBER 00000October 00000OCTOBER",
	},
	{
		format:   "%b %^b %#b %8b %^8b %08b %^08b %#08b",
		t:        time.Date(2020, time.September, 1, 0, 0, 0, 0, time.UTC),
		expected: "Sep SEP SEP      Sep      SEP 00000Sep 00000SEP 00000SEP",
	},
	{
		format:   "%h %^h %#h %8h %^8h %08h %^08h",
		t:        time.Date(2020, time.November, 1, 0, 0, 0, 0, time.UTC),
		expected: "Nov NOV NOV      Nov      NOV 00000Nov 00000NOV",
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
		format:   "%^A %#A %8A %^8A %08A %^08A %^a %#a %8a %^8a %08a %^08a",
		t:        time.Date(2020, time.January, 6, 0, 0, 0, 0, time.UTC),
		expected: "MONDAY MONDAY   Monday   MONDAY 00Monday 00MONDAY MON MON      Mon      MON 00000Mon 00000MON",
	},
	{
		format:   "%4w %4u %-4w %-4u %_4w %_4u %04w %04u",
		t:        time.Date(2020, time.January, 6, 0, 0, 0, 0, time.UTC),
		expected: "0001 0001    1    1    1    1 0001 0001",
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
		format:   "%-V %-U %-W %_V %_U %_W %4V %4U %4W %_4V %_4U %_4W %04V %04U %04W",
		t:        time.Date(2020, time.January, 6, 0, 0, 0, 0, time.UTC),
		expected: "2 1 1  2  1  1 0002 0001 0001    2    1    1 0002 0001 0001",
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
		t:        time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC),
		expected: "2020-010-10",
	},
	{
		format:   "%Y-%j-%-j",
		t:        time.Date(2020, time.February, 2, 0, 0, 0, 0, time.UTC),
		expected: "2020-033-33",
	},
	{
		format:   "%Y-%j-%-j",
		t:        time.Date(2020, time.April, 8, 0, 0, 0, 0, time.UTC),
		expected: "2020-099-99",
	},
	{
		format:   "%Y-%j-%-j",
		t:        time.Date(2020, time.April, 9, 0, 0, 0, 0, time.UTC),
		expected: "2020-100-100",
	},
	{
		format:   "%Y-%j-%-j",
		t:        time.Date(2020, time.April, 10, 0, 0, 0, 0, time.UTC),
		expected: "2020-101-101",
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
		format:   "%04y-%04m-%04d %04H:%04M:%04S.%08f",
		t:        time.Date(2002, time.September, 8, 7, 6, 5, 43210000, time.UTC),
		expected: "0002-0009-0008 0007:0006:0005.00043210",
	},
	{
		format:   "%H:%M:%S.%f",
		t:        time.Date(2020, time.January, 1, 1, 2, 3, 450000000, time.UTC),
		expected: "01:02:03.450000",
	},
	{
		format:   "%I:%M:%S %p %P",
		t:        time.Date(2020, time.January, 1, 0, 2, 3, 0, time.UTC),
		expected: "12:02:03 AM am",
	},
	{
		format:   "%I:%M:%S %p %P",
		t:        time.Date(2020, time.January, 1, 1, 2, 3, 0, time.UTC),
		expected: "01:02:03 AM am",
	},
	{
		format:   "%I:%M:%S %p %P",
		t:        time.Date(2020, time.January, 1, 11, 2, 3, 0, time.UTC),
		expected: "11:02:03 AM am",
	},
	{
		format:   "%I:%M:%S %p %P",
		t:        time.Date(2020, time.January, 1, 12, 2, 3, 0, time.UTC),
		expected: "12:02:03 PM pm",
	},
	{
		format:   "%I:%M:%S %p %P",
		t:        time.Date(2020, time.January, 1, 13, 2, 3, 0, time.UTC),
		expected: "01:02:03 PM pm",
	},
	{
		format:   "%I:%M:%S %p %P",
		t:        time.Date(2020, time.January, 1, 23, 2, 3, 0, time.UTC),
		expected: "11:02:03 PM pm",
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
		format:   "%_I:%_M:%_S %_p %_P",
		t:        time.Date(2020, time.January, 1, 13, 2, 3, 0, time.UTC),
		expected: " 1: 2: 3 PM pm",
	},
	{
		format:   "%_4I:%_4M:%_4S %_4p %_4P",
		t:        time.Date(2020, time.January, 1, 13, 2, 3, 0, time.UTC),
		expected: "   1:   2:   3   PM   pm",
	},
	{
		format:   "%-4I:%-4M:%-4S %-4p %-4P",
		t:        time.Date(2020, time.January, 1, 13, 2, 3, 0, time.UTC),
		expected: "   1:   2:   3   PM   pm",
	},
	{
		format:   "%04I:%04M:%04S %04p %04P",
		t:        time.Date(2020, time.January, 1, 13, 2, 3, 0, time.UTC),
		expected: "0001:0002:0003 00PM 00pm",
	},
	{
		format:   "%p %P %^p %^P %#p %#P %#^p %_4p %_4P %04p %04P %^04p %^04P",
		t:        time.Date(2020, time.January, 1, 11, 2, 3, 0, time.UTC),
		expected: "AM am AM AM am AM am   AM   am 00AM 00am 00AM 00AM",
	},
	{
		format:   "%p %P %^p %^P %#p %#P %#^p %_4p %_4P %04p %04P %^04p %^04P",
		t:        time.Date(2020, time.January, 1, 23, 2, 3, 0, time.UTC),
		expected: "PM pm PM PM pm PM pm   PM   pm 00PM 00pm 00PM 00PM",
	},
	{
		format:   "%k %-k %_k %4k %0k %04k",
		t:        time.Date(2020, time.January, 1, 9, 0, 0, 0, time.UTC),
		expected: " 9 9  9    9 09 0009",
	},
	{
		format:   "%l %-l %_l %4l %0l %04l",
		t:        time.Date(2020, time.January, 1, 20, 0, 0, 0, time.UTC),
		expected: " 8 8  8    8 08 0008",
	},
	{
		format:   "%s %12s %_12s %012s",
		t:        time.Date(2020, time.August, 30, 5, 30, 32, 0, time.UTC),
		expected: "1598765432   1598765432   1598765432 001598765432",
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
		format:   "%#c",
		t:        time.Date(2020, time.February, 9, 23, 4, 5, 0, time.UTC),
		expected: "Sun Feb  9 23:04:05 2020",
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
		format:   "%#+",
		t:        time.Date(2020, time.February, 9, 23, 4, 5, 0, time.UTC),
		expected: "Sun Feb  9 23:04:05 UTC 2020",
	},
	{
		format:   "%F %T %z %-z %_4z %04z",
		t:        time.Date(2020, time.July, 24, 23, 14, 15, 0, time.UTC),
		expected: "2020-07-24 23:14:15 +0000 +000  +000 +0000",
	},
	{
		format:   "%F %T %z",
		t:        time.Date(2020, time.July, 24, 23, 14, 15, 0, time.FixedZone("", -8*60*60)),
		expected: "2020-07-24 23:14:15 -0800",
	},
	{
		format:   "%F %T %z",
		t:        time.Date(2020, time.July, 24, 23, 14, 15, 0, time.FixedZone("", 9*60*60)),
		expected: "2020-07-24 23:14:15 +0900",
	},
	{
		format:   "%F %T %z",
		t:        time.Date(2020, time.July, 24, 23, 14, 15, 0, time.FixedZone("", (5*60+30)*60)),
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
		format:   "%Z %^Z %#Z %#^Z %^#Z",
		t:        time.Date(2020, time.July, 24, 23, 14, 15, 0, time.FixedZone("JST", 9*60*60)),
		expected: "JST JST jst jst jst",
	},
	{
		format:   "%8Z %08Z %8z %_8z %-z %08z %2z %3z %4z %5z %6z %6:z %7:z %:%Z",
		t:        time.Date(2020, time.July, 24, 23, 14, 15, 0, time.FixedZone("JST", 9*60*60)),
		expected: "     JST 00000JST +0000900     +900 +900 +0000900 +0900 +0900 +0900 +0900 +00900 +09:00 +009:00 %:JST",
	},
	{
		format:   "%8Z %08Z %8z %_8z %-z %08z %4z %5z %6z %6:z %7:z",
		t:        time.Date(2020, time.July, 24, 23, 14, 15, 0, time.FixedZone("HAST", -10*60*60)),
		expected: "    HAST 0000HAST -0001000    -1000 -1000 -0001000 -1000 -1000 -01000 -10:00 -010:00",
	},
	{
		format:   "%8z %_8z %-z %08z %4z %5z %6z %6:z %7:z",
		t:        time.Date(2020, time.July, 24, 23, 14, 15, 0, time.FixedZone("", -(5*60+30)*60)),
		expected: "-0000530     -530 -530 -0000530 -0530 -0530 -00530 -05:30 -005:30",
	},
	{
		format:   "%8z %_8z %-z %08z %4z %5z %6z %6:z %7:z",
		t:        time.Date(2020, time.July, 24, 23, 14, 15, 0, time.FixedZone("", 30*60)),
		expected: "+0000030     +030 +030 +0000030 +0030 +0030 +00030 +00:30 +000:30",
	},
	{
		format:   "%:z %::z %:::z %::::z %:Z %-:z %_::z %8z %08:z %_8:z %_8::z %8:::z %-:::z %:-z %:",
		t:        time.Date(2020, time.July, 24, 23, 14, 15, 0, time.UTC),
		expected: "+00:00 +00:00:00 +00 %::::z %:Z +0:00  +0:00:00 +0000000 +0000:00    +0:00  +0:00:00 +0000000 +0 %:-z %:",
	},
	{
		format:   "%:z %::z %:::z %::::z %:Z %-:z %_::z %8z %08:z %_8:z %_8::z %8:::z %-:::z %:-z %:",
		t:        time.Date(2020, time.July, 24, 23, 14, 15, 0, time.FixedZone("JST", 9*60*60)),
		expected: "+09:00 +09:00:00 +09 %::::z %:Z +9:00  +9:00:00 +0000900 +0009:00    +9:00  +9:00:00 +0000009 +9 %:-z %:",
	},
	{
		format:   "%:z %::z %:::z %::::z %:Z %-:z %_::z %8z %08:z %_8:z %_8::z %8:::z %-:::z %:-z %:",
		t:        time.Date(2020, time.July, 24, 23, 14, 15, 0, time.FixedZone("HAST", -(10*60+30)*60)),
		expected: "-10:30 -10:30:00 -10:30 %::::z %:Z -10:30 -10:30:00 -0001030 -0010:30   -10:30 -10:30:00 -0010:30 -10:30 %:-z %:",
	},
	{
		format:   "%:z %::z %:::z %::::z %:Z %-:z %_::z %8z %08:z %_8:z %_8::z %8:::z %-:::z %:-z %:",
		t:        time.Date(2020, time.July, 24, 23, 14, 15, 0, time.FixedZone("", 60*60+2)),
		expected: "+01:00 +01:00:02 +01:00:02 %::::z %:Z +1:00  +1:00:02 +0000100 +0001:00    +1:00  +1:00:02 +01:00:02 +1:00:02 %:-z %:",
	},
	{
		format:   "%:z %::z %:::z %::::z %:Z %-:z %_::z %8z %08:z %_8:z %_8::z %8:::z %-:::z %:-z %:",
		t:        time.Date(2020, time.July, 24, 23, 14, 15, 0, time.FixedZone("", -9*60*60-62)),
		expected: "-09:01 -09:01:02 -09:01:02 %::::z %:Z -9:01  -9:01:02 -0000901 -0009:01    -9:01  -9:01:02 -09:01:02 -9:01:02 %:-z %:",
	},
	{
		format:   "%:: %6: %z",
		t:        time.Date(2020, time.July, 24, 23, 14, 15, 0, time.FixedZone("", 9*60*60)),
		expected: "%::    %6: +0900",
	},
	{
		format:   "%H%%%M%t%S%n%f",
		t:        time.Date(2020, time.January, 1, 1, 2, 3, 450000000, time.UTC),
		expected: "01%02\t03\n450000",
	},
	{
		format:   "%t,%4t,%04t,%n,%4n,%04n,%4",
		t:        time.Date(2020, time.January, 1, 1, 1, 1, 0, time.UTC),
		expected: "\t,   \t,000\t,\n,   \n,000\n,  %4",
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
		format:   "%4Y %4y %4C %4g %4G %4m %4V %4U %4d %4j %4H %4M %4S %4f",
		t:        time.Date(2009, time.January, 2, 3, 4, 5, 60000000, time.UTC),
		expected: "2009 0009 0020 0009 2009 0001 0001 0000 0002 0002 0003 0004 0005 60000",
	},
	{
		format:   "%1d %2d %3d %4d %5d %6d %7d %8d %9d %10d",
		t:        time.Date(2020, time.January, 1, 1, 1, 1, 0, time.UTC),
		expected: "01 01 001 0001 00001 000001 0000001 00000001 000000001 0000000001",
	},
	{
		format:   "%10000Y",
		t:        time.Date(2020, time.January, 1, 1, 1, 1, 0, time.UTC),
		expected: fmt.Sprintf("%01024d", 2020),
	},
	{
		format:   "%922337203685477580Y",
		t:        time.Date(2020, time.January, 1, 1, 1, 1, 0, time.UTC),
		expected: fmt.Sprintf("%01024d", 2020),
	},
	{
		format:   "%9223372036854775809Y",
		t:        time.Date(2020, time.January, 1, 1, 1, 1, 0, time.UTC),
		expected: fmt.Sprintf("%01024d", 2020),
	},
	{
		format:   "%18446744073709551630Y",
		t:        time.Date(2020, time.January, 1, 1, 1, 1, 0, time.UTC),
		expected: fmt.Sprintf("%01024d", 2020),
	},
	{
		format:   "%!%.%[%]%|%$%-",
		expected: "%!%.%[%]%|%$%-",
	},
	{
		format:   "%4_",
		expected: " %4_",
	},
	{
		format:   "%09_",
		expected: "00000%09_",
	},
	{
		format:   "%^",
		expected: "%^",
	},
	{
		format:   "%#",
		expected: "%#",
	},
	{
		format:   "%0",
		expected: "%0",
	},
	{
		format:   "%4",
		expected: "  %4",
	},
	{
		format:   "%-9",
		expected: "      %-9",
	},
	{
		format:   "%-9^",
		expected: "     %-9^",
	},
	{
		format:   "%-09^",
		expected: "0000%-09^",
	},
	{
		format:   "%06",
		expected: "000%06",
	},
	{
		format:   "%6J",
		expected: "   %6J",
	},
	{
		format:   "%",
		expected: "%",
	},
}

func TestFormat(t *testing.T) {
	for _, tc := range formatTestCases {
		var name string
		if len(tc.expected) < 1000 {
			name = tc.expected + "/" + tc.format
		} else {
			name = strings.ReplaceAll(tc.expected+"/"+tc.format, strings.Repeat("0", 30), "0.")
		}
		t.Run(name, func(t *testing.T) {
			got := timefmt.Format(tc.t, tc.format)
			if got != tc.expected {
				t.Error(diff(tc.expected, got))
			}
		})
	}
}

func TestAppendFormat(t *testing.T) {
	tm := time.Date(2020, time.January, 2, 3, 4, 5, 0, time.UTC)
	buf := timefmt.AppendFormat(make([]byte, 0, 64), tm, "%c")
	if got, expected := string(buf), "Thu Jan  2 03:04:05 2020"; got != expected {
		t.Errorf("expected: %q, got: %q", expected, got)
	}
}

func ExampleFormat() {
	t := time.Date(2020, time.July, 24, 9, 7, 29, 0, time.UTC)
	str := timefmt.Format(t, "%Y-%m-%d %H:%M:%S")
	fmt.Println(str)
	// Output: 2020-07-24 09:07:29
}

func ExampleAppendFormat() {
	t := time.Date(2020, time.July, 24, 9, 7, 29, 0, time.UTC)
	buf := make([]byte, 0, 64)
	buf = append(buf, '(')
	buf = timefmt.AppendFormat(buf, t, "%Y-%m-%d %H:%M:%S")
	buf = append(buf, ')')
	fmt.Println(string(buf))
	// Output: (2020-07-24 09:07:29)
}
