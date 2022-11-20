package timefmt_test

import (
	"errors"
	"fmt"
	"log"
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
		source: "20",
		format: "%Y",
		t:      time.Date(20, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "2",
		format: "%Y",
		t:      time.Date(2, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "20",
		format: "%Y",
		t:      time.Date(20, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source:   "20xxx",
		format:   "%Y",
		parseErr: errors.New(`unconverted string: "xxx"`),
	},
	{
		source:   "a",
		format:   "%Y",
		parseErr: errors.New("cannot parse %Y"),
	},
	{
		source: "20",
		format: "%C",
		t:      time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "1758",
		format: "%C%y",
		t:      time.Date(1758, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source:   "xx",
		format:   "%C",
		parseErr: errors.New("cannot parse %C"),
	},
	{
		source: "68",
		format: "%y",
		t:      time.Date(2068, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "69",
		format: "%y",
		t:      time.Date(1969, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source:   "xx",
		format:   "%y",
		parseErr: errors.New("cannot parse %y"),
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
		source: "9-9-9",
		format: "%y-%m-%d",
		t:      time.Date(2009, time.September, 9, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "2020 02  9",
		format: "%Y %m %e",
		t:      time.Date(2020, time.February, 9, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "Jan",
		format: "%b",
		t:      time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source:   "Ja",
		format:   "%b",
		parseErr: errors.New("cannot parse %b"),
	},
	{
		source: "Jul",
		format: "%b",
		t:      time.Date(1900, time.July, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "Sep",
		format: "%b",
		t:      time.Date(1900, time.September, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "September",
		format: "%B",
		t:      time.Date(1900, time.September, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source:   "Sep",
		format:   "%B",
		parseErr: errors.New("cannot parse %B"),
	},
	{
		source: "Sep",
		format: "%h",
		t:      time.Date(1900, time.September, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "100",
		format: "%j",
		t:      time.Date(1900, time.April, 10, 0, 0, 0, 0, time.UTC),
	},
	{
		source:   ".10",
		format:   "%j",
		parseErr: errors.New("cannot parse %j"),
	},
	{
		source: "20203",
		format: "%Y%j",
		t:      time.Date(2020, time.January, 3, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "2020366",
		format: "%Y%j",
		t:      time.Date(2020, time.December, 31, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "2020-9-33",
		format: "%Y-%m-%j",
		t:      time.Date(2020, time.February, 2, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "MAY",
		format: "%b",
		t:      time.Date(1900, time.May, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "SATURDAY",
		format: "%A",
		t:      time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "[sunday]",
		format: "[%A]",
		t:      time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "[Mon]",
		format: "[%a]",
		t:      time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source:   "Teu",
		format:   "%a",
		parseErr: errors.New("cannot parse %a"),
	},
	{
		source: "mondaymon1",
		format: "%A%a%w",
		t:      time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source:   "mooday",
		format:   "%A",
		parseErr: errors.New("cannot parse %A"),
	},
	{
		source: "0",
		format: "%w",
		t:      time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "6",
		format: "%w",
		t:      time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source:   "7",
		format:   "%w",
		parseErr: errors.New("cannot parse %w"),
	},
	{
		source:   "0",
		format:   "%u",
		parseErr: errors.New("cannot parse %u"),
	},
	{
		source: "1",
		format: "%u",
		t:      time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source: "7",
		format: "%u",
		t:      time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source:   "8",
		format:   "%u",
		parseErr: errors.New("cannot parse %u"),
	},
	{
		source: "20 01",
		format: "%g %V",
		t:      time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source:   "xx",
		format:   "%g",
		parseErr: errors.New("cannot parse %g"),
	},
	{
		source:   "xx",
		format:   "%V",
		parseErr: errors.New("cannot parse %V"),
	},
	{
		source: "2009 01 00 00",
		format: "%G %V %U %W",
		t:      time.Date(2009, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		source:   "xxxx",
		format:   "%G",
		parseErr: errors.New("cannot parse %G"),
	},
	{
		source: "2020-09-08 07:06:05",
		format: "%Y-%m-%d %H:%M:%S",
		t:      time.Date(2020, time.September, 8, 7, 6, 5, 0, time.UTC),
	},
	{
		source: "1:2:3.45",
		format: "%H:%M:%S.%f",
		t:      time.Date(1900, time.January, 1, 1, 2, 3, 450000000, time.UTC),
	},
	{
		source: "1213145678912",
		format: "%H%M%S%f%d",
		t:      time.Date(1900, time.January, 2, 12, 13, 14, 567891000, time.UTC),
	},
	{
		source:   "xx",
		format:   "%H",
		parseErr: errors.New("cannot parse %H"),
	},
	{
		source:   "xx",
		format:   "%M",
		parseErr: errors.New("cannot parse %M"),
	},
	{
		source:   "xx",
		format:   "%S",
		parseErr: errors.New("cannot parse %S"),
	},
	{
		source:   "1:2:3.",
		format:   "%H:%M:%S.%f",
		parseErr: errors.New("cannot parse %f"),
	},
	{
		source: "12:13:14 AM",
		format: "%I:%M:%S %p",
		t:      time.Date(1900, time.January, 1, 0, 13, 14, 0, time.UTC),
	},
	{
		source: "01:14:15PM",
		format: "%I:%M:%S%p",
		t:      time.Date(1900, time.January, 1, 13, 14, 15, 0, time.UTC),
	},
	{
		source: "PM 11:14:15",
		format: "%p %I:%M:%S",
		t:      time.Date(1900, time.January, 1, 23, 14, 15, 0, time.UTC),
	},
	{
		source: "12:13:14 PM",
		format: "%I:%M:%S %p",
		t:      time.Date(1900, time.January, 1, 12, 13, 14, 0, time.UTC),
	},
	{
		source:   "xx",
		format:   "%I",
		parseErr: errors.New("cannot parse %I"),
	},
	{
		source: "am  9:10:11",
		format: "%P %k:%M:%S",
		t:      time.Date(1900, time.January, 1, 9, 10, 11, 0, time.UTC),
	},
	{
		source: " 9:10:11 pm",
		format: "%l:%M:%S %P",
		t:      time.Date(1900, time.January, 1, 21, 10, 11, 0, time.UTC),
	},
	{
		source: "1598765432",
		format: "%s",
		t:      time.Date(2020, time.August, 30, 5, 30, 32, 0, time.UTC),
	},
	{
		source:   ".",
		format:   "%s",
		parseErr: errors.New("cannot parse %s"),
	},
	{
		source: "23:14",
		format: "%R",
		t:      time.Date(1900, time.January, 1, 23, 14, 0, 0, time.UTC),
	},
	{
		source:   "23:",
		format:   "%R",
		parseErr: errors.New("cannot parse %M"),
	},
	{
		source: "3:14:15 PM",
		format: "%r",
		t:      time.Date(1900, time.January, 1, 15, 14, 15, 0, time.UTC),
	},
	{
		source:   "3:1415 PM",
		format:   "%r",
		parseErr: errors.New("expected ':'"),
	},
	{
		source:   "3:14:15PM",
		format:   "%r",
		parseErr: errors.New("expected ' '"),
	},
	{
		source: "2/20/21 23:14:15",
		format: "%D %T",
		t:      time.Date(2021, time.February, 20, 23, 14, 15, 0, time.UTC),
	},
	{
		source: "02/09/20 23:14:15",
		format: "%x %X",
		t:      time.Date(2020, time.February, 9, 23, 14, 15, 0, time.UTC),
	},
	{
		source: "2020-02-09 \t\n\v\f\r 23:14:15",
		format: "%F%t%T",
		t:      time.Date(2020, time.February, 9, 23, 14, 15, 0, time.UTC),
	},
	{
		source:   "2020-02-0923:14:15",
		format:   "%F%t%T",
		parseErr: errors.New("expected a space for %t"),
	},
	{
		source: " 9-Jul-2020 23:14:15",
		format: "%v %X",
		t:      time.Date(2020, time.July, 9, 23, 14, 15, 0, time.UTC),
	},
	{
		source: "Sun Feb  9 23:14:15 2020",
		format: "%c",
		t:      time.Date(2020, time.February, 9, 23, 14, 15, 0, time.UTC),
	},
	{
		source: "Sun Feb  9 23:14:15 UTC 2020",
		format: "%+",
		t:      time.Date(2020, time.February, 9, 23, 14, 15, 0, time.UTC),
	},
	{
		source: "2020-07-24 23:14:15 +0000",
		format: "%F %T %z",
		t:      time.Date(2020, time.July, 24, 23, 14, 15, 0, time.FixedZone("", 0)),
	},
	{
		source: "2020-07-24T23:14:15Z",
		format: "%FT%T%z",
		t:      time.Date(2020, time.July, 24, 23, 14, 15, 0, time.UTC),
	},
	{
		source: "2020-07-24 23:14:15 -0800",
		format: "%F %T %z",
		t:      time.Date(2020, time.July, 24, 23, 14, 15, 0, time.FixedZone("", -8*60*60)),
	},
	{
		source: "2020-07-24 23:14:15 +0900",
		format: "%F %T %z",
		t:      time.Date(2020, time.July, 24, 23, 14, 15, 0, time.FixedZone("", 9*60*60)),
	},
	{
		source: "2020-07-24 23:14:15 +0530",
		format: "%F %T %z",
		t:      time.Date(2020, time.July, 24, 23, 14, 15, 0, time.FixedZone("", (5*60+30)*60)),
	},
	{
		source: "2020-07-24 23:14:15 +04:30",
		format: "%F %T %z",
		t:      time.Date(2020, time.July, 24, 23, 14, 15, 0, time.FixedZone("", (4*60+30)*60)),
	},
	{
		source: "2020-07-24 23:14:15 +05:43:21",
		format: "%F %T %z",
		t:      time.Date(2020, time.July, 24, 23, 14, 15, 0, time.FixedZone("", (5*60+43)*60+21)),
	},
	{
		source: "2020-07-24 23:14:15 +05:43zzz",
		format: "%F %T %zzzz",
		t:      time.Date(2020, time.July, 24, 23, 14, 15, 0, time.FixedZone("", (5*60+43)*60)),
	},
	{
		source: "2020-07-24 23:14:15 +05:43:",
		format: "%F %T %z:",
		t:      time.Date(2020, time.July, 24, 23, 14, 15, 0, time.FixedZone("", (5*60+43)*60)),
	},
	{
		source: "2020-07-24 23:14:15 +05:43:0",
		format: "%F %T %z:0",
		t:      time.Date(2020, time.July, 24, 23, 14, 15, 0, time.FixedZone("", (5*60+43)*60)),
	},
	{
		source: "2020-07-24 23:14:15 +05",
		format: "%F %T %z",
		t:      time.Date(2020, time.July, 24, 23, 14, 15, 0, time.FixedZone("", 5*60*60)),
	},
	{
		source: "2020-07-24T23:14:15+05Z",
		format: "%FT%T%z%z",
		t:      time.Date(2020, time.July, 24, 23, 14, 15, 0, time.UTC),
	},
	{
		source:   "2020-07-24 23:14:15 ",
		format:   "%F %T %z",
		parseErr: errors.New("cannot parse %z"),
	},
	{
		source:   "2020-07-24 23:14:15 +",
		format:   "%F %T %z",
		parseErr: errors.New("cannot parse %z"),
	},
	{
		source:   "2020-07-24 23:14:15 +0",
		format:   "%F %T %z",
		parseErr: errors.New("cannot parse %z"),
	},
	{
		source:   "2020-07-24 23:14:15 +053",
		format:   "%F %T %z",
		parseErr: errors.New(`unconverted string: "3"`),
	},
	{
		source:   "2020-07-24 23:14:15 +04:3",
		format:   "%F %T %z",
		parseErr: errors.New("cannot parse %z"),
	},
	{
		source:   "2020-07-24 23:14:15 +04:30:",
		format:   "%F %T %z",
		parseErr: errors.New(`unconverted string: ":"`),
	},
	{
		source:   "2020-07-24 23:14:15 +04:30:0",
		format:   "%F %T %z",
		parseErr: errors.New(`unconverted string: ":0"`),
	},
	{
		source:   "2020-07-24 23:14:15 +04:3:00",
		format:   "%F %T %z",
		parseErr: errors.New("cannot parse %z"),
	},
	{
		source:   "2020-07-24 23:14:15 +0430:10",
		format:   "%F %T %z",
		parseErr: errors.New(`unconverted string: ":10"`),
	},
	{
		source:   "2020-07-24 23:14:15 +04:3010",
		format:   "%F %T %z",
		parseErr: errors.New(`unconverted string: "10"`),
	},
	{
		source:   "2020-07-24 23:14:15 +0:30",
		format:   "%F %T %z",
		parseErr: errors.New("cannot parse %z"),
	},
	{
		source:   "2020-07-24 23:14:15 +003a",
		format:   "%F %T %z",
		parseErr: errors.New(`unconverted string: "3a"`),
	},
	{
		source:   "2020-07-24 23:14:15 $0000",
		format:   "%F %T %z",
		parseErr: errors.New("cannot parse %z"),
	},
	{
		source:   "2020-07-24 23:14:15 +05:43:2a",
		format:   "%F %T %z",
		parseErr: errors.New(`unconverted string: ":2a"`),
	},
	{
		source: "2020-07-24 23:14:15 +05:30%",
		format: "%F %T %:z%%",
		t:      time.Date(2020, time.July, 24, 23, 14, 15, 0, time.FixedZone("", (5*60+30)*60)),
	},
	{
		source: "2020-07-24 23:14:15 Z",
		format: "%F %T %:z",
		t:      time.Date(2020, time.July, 24, 23, 14, 15, 0, time.UTC),
	},
	{
		source:   "2020-07-24 23:14:15 +05",
		format:   "%F %T %:z",
		parseErr: errors.New("expected ':' for %:z"),
	},
	{
		source:   "2020-07-24 23:14:15 +05-30",
		format:   "%F %T %:z",
		parseErr: errors.New("expected ':' for %:z"),
	},
	{
		source:   "2020-07-24 23:14:15 +0530",
		format:   "%F %T %:z",
		parseErr: errors.New(`expected ':' for %:z`),
	},
	{
		source:   "2020-07-24 23:14:15 *05:30",
		format:   "%F %T %:z",
		parseErr: errors.New("cannot parse %:z"),
	},
	{
		source:   "2020-07-24 23:14:15 +0x:30",
		format:   "%F %T %:z",
		parseErr: errors.New("cannot parse %:z"),
	},
	{
		source:   "2020-07-24 23:14:15 +00:3x",
		format:   "%F %T %:z",
		parseErr: errors.New("cannot parse %:z"),
	},
	{
		source:   "2020-07-24 23:14:15 ",
		format:   "%F %T %:",
		parseErr: errors.New(`expected 'z' after "%:"`),
	},
	{
		source:   "2020-07-24 23:14:15 ",
		format:   "%F %T %:H",
		parseErr: errors.New(`expected 'z' after "%:"`),
	},
	{
		source: "2020-07-24 23:14:15 +05:30:10",
		format: "%F %T %::z",
		t:      time.Date(2020, time.July, 24, 23, 14, 15, 0, time.FixedZone("", (5*60+30)*60+10)),
	},
	{
		source: "2020-07-24 23:14:15 -05:30:10",
		format: "%F %T %::z",
		t:      time.Date(2020, time.July, 24, 23, 14, 15, 0, time.FixedZone("", -((5*60+30)*60+10))),
	},
	{
		source: "2020-07-24 23:14:15 ::-05:30::",
		format: "%F %T ::%:z::",
		t:      time.Date(2020, time.July, 24, 23, 14, 15, 0, time.FixedZone("", -(5*60+30)*60)),
	},
	{
		source: "2020-07-24 23:14:15 -05:30:10 -04:20 +0300",
		format: "%F %T %::z %:z %z",
		t:      time.Date(2020, time.July, 24, 23, 14, 15, 0, time.FixedZone("", 3*60*60)),
	},
	{
		source:   "2020-07-24 23:14:15 +05",
		format:   "%F %T %::z",
		parseErr: errors.New("expected ':' for %::z"),
	},
	{
		source:   "2020-07-24 23:14:15 +05:30:0",
		format:   "%F %T %::z",
		parseErr: errors.New("cannot parse %::z"),
	},
	{
		source:   "2020-07-24 23:14:15 +05:30:0x",
		format:   "%F %T %::z",
		parseErr: errors.New("cannot parse %::z"),
	},
	{
		source:   "2020-07-24 23:14:15 /05:30:00",
		format:   "%F %T %::z",
		parseErr: errors.New("cannot parse %::z"),
	},
	{
		source:   "2020-07-24 23:14:15 +05300000",
		format:   "%F %T %::z",
		parseErr: errors.New("expected ':' for %::z"),
	},
	{
		source:   "2020-07-24 23:14:15 +05-30:00",
		format:   "%F %T %::z",
		parseErr: errors.New("expected ':' for %::z"),
	},
	{
		source:   "2020-07-24 23:14:15 +05:30-00",
		format:   "%F %T %::z",
		parseErr: errors.New("expected ':' for %::z"),
	},
	{
		source:   "2020-07-24 23:14:15 +05:30",
		format:   "%F %T %::z",
		parseErr: errors.New("expected ':' for %::z"),
	},
	{
		source:   "2020-07-24 23:14:15 ",
		format:   "%F %T %::",
		parseErr: errors.New(`expected 'z' after "%::"`),
	},
	{
		source:   "2020-07-24 23:14:15 ",
		format:   "%F %T %::Z",
		parseErr: errors.New(`expected 'z' after "%::"`),
	},
	{
		source: "2020-07-24 23:14:15 UTC",
		format: "%F %T %Z",
		t:      time.Date(2020, time.July, 24, 23, 14, 15, 0, time.FixedZone("UTC", 0)),
	},
	{
		source:   "X",
		format:   "%Z",
		parseErr: errors.New(`cannot parse "X" with "%Z"`),
	},
	{
		source: "2020-07-24 23:14:15 +0530 (AAA)",
		format: "%F %T %z (%Z)",
		t:      time.Date(2020, time.July, 24, 23, 14, 15, 0, time.FixedZone("AAA", (5*60+30)*60)),
	},
	{
		source: "2020-07-24 23:14:15 (AAA) +0530",
		format: "%F %T (%Z) %z",
		t:      time.Date(2020, time.July, 24, 23, 14, 15, 0, time.FixedZone("AAA", (5*60+30)*60)),
	},
	{
		source: "01%02\t03\n450000",
		format: "%H%%%M%t%S%n%f",
		t:      time.Date(1900, time.January, 1, 1, 2, 3, 450000000, time.UTC),
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
	{
		format:   "%",
		parseErr: errors.New("stray %"),
	},
	{
		source:   "",
		format:   "%%",
		parseErr: errors.New("expected '%'"),
	},
	{
		source:   "",
		format:   "x",
		parseErr: errors.New("expected 'x'"),
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
				name, offset := tc.t.Zone()
				gotName, gotOffset := got.Zone()
				if name != gotName || offset != gotOffset {
					t.Errorf("expected zone: name = %s, offset = %d, got zone: name = %s, offset = %d",
						name, offset,
						gotName, gotOffset,
					)
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

func ExampleParse() {
	str := "2020-07-24 09:07:29"
	t, err := timefmt.Parse(str, "%Y-%m-%d %H:%M:%S")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t)
	// Output: 2020-07-24 09:07:29 +0000 UTC
}

func ExampleParseInLocation() {
	loc := time.FixedZone("JST", 9*60*60)
	str := "2020-07-24 09:07:29"
	t, err := timefmt.ParseInLocation(str, "%Y-%m-%d %H:%M:%S", loc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t)
	// Output: 2020-07-24 09:07:29 +0900 JST
}
