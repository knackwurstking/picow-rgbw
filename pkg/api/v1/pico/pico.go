package pico

import (
	"fmt"
	"strconv"
)

func PathGetDuty() string {
	return "/rgbw/get_duty"
}

func PathGetPins() string {
	return "/rgbw/get_pins"
}

// PathSetDuty returns the server path to change the current duty cycle
// (use -1 or 0 if pin is not in use)
func PathSetDuty(r, g, b, w int) string {
	query := ""

	for i, n := range [4]int{r, g, b, w} {
		switch i {
		case 0:
			query += "r="
		case 1:
			query += "g="
		case 2:
			query += "b="
		case 3:
			query += "w="
		}

		query += strconv.Itoa(n)

		if i < 3 {
			query += "&"
		}
	}

	return fmt.Sprintf("/rgbw/set_pwm?%s", query)
}

func PathSetPins(r, g, b, w int) string {
	query := ""

	for i, n := range [4]int{r, g, b, w} {
		switch i {
		case 0:
			query += "r="
		case 1:
			query += "g="
		case 2:
			query += "b="
		case 3:
			query += "w="
		}

		query += strconv.Itoa(n)

		if i < 3 {
			query += "&"
		}
	}

	return fmt.Sprintf("/rgbw/set_pins?%s", query)
}
