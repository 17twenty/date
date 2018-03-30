package date

import (
	"fmt"
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	Sydney, err := time.LoadLocation("Australia/Sydney")
	if err != nil {
		panic(err)
	}

	for i, test := range []struct {
		want string
		got  interface{}
	}{
		{
			"1970-01-01",
			MustFromString("1970-01-01"),
		},
		{
			"2018-03-30",
			FromTime(time.Date(2018, time.March, 30, 12, 01, 45, 0, time.UTC)),
		},
		{
			"2005-04-21",
			FromTime(MustFromString("2005-04-21").Time()),
		},
		{
			"2001-04-11 00:00:00 +1000 AEST",
			MustFromString("2001-04-11").TimeIn(Sydney),
		},
		{
			"2000-02-25",
			MustFromString("1999-12-25").AddMonths(2),
		},
		{
			"1999-10-25",
			MustFromString("1999-12-25").AddMonths(-2),
		},
		{
			"1999-07-01", // May only has 31 days, so normalises to 1st of July.
			MustFromString("1999-05-31").AddMonths(1),
		},
		{
			"2000-12-25",
			MustFromString("1999-12-25").AddYears(1),
		},
		{
			"2017-03-01", // Feb 2016 has 31 days, but Feb 2017 has 28 days, so normalises.
			MustFromString("2016-02-29").AddYears(1),
		},
	} {
		if gotStr := fmt.Sprintf("%v", test.got); gotStr != test.want {
			t.Errorf("i=%d got=%v want=%v", i, gotStr, test.want)
		}
	}
}

func TestFromStringErr(t *testing.T) {
	if _, err := FromString("not a date"); err == nil {
		t.Error("expected error")
	}
}
