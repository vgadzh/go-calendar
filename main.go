package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/urfave/cli"
	calendar "github.com/vgadzh/go-calendar/pkg"
)

// main
func main() {
	app := cli.NewApp()
	app.Name = "Calendar"
	app.Usage = "Prints a calendar for the current week. Optionally, it can print the given count of weeks before and after the current week."
	app.Commands = []cli.Command{
		{
			Name:  "version",
			Usage: "Print version information",
			Action: func(c *cli.Context) error {
				fmt.Println("v0.3.3")
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
		os := calendar.OutputSettings{
			Month:    c.GlobalBool("print-month"),
			Weekdays: c.GlobalBool("print-weekdays"),
			Colors:   c.GlobalBool("colored-output"),
		}
		cal := calendar.New(weeksBefore, weeksAfter, os)
		fmt.Println(cal.String())
		return nil
	}
	app.Run(os.Args)

}
