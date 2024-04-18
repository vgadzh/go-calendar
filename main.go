package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Calendar"
	app.Usage = "Prints current weekdays, optional previous weeks and optional next weeks"
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
			Name:  "print-colors, c",
			Usage: "Print colored output",
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
		printColors := c.GlobalBool("print-colors")

		printCalendar(weeksBefore, weeksAfter, printMonth, printColors)
		//black := color.New(color.FgBlack)
		//whiteBackground := black.Add(color.BgWhite).Add(color.Bold)
		//whiteBackground.Println("Black text with white background.")
		//color.Yellow("cyan")
		//fmt.Println("Regular text")
		return nil
	}
	app.Run(os.Args)

}
func printCalendar(weeksBefore, weeksAfter int, printMonth, printColors bool) {
	now := time.Now()
	if printMonth {
		if printColors {
			color := color.New(color.Bold)
			color.Println(now.Month())
		} else {
			fmt.Println(now.Month())
		}
	}

	weekdayNumber := int(now.Weekday())
	for i := 1 - 7*weeksBefore; i <= 7+7*weeksAfter; i++ {
		newDate := now.AddDate(0, 0, i-weekdayNumber)
		var str string
		if printColors && newDate.Equal(now) {
			black := color.New(color.FgBlack)
			whiteBackground := black.Add(color.BgWhite).Add(color.Bold)
			str = whiteBackground.Sprint(newDate.Day())
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
