package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// Julian date seperti di website ini https://www.longpelaexpertise.com.au/toolsJulian.php
	now := "2023-01-29"
	res := GetJulianDays(now)
	fmt.Println(res)
}

func GetJulianDays(nowTime string) string {
	// getDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	// layout menyesuaikan parameter nowTime
	layout := "2006-01-02"
	getDate, _ := time.Parse(layout, nowTime)
	//fmt.Println(getDate)
	yearString := nowTime[0:4]
	// monthString := nowTime[5:7]
	// fmt.Println(monthString)
	var isLeapYear bool
	var numberOfDays = 366
	getIntYear, _ := strconv.Atoi(yearString)
	// getIntMonth, _ := strconv.Atoi(monthString)
	// fmt.Println(getIntMonth)
	if (getIntYear%4 == 0 && getIntYear%100 != 0) || getIntYear%400 == 0 {
		isLeapYear = true
	}
	fmt.Println(nowTime[8:10])

	startedTime, _ := time.Parse(layout, fmt.Sprintf("%s-01-01", yearString))
	plusOneYear := startedTime.AddDate(1, 0, 0)
	days := plusOneYear.Sub(getDate).Hours() / 24
	// fmt.Println("cek days before round:", days)
	//var fixDays int
	if isLeapYear {
		// if (getIntMonth == 2 && nowTime[8:10] == "29") || (getIntMonth > 1 && getIntMonth != 2) {
		// 	numberOfDays = 367
		// }
		numberOfDays = 367
	}
	// fmt.Println("cek fixDays:", fixDays)
	// fmt.Println("cek newDays:", newDays)
	// fmt.Println("cek hari:", numberOfDays)
	// fmt.Println("cek days after round:", days)
	days = float64(numberOfDays) - days
	fmt.Println("cek hasil:", days)
	daysToString := strconv.Itoa(int(days))
	if len(daysToString) > 2 {
		return fmt.Sprintf("%s%s", yearString[2:], daysToString)
	}
	return fmt.Sprintf("%s0%s", yearString[2:], daysToString)
}
