package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/urfave/cli"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var BoldGray = "\033[37;1m"
var White = "\033[37m"
var WhiteBold = "\033[37;0m"
var WhiteBg = "\033[47m"

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
	//bold := color.New(color.Bold)
	if printMonth {
		if printColors {
			//bold.Println(now.Month())
			fmt.Println(WhiteBold + now.Month().String() + Reset)
		} else {
			fmt.Println(now.Month())
		}
	}

	weekdayNumber := int(now.Weekday())
	for i := 1 - 7*weeksBefore; i <= 7+7*weeksAfter; i++ {
		newDate := now.AddDate(0, 0, i-weekdayNumber)
		var str string
		if printColors && newDate.Equal(now) {
			//black := color.New(color.FgBlack)
			//whiteBackground := black.Add(color.BgWhite).Add(color.Bold)
			//str = whiteBackground.Sprint(newDate.Day())
			str = WhiteBg + strconv.Itoa(newDate.Day()) + Reset
		} else if printColors && newDate.Before(now) {
			//str = strconv.Itoa(newDate.Day())
			//str = hiWhite.Sprint(newDate.Day())
			str = Gray + strconv.Itoa(newDate.Day()) + Reset
		} else {
			//str = bold.Sprint(newDate.Day())
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
