package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/urfave/cli"
	colors "github.com/vgadzh/go-calendar/pkg/helper"
)

// main
func main() {
	app := cli.NewApp()
	app.Name = "Calendar"
	app.Usage = "Prints a calendar for the current week. Optionally, it can print the given count of weeks before and after the current week."
	app.Commands = []cli.Command{
		{
			Name: "version",
			Usage: "Print version information",
			Action: func(c *cli.Context) error {
				fmt.Println("v0.3")
				return nil
			},
		},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "before, b",
			Value: "0",
			Usage: "Previous week count",
		},
		cli.StringFlag{
			Name:  "after, a",
			Value: "0",
			Usage: "Next week count",
		},
		cli.BoolFlag{
			Name:  "print-month, m",
			Usage: "Print current month",
		},
		cli.BoolFlag{
			Name:  "colored-output, c",
			Usage: "Use colors to highlight the output",
		},
		cli.BoolFlag{
			Name:  "print-weekdays, w",
			Usage: "Print weekday names",
		},
	}
	app.Action = func(c *cli.Context) error {
		weeksBefore, err := strconv.Atoi(c.GlobalString("before"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "weeksBefore parsing error %v", err)
			return err
		}
		weeksAfter, err := strconv.Atoi(c.GlobalString("after"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "weeksAfter parsing error %v", err)
			return err
		}
		printMonth := c.GlobalBool("print-month")
		useColors := c.GlobalBool("colored-output")
		printWeekdays := c.GlobalBool("print-weekdays")

		printCalendar(weeksBefore, weeksAfter, printMonth, useColors, printWeekdays)
		return nil
	}
	app.Run(os.Args)

}

// printCalendar
func printCalendar(weeksBefore, weeksAfter int, printMonth, useColors, printWeekdays bool) {
	now := time.Now()

	if printMonth {
		if useColors {
			fmt.Println(colors.GetColoredString(now.Month().String(), colors.BoldWhite))
		} else {
			fmt.Println(now.Month())
		}
	}

	tabSize := 5
	weekDayLetterCount := 3
	spaceCount := tabSize - weekDayLetterCount
	if printWeekdays {
		// Monday to Saturday
		for i := 1; i <= 6; i++ {
			name := time.Weekday(i).String()[0:weekDayLetterCount]
			if useColors && i == 6 {
				fmt.Print(getWeekendStyledString(name))
			} else {
				fmt.Print(name)
			}
			fmt.Print(getSpaces(spaceCount))
		}
		// Sunday
		name := time.Weekday(0).String()[0:weekDayLetterCount]
		if useColors {
			fmt.Print(getWeekendStyledString(name))
		} else {
			fmt.Print(name)
		}
		fmt.Print(getSpaces(spaceCount))
		fmt.Println()
	}

	weekdayNumber := int(now.Weekday())
	for i := 1 - 7*weeksBefore; i <= 7+7*weeksAfter; i++ {
		newDate := now.AddDate(0, 0, i-weekdayNumber)
		str := strconv.Itoa(newDate.Day())

		if useColors {
			if newDate.Equal(now) && isWeekend(newDate) {
				// Today is weekend
				str = colors.GetColoredString(strconv.Itoa(newDate.Day()), colors.Black, colors.OnRed)
			} else if newDate.Equal(now) && !isWeekend(newDate) {
				// Today
				str = colors.GetColoredString(strconv.Itoa(newDate.Day()), colors.Black, colors.OnWhite)
			} else if newDate.Before(now) && !isWeekend(newDate) {
				// Past day
				str = colors.GetColoredString(strconv.Itoa(newDate.Day()), colors.FaintWhite)
			} else if newDate.Before(now) && isWeekend(newDate) {
				// Past day and weekend
				str = colors.GetColoredString(strconv.Itoa(newDate.Day()), colors.FaintRed)
			} else if isWeekend(newDate) {
				// Future weekend
				str = colors.GetColoredString(strconv.Itoa(newDate.Day()), colors.Red)
			}
		}

		spaceCount := tabSize - len(strconv.Itoa(newDate.Day()))
		fmt.Print(str)
		fmt.Print(getSpaces(spaceCount))
		if i%7 == 0 {
			fmt.Println()
		}
	}
}

// getSpaces
func getSpaces(count int) string {
	return strings.Repeat(" ", count)
}

// getWeekendStyledString
func getWeekendStyledString(text string) string {
	return colors.GetColoredString(text, colors.Red)
}

// isWeekend
func isWeekend(day time.Time) bool {
	return day.Weekday() == 0 || day.Weekday() == 6
}
