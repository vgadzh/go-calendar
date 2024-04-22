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

func main() {
	app := cli.NewApp()
	app.Name = "Calendar"
	app.Usage = "Prints a calendar for the current week. Optionally, it can print the given count of weeks before and after the current week."
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

		printCalendar(weeksBefore, weeksAfter, printMonth, useColors)
		return nil
	}
	app.Run(os.Args)

}
func printCalendar(weeksBefore, weeksAfter int, printMonth, useColors bool) {
	now := time.Now()

	if printMonth {
		if useColors {
			fmt.Println(colors.GetColoredString(now.Month().String(), colors.BoldWhite))
		} else {
			fmt.Println(now.Month())
		}
	}

	weekdayNumber := int(now.Weekday())
	for i := 1 - 7*weeksBefore; i <= 7+7*weeksAfter; i++ {
		newDate := now.AddDate(0, 0, i-weekdayNumber)
		var str string
		if useColors && newDate.Equal(now) {
			str = colors.GetColoredString(strconv.Itoa(newDate.Day()), colors.Black, colors.OnWhite)
		} else if newDate.Before(now){
			str = colors.GetColoredString(strconv.Itoa(newDate.Day()), colors.FaintWhite)
		} else {
			str = strconv.Itoa(newDate.Day())
		}
		tabSize := 5
		tabs := tabSize - len(strconv.Itoa(newDate.Day()))
		fmt.Print(str)
		fmt.Print(strings.Repeat(" ", tabs))
		if i%7 == 0 {
			fmt.Println()
		}
	}
}
