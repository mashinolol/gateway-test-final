package transactions

import (
	"fmt"
	"sort"
	"time"
)

type Transaction struct {
	Value     int
	Timestamp time.Time
}

func intervalSlices(t time.Time, interval string) (time.Time, error) {
	switch interval {
	case "month":
		return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location()), nil
	case "week":
		year, week := t.ISOWeek()
		return time.Date(year, 0, (week-1)*7+1, 0, 0, 0, 0, t.Location()), nil
	case "day":
		return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()), nil
	case "hour":
		return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location()), nil
	default:
		return time.Time{}, fmt.Errorf("invalid interval specified")
	}
}

func Group(transactions []*Transaction, interval string) ([]*Transaction, error) {
	grouped := make(map[time.Time]*Transaction)

	for _, tr := range transactions {
		roundedTime, err := intervalSlices(tr.Timestamp, interval)
		if err != nil {
			return nil, err
		}
		if existing, found := grouped[roundedTime]; !found || existing.Timestamp.Before(tr.Timestamp) {
			grouped[roundedTime] = tr
		}
	}

	var result []*Transaction
	for _, tr := range grouped {
		result = append(result, tr)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Timestamp.Before(result[j].Timestamp)
	})

	return result, nil
}
