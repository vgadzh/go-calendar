package calendar

import "time"

type Calendar struct {
	now       time.Time
	curr      time.Time
	startDate time.Time
	endDate   time.Time
}

func NewCalendar(now time.Time, weeksBefore, weeksAfter int) *Calendar {
	var weekday int
	if now.Weekday() == time.Sunday {
		weekday = 7
	} else {
		weekday = int(now.Weekday())
	}
	currentMonday := now.AddDate(0, 0, -(weekday - 1))

	minWeekIndex := 0
	if weeksBefore > 0 {
		minWeekIndex = -weeksBefore
	}

	maxWeekIndex := 0
	if weeksAfter > 0 {
		maxWeekIndex = weeksAfter
	}

	startDate := currentMonday.AddDate(0, 0, minWeekIndex*7)
	currentSunday := currentMonday.AddDate(0, 0, 6)
	endDate := currentSunday.AddDate(0, 0, maxWeekIndex*7)

	c := &Calendar{
		now:       now,
		curr:      startDate,
		startDate: startDate,
		endDate:   endDate,
	}

	return c
}

type DateCell struct {
	Day           int
	Month         string
	Year          int
	WeekdayPos    int
	IsCurrentYear bool
	IsToday       bool
	IsWeekend     bool
	IsMonthStart  bool
	IsLastWeekday bool
	InPast        bool
}

func (c *Calendar) Next() *DateCell {
	if c.curr.After(c.endDate) {
		return nil
	}

	var weekdayPos int
	if c.curr.Weekday() == time.Sunday {
		weekdayPos = 6
	} else {
		weekdayPos = int(c.curr.Weekday()) - 1
	}

	cell := &DateCell{
		Day:           c.curr.Day(),
		Month:         c.curr.Month().String(),
		Year:          c.curr.Year(),
		WeekdayPos:    weekdayPos,
		IsCurrentYear: c.curr.Year() == c.now.Year(),
		IsToday:       c.curr.Equal(c.now),
		InPast:        c.curr.Before(c.now),
		IsWeekend:     c.curr.Weekday() == time.Sunday || c.curr.Weekday() == time.Saturday,
		IsMonthStart:  c.curr.Day() == 1,
		IsLastWeekday: c.curr.Weekday() == time.Sunday,
	}

	c.curr = c.curr.Add(24 * time.Hour)
	return cell
}
