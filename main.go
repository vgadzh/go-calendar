package main

import (
	"fmt"
	"time"
)

func main() {
	//now := time.Now()
	now, err := time.Parse("2006-01-02", "2024-04-15")
	if err != nil {
		fmt.Println(err)
	}
	weekdayNumber := 1
	for i := -7 + 1; i <= 28+7; i++ {
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
