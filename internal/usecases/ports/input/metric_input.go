package input

import "time"

type Metric struct {
	Name      string    `json:"name"`
	Attribute string    `json:"attribute"`
	Value     float64   `json:"value"`
	Date      time.Time `json:"date"`
	Child     []*Metric `json:"child"`
}
