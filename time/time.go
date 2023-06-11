package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Println(now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	fmt.Println(year, month, day)
	timestamp := now.Unix()
	nanoTimestamp := now.UnixNano()
	fmt.Println(timestamp, nanoTimestamp)

	later := now.Add(time.Hour)
	fmt.Println(later)

	before := now.Sub(later)
	fmt.Println(before)

	ticker := time.Tick(time.Second * 5)
	for i := range ticker {
		fmt.Println(i)
	}
}
