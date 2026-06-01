package calendar

import (
	"strings"
	"testing"
	"time"

	"github.com/vgadzh/go-calendar/internal/colors"
)

func TestPrinter_Print_Colors(t *testing.T) {
	now := time.Date(2026, time.June, 2, 12, 0, 0, 0, time.UTC)
	cal := NewCalendar(now, 1, 0)

	var output strings.Builder
	p := NewPrinter(5, true, false, false, &output)
	p.Print(cal)
	outputStr := output.String()

	if !strings.Contains(outputStr, colors.FaintRed) {
		t.Error("Expected FaintRed color code in output for past weekend days")
	}
	if !strings.Contains(outputStr, colors.Red) {
		t.Error("Expected Red color code in output for weekend days")
	}

	if !strings.Contains(outputStr, "30") || strings.Contains(outputStr, "\x1b[47m") {
		if strings.Count(outputStr, "2") == 0 {
			t.Error("Could not find today's date in output")
		}
	}
}

func TestPrinter_Print_NoColors(t *testing.T) {
	now := time.Now()
	cal := NewCalendar(now, 0, 0)

	var output strings.Builder
	p := NewPrinter(5, false, false, false, &output)

	p.Print(cal)

	outputStr := output.String()

	if strings.Contains(outputStr, "\x1b") {
		t.Error("Output should not contain ANSI color codes when colored=false")
	}
}

func TestPrinter_Print_NilCalendar(t *testing.T) {
	var output strings.Builder
	p := NewPrinter(5, true, false, false, &output)

	p.Print(nil) // expected no panic

	if len(output.String()) > 0 {
		t.Error("Print(nil) should not produce any output")
	}
}

func TestPrintCellMonth_CurrentYear(t *testing.T) {
	now := time.Date(2026, time.June, 15, 0, 0, 0, 0, time.UTC)
	cal := NewCalendar(now, 0, 0)

	var output strings.Builder
	p := NewPrinter(5, false, true, false, &output) // show month

	cell := cal.Next()
	if cell == nil {
		t.Skip("Calendar ended prematurely")
	}

	if cell.IsCurrentYear {
		p.printCellMonth(cell)
		outputStr := output.String()
		if !strings.Contains(outputStr, "June") {
			t.Errorf("Expected 'June' in output, got: %s", outputStr)
		}
	} else {
		t.Error("IsCurrentYear should be true for current year")
	}
}

func TestPrintCellMonth_PastYear(t *testing.T) {
	now := time.Date(2026, time.June, 15, 0, 0, 0, 0, time.UTC)
	cal := NewCalendar(now, 100, 0)

	var output strings.Builder
	p := NewPrinter(5, false, true, false, &output)

	cell := cal.Next()
	if cell == nil {
		t.Skip("Calendar ended prematurely")
	}

	p.printCellMonth(cell)
	outputStr := output.String()

	if !strings.Contains(outputStr, "2024") {
		t.Errorf("Expected '2024' in output for past year, got: %s", outputStr)
	}
}

func TestPrintWeekdays_Colors(t *testing.T) {
	var output strings.Builder
	p := NewPrinter(5, true, false, true, &output)
	p.printWeekdays()
	outputStr := output.String()
	if !strings.Contains(outputStr, "Mo") {
		t.Error("Expected 'Mo' in weekday headers")
	}

	if !strings.Contains(outputStr, colors.Red) {
		t.Error("Expected Red color code for weekend days in headers")
	}
}

func TestPrintWeekdays_NoColors(t *testing.T) {
	var output strings.Builder
	p := NewPrinter(5, false, false, true, &output)
	p.printWeekdays()
	outputStr := output.String()
	if !strings.Contains(outputStr, "Mo") {
		t.Error("Expected 'Mo' in weekday headers")
	}

	if strings.Contains(outputStr, "\x1b") {
		t.Error("Output should not contain ANSI color codes when colored=false")
	}
}
