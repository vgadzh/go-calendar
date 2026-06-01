package calendar

import (
	"testing"
	"time"
)

func TestNewCalendar(t *testing.T) {
	tests := []struct {
		now               time.Time
		weeksBefore       int
		weeksAfter        int
		expectedStartDate time.Time
		expectedEndDate   time.Time
	}{
		{
			now:               time.Date(2026, time.June, 2, 12, 0, 0, 0, time.UTC),
			weeksBefore:       0,
			weeksAfter:        0,
			expectedStartDate: time.Date(2026, time.June, 1, 12, 0, 0, 0, time.UTC),
			expectedEndDate:   time.Date(2026, time.June, 7, 12, 0, 0, 0, time.UTC),
		},
		{
			now:               time.Date(2026, time.June, 2, 12, 0, 0, 0, time.UTC),
			weeksBefore:       1,
			weeksAfter:        0,
			expectedStartDate: time.Date(2026, time.May, 25, 12, 0, 0, 0, time.UTC),
			expectedEndDate:   time.Date(2026, time.June, 7, 12, 0, 0, 0, time.UTC),
		},
		{
			now:               time.Date(2026, time.June, 2, 12, 0, 0, 0, time.UTC),
			weeksBefore:       2,
			weeksAfter:        0,
			expectedStartDate: time.Date(2026, time.May, 18, 12, 0, 0, 0, time.UTC),
			expectedEndDate:   time.Date(2026, time.June, 7, 12, 0, 0, 0, time.UTC),
		},
		{
			now:               time.Date(2026, time.June, 2, 12, 0, 0, 0, time.UTC),
			weeksBefore:       2,
			weeksAfter:        1,
			expectedStartDate: time.Date(2026, time.May, 18, 12, 0, 0, 0, time.UTC),
			expectedEndDate:   time.Date(2026, time.June, 14, 12, 0, 0, 0, time.UTC),
		},
		{
			now:               time.Date(2026, time.June, 2, 12, 0, 0, 0, time.UTC),
			weeksBefore:       2,
			weeksAfter:        2,
			expectedStartDate: time.Date(2026, time.May, 18, 12, 0, 0, 0, time.UTC),
			expectedEndDate:   time.Date(2026, time.June, 21, 12, 0, 0, 0, time.UTC),
		},
	}
	for i, tt := range tests {
		c := NewCalendar(tt.now, tt.weeksBefore, tt.weeksAfter)
		if !c.now.Equal(tt.now) {
			t.Errorf("%d: now date mismatch", i)
		}
		if !c.startDate.Equal(tt.expectedStartDate) {
			t.Errorf("%d: startDate date mismatch", i)
		}
		if !c.endDate.Equal(tt.expectedEndDate) {
			t.Errorf("%d: endDate date mismatch", i)
		}
	}
}

func TestNext(t *testing.T) {
	now := time.Date(2026, time.June, 2, 12, 0, 0, 0, time.UTC)
	cal := NewCalendar(now, 0, 0)
	count := 0
	for cell := cal.Next(); cell != nil; cell = cal.Next() {
		count++
		if cell.Day == 1 && !cell.IsMonthStart {
			t.Error("day 1 is not month start")
		}
		if cell.Day != 1 && cell.IsMonthStart {
			t.Errorf("day %d is not month start", cell.Day)
		}
		if cell.Day == 2 && !cell.IsToday {
			t.Error("day 2 is today")
		}
		if cell.Day != 2 && cell.IsToday {
			t.Errorf("day %d is today", cell.Day)
		}
		if cell.Month != time.June.String() {
			t.Errorf("month %s is not %s", cell.Month, time.June.String())
		}
		if cell.Year != now.Year() {
			t.Errorf("year %d is not %d", cell.Year, now.Year())
		}
		if !cell.IsCurrentYear {
			t.Errorf("isCurrentYear is false, day %d", cell.Day)
		}
		if count < 6 && cell.IsWeekend {
			t.Errorf("day %d should not to be weekend", cell.Day)
		}
		if count >= 6 && !cell.IsWeekend {
			t.Errorf("day %d should be weekend", cell.Day)
		}
		if count == 7 && !cell.IsLastWeekday {
			t.Errorf("day %d should be last weekday", cell.Day)
		}
		if cell.Day == 1 && !cell.InPast {
			t.Errorf("day 1 is past day %d", cell.Day)
		}
		if cell.WeekdayPos != count-1 {
			t.Errorf("weekdayPos %d is not %d", cell.WeekdayPos, count)
		}
	}
	if count != 7 {
		t.Errorf("count: got %d, want %d", count, 7)
	}
}
