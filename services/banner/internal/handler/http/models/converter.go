package models

import (
	"time"
)

const (
	UnknownVisibility = "unknown"
)

func ToRequestTime(t time.Time) *time.Time {
	var resT *time.Time
	if t == (time.Time{}) {
		resT = nil
	} else {
		resT = &t
	}
	return resT
}
