package booking

import (
    "time"
    "fmt"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    t, err := time.Parse("1/02/2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    layout := "January 2, 2006 15:04:05"
    dt, err := time.Parse(layout, date)
    if(err!=nil) {
        panic(err)
    }
	return dt.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    layout := "Monday, January 2, 2006 15:04:05"
    formattedTime, _ := time.Parse(layout, date)
	hour := formattedTime.Hour()
    fmt.Printf("Hour : %v\n", hour)
    return hour>=12 && hour<18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
    layout := "1/2/2006 15:04:05"
    formattedTime, _ := time.Parse(layout, date)
    weekday := formattedTime.Weekday()
    month := formattedTime.Month()
    day := formattedTime.Day()
    year := formattedTime.Year()
	hour := formattedTime.Hour()
    minute := formattedTime.Minute()
	return fmt.Sprintf("You have an appointment on %v, %v %v, %v, at %v:%v.", weekday, month, day, year, hour, minute  )
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	anniversary := time.Date(2024, time.September, 15, 0, 0, 0, 0, time.UTC)
    return anniversary
}
