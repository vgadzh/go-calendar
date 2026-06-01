package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/vgadzh/go-calendar/internal/calendar"
)

const version = "0.3.4"

func main() {
	before := flag.Int("b", 0, "Weeks before current")
	after := flag.Int("a", 0, "Weeks after current")
	printMonth := flag.Bool("m", false, "Print month names")
	printWeekdays := flag.Bool("w", false, "Print weekday names")
	colored := flag.Bool("c", false, "Use colors")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Version: %s\n", version)
		fmt.Fprintf(os.Stderr, "Usage: go-calendar [-b weeks] [-a weeks] [-m] [-w] [-c]\n")
		fmt.Fprintf(os.Stderr, "-b: weeks before current\n")
		fmt.Fprintf(os.Stderr, "-a: weeks after current\n")
		fmt.Fprintf(os.Stderr, "-m: print month names\n")
		fmt.Fprintf(os.Stderr, "-w: print weekday names\n")
		fmt.Fprintf(os.Stderr, "-c: use colors\n")
		fmt.Fprintln(os.Stderr, "\nExample: go-calendar -b 1 -a 3 -m -w -c")
	}

	flag.Parse()

	if flag.NArg() > 0 {
		fmt.Println("No arguments supported directly. Use flags.")
		os.Exit(1)
	}

	if *after < 0 {
		fmt.Println("Invalid -a argument.")
		os.Exit(1)
	}
	if *before < 0 {
		fmt.Println("Invalid -b argument.")
		os.Exit(1)
	}

	now := time.Now()
	cal := calendar.NewCalendar(now, *before, *after)
	printer := calendar.NewPrinter(5, *colored, *printMonth, *printWeekdays)

	printer.Print(cal)
}
