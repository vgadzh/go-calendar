package main

import (
	"fmt"
	"os"
	"strconv"
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

		printCalendar(weeksBefore, weeksAfter)
		return nil
	}
	app.Run(os.Args)

}
func printCalendar(weeksBefore, weeksAfter int) {
	now := time.Now()
	weekdayNumber := int(now.Weekday())

	for i := 1 - 7*weeksBefore; i <= 7+7*weeksAfter; i++ {
		newDay := now.AddDate(0, 0, i-weekdayNumber).Day()
		//if newDay == now.Day() {
		//	fmt.Printf("%s", "*")
		//}
		fmt.Printf("%3d ", newDay)
		if i%7 == 0 {
			fmt.Println()
		}
	}

}
