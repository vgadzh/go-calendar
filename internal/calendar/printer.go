package calendar

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/vgadzh/go-calendar/internal/colors"
)

type Printer struct {
	tabSize      int
	useColors    bool
	showMonth    bool
	showWeekdays bool
	writer       io.Writer
}

func NewPrinter(tabSize int, colored, printMonth, printWeekdays bool, w io.Writer) *Printer {
	return &Printer{
		tabSize:      tabSize,
		useColors:    colored,
		showMonth:    printMonth,
		showWeekdays: printWeekdays,
		writer:       w,
	}
}

func (p *Printer) Print(cal *Calendar) {
	if cal == nil {
		return
	}

	firstLine := true
	for cell := cal.Next(); cell != nil; cell = cal.Next() {
		switch {
		case firstLine:
			if p.showMonth {
				p.printCellMonth(cell)
			}
			if p.showWeekdays {
				p.printWeekdays()
			}

		case cell.IsMonthStart:
			_, _ = fmt.Fprintln(p.writer)
			_, _ = fmt.Fprintln(p.writer)
			if p.showMonth {
				p.printCellMonth(cell)
			}
			if p.showWeekdays {
				p.printWeekdays()
			}
			for range cell.WeekdayPos {
				p.printTabbed("") // placeholder for a first day of the month
			}
		}

		firstLine = false

		switch {
		case cell.InPast && p.useColors && cell.IsWeekend:
			dayStr := strconv.Itoa(cell.Day)
			text := colors.GetColoredString(dayStr, colors.FaintRed)
			p.printTabbedColored(text, len(dayStr))

		case cell.InPast && p.useColors:
			dayStr := strconv.Itoa(cell.Day)
			text := colors.GetColoredString(dayStr, colors.FaintWhite)
			p.printTabbedColored(text, len(dayStr))

		case cell.IsToday && cell.IsWeekend && p.useColors:
			dayStr := strconv.Itoa(cell.Day)
			text := colors.GetColoredString(dayStr, colors.Black, colors.OnRed)
			p.printTabbedColored(text, len(dayStr))

		case cell.IsWeekend && p.useColors:
			dayStr := strconv.Itoa(cell.Day)
			text := colors.GetColoredString(dayStr, colors.Red)
			p.printTabbedColored(text, len(dayStr))

		case cell.IsToday && p.useColors:
			dayStr := strconv.Itoa(cell.Day)
			text := colors.GetColoredString(dayStr, colors.Black, colors.OnWhite)
			p.printTabbedColored(text, len(dayStr))

		default:
			p.printTabbed(strconv.Itoa(cell.Day))
		}

		if cell.IsLastWeekday {
			_, _ = fmt.Fprintln(p.writer)
		}
	}
}

func (p *Printer) printCellMonth(cell *DateCell) {
	if cell == nil {
		return
	}
	if cell.IsCurrentYear {
		_, _ = fmt.Fprintln(p.writer, cell.Month)
	} else {
		_, _ = fmt.Fprintln(p.writer, cell.Month, cell.Year)
	}
}

func (p *Printer) printWeekdays() {
	weekdays := []string{"Mo", "Tu", "We", "Th", "Fr", "Sa", "Su"}
	for i, w := range weekdays {
		if i < 5 {
			p.printTabbed(w)
		} else if p.useColors {
			text := colors.GetColoredString(w, colors.Red)
			p.printTabbedColored(text, len(w))
		} else {
			p.printTabbed(w)
		}
	}
	_, _ = fmt.Fprintln(p.writer)
}

func (p *Printer) printTabbed(text string) {
	_, _ = fmt.Fprintf(p.writer, "%-*s", p.tabSize, text)
}

func (p *Printer) printTabbedColored(ansiStr string, visibleLen int) {
	spaces := p.tabSize - visibleLen
	text := ansiStr + strings.Repeat(" ", spaces)
	_, _ = fmt.Fprintf(p.writer, "%-*s", p.tabSize, text)
}
