package service

import "time"

type TimeProvider interface {
	Now() time.Time
}

type RealTime struct {
}

func (r RealTime) Now() time.Time {
	return time.Now()
}
func NewRealTime() TimeProvider {
	return &RealTime{}
}
