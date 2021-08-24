package parse_string_datetime_to_time

import (
	"log"
	"time"
)

func ParseStringDatetimeToTime(datetime string) time.Time {
	layout := "20060102150405" // you can explore layout here https://yourbasic.org/golang/format-parse-string-time-date-example/
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Printf("Error get time, cause:%+v\n", err)
	}
	getTime, _ := time.ParseInLocation(layout, datetime, location)
	return getTime
}
