package domain

import "time"

type GyroscopeDomain struct {
	ID             int
	X              float64
	Y              float64
	Z              float64
	DeviceID       int
	MacAddress     string
	CollectionDate time.Time
}
