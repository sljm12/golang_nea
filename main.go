package main

import (
	"fmt"
	"net/http"
)

func get_nearest_5min(minute int) int {
	return minute - (minute % 5)
}

func main() {
	resp, err := http.Get("http://example.com/")

	if err != nil {
		//handle error
		fmt.Println("Error in loading request")
	}

	defer resp.Body.Close()
}
