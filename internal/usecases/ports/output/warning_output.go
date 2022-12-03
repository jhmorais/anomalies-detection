package output

import "time"

type WarningOutput struct {
	OutlierPeriodStart time.Time
	OutlierPeriodEnd   time.Time
	Metric             string
	Attribute          string
}
