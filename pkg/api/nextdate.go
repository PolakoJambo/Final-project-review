package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const dateFormat = "20060102"

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	date, err := time.Parse(dateFormat, dstart)
	if err != nil {
		return "", err
	}

	if repeat == "" {
		return "", nil
	}

	parts := strings.Split(repeat, " ")
	rule := parts[0]

	switch rule {
	case "d":
		if len(parts) != 2 {
			return "", fmt.Errorf("invalid 'd' rule format")
		}
		days, err := strconv.Atoi(parts[1])
		if err != nil || days < 1 || days > 400 {
			return "", fmt.Errorf("invalid day interval")
		}

		for {
			date = date.AddDate(0, 0, days)
			if afterNow(date, now) {
				break
			}
		}
	case "y":
		if len(parts) != 1 {
			return "", fmt.Errorf("invalid 'y' rule format")
		}

		for {
			date = date.AddDate(1, 0, 0)
			if afterNow(date, now) {
				break
			}
		}
	default:
		return "", fmt.Errorf("unsupported format")
	}

	return date.Format(dateFormat), nil
}

func afterNow(date, now time.Time) bool {
	dateStr := date.Format(dateFormat)
	nowStr := now.Format(dateFormat)
	return dateStr > nowStr
}

func nextDateHandler(w http.ResponseWriter, r *http.Request) {
	nowParam := r.FormValue("now")
	dateParam := r.FormValue("date")
	repeatParam := r.FormValue("repeat")

	var now time.Time
	var err error
	if nowParam == "" {
		now = time.Now()
	} else {
		now, err = time.Parse(dateFormat, nowParam)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	nextDate, err := NextDate(now, dateParam, repeatParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, nextDate)
}
