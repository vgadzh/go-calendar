# Go calendar

## About

Prints a calendar for the current week. Optionally, it can print the given count of weeks before and after the current week.

## Using examples

Print the current week

```bash
./go-calendar 
22   23   24   25   26   27   28
```

Print the month name

```bash
./go-calendar -m
April
22   23   24   25   26   27   28
```

Print one week before the current week

```bash
./go-calendar -b 1
15   16   17   18   19   20   21
22   23   24   25   26   27   28
```

Print one week after the current week

```bash
./go-calendar -a 1
22   23   24   25   26   27   28
29   30   1    2    3    4    5
```

Print one week before and 3 weeks after the current week

```bash
./go-calendar -b 1 -a 3
15   16   17   18   19   20   21
22   23   24   25   26   27   28
29   30   1    2    3    4    5
6    7    8    9    10   11   12
13   14   15   16   17   18   19
```

Print help

```bash
./go-calendar help
NAME:
   Calendar - Prints a calendar for the current week. Optionally, it can print the given count of weeks before and after the current week.

USAGE:
   main [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --before value, -b value  Previous week count (default: "0")
   --after value, -a value   Next week count (default: "0")
   --print-month, -m         Print current month
   --help, -h  
```