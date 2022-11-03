package main

import (
	"fmt"
	"net/http"
	"time"
)

func get_date_string(time time.Time, delta int) (string, string) {
	dt := time
	d := fmt.Sprintf("%02d%02d%02d", dt.Year(), dt.Month(), dt.Day())
	min := dt.Minute()
	min = get_nearest_5min(min)
	min = min - delta
	t := fmt.Sprintf("%02d%02d", dt.Hour(), min)
	return d, t
}

func get_nearest_5min(minute int) int {
	return minute - (minute % 5)
}

func get_url_string(date string, time string) string {
	return fmt.Sprintf("https://www.nea.gov.sg/docs/default-source/rain-area/dpsri_70km_%s%s0000dBR.dpsri.png", date, time)
}

func main() {
	resp, err := http.Get("http://example.com/")

	if err != nil {
		//handle error
		fmt.Println("Error in loading request")
	}

	defer resp.Body.Close()
}
