package output

import "time"

type AlarmOutput struct {
	OutlierPeriodStart time.Time
	OutlierPeriodEnd   time.Time
	Metric             string
	Attribute          string
}
