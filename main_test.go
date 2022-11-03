package main

import (
	"fmt"
	"testing"
	"time"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestNearest5(t *testing.T) {
	for i := 1; i < 5; i++ {
		ans := get_nearest_5min(i)
		if ans != 0 {
			t.Fatalf("failed")
		}
	}

	for i := 21; i < 25; i++ {
		fmt.Println("Testing ", i)
		ans := get_nearest_5min(i)
		if ans != 20 {
			t.Fatalf("failed")
		}
	}
}

func TestURLString(t *testing.T) {
	s := get_url_string("20221103", "0840")
	fmt.Println(s)
	if s != "https://www.nea.gov.sg/docs/default-source/rain-area/dpsri_70km_2022110308400000dBR.dpsri.png" {
		t.Fatalf("failed")
	}

}

func TestDateString(t *testing.T) {
	delta := 5
	dt := time.Date(2022, time.August, 1, 10, 11, 00, 00, time.UTC)
	ds, ts := get_date_string(dt, delta)
	fmt.Println(fmt.Sprintf("Date Time %s %s", ds, ts))

	if ds != "20220801" || ts != "1005" {
		t.Fatalf("Failed")
	}
}
