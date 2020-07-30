package dates

import "time"

// StringToDate convertie uen chaine en date
func StringToDate(date string) time.Time {
	layout := "2006-01-02"
	d, _ := time.Parse(layout, date)
	return d
}
