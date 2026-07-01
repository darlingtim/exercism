package booking

import (
    "time"
    "fmt"
    "reflect"
    "strconv"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    layout:= "1/2/2006 15:04:05"
    dateTime, err:= time.Parse(layout, date)
    if err != nil {
        panic("unable to parse time")
    }
    
    return dateTime
	panic("Please implement the Schedule function")
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    layout := "January 2, 2006 15:04:05"
    dateTime, err:= time.Parse(layout, date)
    if err != nil {
        panic("unable to parse time")
    }
   return dateTime.Before(time.Now())
	panic("Please implement the HasPassed function")
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    layout := "Monday, January 2, 2006 15:04:05"
    dateTime, err := time.Parse(layout, date)
    if err != nil {
        panic("unable to parse time")
    }
    if (dateTime.Hour() >= 12) && (dateTime.Hour() < 18) {
        return true
    }else{
        return false
    }
    /*if (dateTime.Hour() * 3600) + (dateTime.Minute() * 60) + (dateTime.Second()) < 43200 || (dateTime.Hour() * 3600) + (dateTime.Minute() * 60) + (dateTime.Second()) >= 64800  {
        return false
    }else{
        return true
    } */
    
	panic("Please implement the IsAfternoonAppointment function")
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
    dateTime := Schedule(date)
    if reflect.TypeOf(dateTime) == reflect.TypeOf(time.Time{}) {
        //year, month, day := dateTime.Date()
        //hour, minute, second := dateTime()
        return fmt.Sprintf("You have an appointment on %v, at %v.", dateTime.Format("Monday, January 2, 2006"), dateTime.Format("15:04"))
    }else{
        return fmt.Sprint(fmt.Errorf("Wrong date type was recieved"))
    }
	panic("Please implement the Description function")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    layout:= "2006-1-2"
    date := strconv.Itoa(time.Now().Year()) + "-9-15"
    anniversaryDate, err := time.Parse(layout, date)
    if err != nil {
        fmt.Errorf("issue parsing time")
    }
    return anniversaryDate
	panic("Please implement the AnniversaryDate function")
}
