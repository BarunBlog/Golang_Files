// Converting 12hr clock format to 24hr

package main

import (
	"fmt"
	"strconv"
)

func timeConversion(s string) string {

	var ln = len(s)

	if s[:2] == "12" && s[ln-2:] == "AM" {
		return "00" + s[2:ln-2]
	}
	if s[:2] == "12" && s[ln-2:] == "PM" {
		return s[:ln-2]
	}

	if s[ln-2:] == "PM" {
		hr, _ := strconv.Atoi(s[0:2])

		hr += 12
		return strconv.Itoa(hr) + s[2:ln-2]
	} else {
		return s[:ln-2]
	}

}

func main() {
	var time string = "12:01:00PM" // 07:05:45PM

	var timein24 = timeConversion(time)

	fmt.Println(timein24)
}
