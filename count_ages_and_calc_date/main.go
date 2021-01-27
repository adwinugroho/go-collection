package main

import (
	"log"
	"time"
)

func TimeHostNow() time.Time {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Printf("Error get time, cause:%+v\n", err)
	}
	now := time.Now()
	timeInLoc := now.In(location)
	return timeInLoc
}

func CalcDate(years int) string {
	if years == 0 {
		return ""
	}
	now := TimeHostNow()
	years = (years * -1)
	result := now.AddDate(years, 0, 0)
	return result.Format("20060102")
}

func CountAges(birthdayDate string) int {
	now := TimeHostNow()
	birthdayToTime, _ := time.Parse("20060102", birthdayDate)
	years := now.Year() - birthdayToTime.Year()
	if now.YearDay() < birthdayToTime.YearDay() {
		years--
	}
	return years
}
