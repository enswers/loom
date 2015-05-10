package config

import (
	"time"
)

type Retry struct {
	Number           int
	Interval         string
	intervalDuration *time.Duration
}

func (r *Retry) GetInterval() (time.Duration, error) {
	if r.intervalDuration != nil {
		return *r.intervalDuration, nil
	}
	if r.Interval != "" {
		d, err := time.ParseDuration(r.Interval)
		if err != nil {
			return time.Duration(0), err
		}
		r.intervalDuration = &d
		return d, err
	}
	return *r.intervalDuration, nil
}
