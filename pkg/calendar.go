package calendar

import (
	"strconv"
	"strings"
	"time"

	colors "github.com/vgadzh/go-calendar/pkg/helper"
)

const (
	daysInWeek         = 7
	weekDayLetterCount = 2
	tabSize            = 5
	eol                = "\n"
)

type Calendar struct {
	daysBefore     int
	daysAfter      int
	outputSettings OutputSettings
}
type OutputSettings struct {
	Colors   bool
	Month    bool
	Weekdays bool
}

type day struct {
	now    time.Time
	dt     time.Time
	colors bool
}

func New(weeksBefore, weeksAfter int, os OutputSettings) Calendar {
	weekdayNumber := int(time.Now().Weekday())
	return Calendar{
		daysBefore:     1 - daysInWeek*weeksBefore - weekdayNumber,
		daysAfter:      daysInWeek + daysInWeek*weeksAfter - weekdayNumber,
		outputSettings: os,
	}
}

func (c *Calendar) String() string {
	var output string
	now := time.Now()
	for i := c.daysBefore; i <= c.daysAfter; i++ {
		current := day{
			now:    now,
			dt:     now.AddDate(0, 0, i),
			colors: c.outputSettings.Colors,
		}
		if current.isFirstMonthDay() {
			output += eol
			if c.outputSettings.Month {
				output += eol + current.Month() + eol
			}
			if c.outputSettings.Weekdays {
				output += current.getWeekdayNames() + eol
			}
		} else if current.isFirstWeekday() && i > c.daysBefore {
			output += eol
		} else if i == c.daysBefore {
			if c.outputSettings.Month {
				output += current.Month() + eol
			}
			if c.outputSettings.Weekdays {
				output += current.getWeekdayNames() + eol
			}
		}
		output += current.String()
	}
	return output
}

func (d *day) isFirstWeekday() bool {
	return int(d.dt.Weekday()) == 1
}

func (d *day) isFirstMonthDay() bool {
	return d.dt.Day() == 1
}

func (d *day) isWeekend() bool {
	return d.dt.Weekday() == 0 || d.dt.Weekday() == 6
}

func (d *day) isToday() bool {
	return d.dt.Equal(d.now)
}

func (d *day) isInThePast() bool {
	return d.dt.Before(d.now)
}

func (d *day) Month() string {
	if d.dt.Year() == d.now.Year() {
		return d.dt.Month().String()
	} else {
		return d.dt.Month().String() + " " + strconv.Itoa(d.dt.Year())
	}
}

func (d *day) String() string {
	str := strconv.Itoa(d.dt.Day())
	spaceCount := tabSize - len(str)
	if d.colors {
		if d.isToday() && d.isWeekend() {
			// Today is weekend
			str = colors.GetColoredString(str, colors.Black, colors.OnRed)
		} else if d.isToday() && !d.isWeekend() {
			// Today
			str = colors.GetColoredString(str, colors.Black, colors.OnWhite)
		} else if d.isInThePast() && !d.isWeekend() {
			// Past day
			str = colors.GetColoredString(str, colors.FaintWhite)
		} else if d.isInThePast() && d.isWeekend() {
			// Past day and weekend
			str = colors.GetColoredString(str, colors.FaintRed)
		} else if d.isWeekend() {
			// Future weekend
			str = colors.GetColoredString(str, colors.Red)
		}
	}
	var offset int
	if d.dt.Day() == 1 {
		if int(d.dt.Weekday()) == 0 {
			offset = 6
		} else {
			offset = int(d.dt.Weekday()) - 1
		}
		str = strings.Repeat(" ", offset*tabSize) + str

	}
	str += strings.Repeat(" ", spaceCount)
	return str
}

func (d *day) getWeekdayNames() string {
	var output string
	var weekday string
	for i := 1; i <= 6; i++ {
		weekday = time.Weekday(i).String()[0:weekDayLetterCount]
		if d.colors && i == 6 {
			weekday = colors.GetColoredString(weekday, colors.Red)
		}
		output += weekday + strings.Repeat(" ", tabSize-weekDayLetterCount)
	}
	weekday = time.Weekday(0).String()[0:weekDayLetterCount]
	if d.colors {
		weekday = colors.GetColoredString(weekday, colors.Red)
	}
	output += weekday + strings.Repeat(" ", tabSize-weekDayLetterCount)
	return output
}
