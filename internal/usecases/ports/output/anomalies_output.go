package output

import "time"

type AnomaliesOutput struct {
	SiteId                  string
	OutliersDetectionMethod string
	checkTimeStart          time.Time
	checkTimeEnd            time.Time
	TimeAgo                 string
	TimeStep                string
	DateStart               time.Time
	DateEnd                 time.Time
	Result                  Result
}
