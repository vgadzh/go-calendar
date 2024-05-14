# Go calendar

## About

Prints a calendar for the current week. Optionally, it can print the given count of weeks before and after the current week. Supports ANSI colored output.

## Using examples

Print the current week

```bash
./go-calendar 
22   23   24   25   26   27   28
```

Print one week before the current week, 3 weeks after the current week. Print month name and weekday names.

```bash
./go-calendar -m -w -b 1 -a 3
April
Mo   Tu   We   Th   Fr   Sa   Su
15   16   17   18   19   20   21
22   23   24   25   26   27   28
29   30

May
Mo   Tu   We   Th   Fr   Sa   Su
          1    2    3    4    5
6    7    8    9    10   11   12
13   14   15   16   17   18   19
```

## Help

```bash
./go-calendar help
NAME:
   Calendar - Prints a calendar for the current week. Optionally, it can print the given count of weeks before and after the current week.

USAGE:
   main [global options] command [command options] [arguments...]

COMMANDS:
   version  Print version information
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --before value, -b value  Previous week count (default: "0")
   --after value, -a value   Next week count (default: "0")
   --print-month, -m         Print current month
   --colored-output, -c      Use colors to highlight the output
   --print-weekdays, -w      Print weekday names
   --help, -h  
```

## Screenshots

![image](/screenshots/1.png?raw=true)
