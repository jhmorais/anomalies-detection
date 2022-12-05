package entities

import "time"

type Metric struct {
	Name      string
	Attribute string
	Value     float64
	Date      time.Time
}
