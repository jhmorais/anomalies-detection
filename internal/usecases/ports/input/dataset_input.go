package input

type DatasetInput struct {
	SiteID                  string                  `json:"siteId"`
	TimeAgo                 string                  `json:"timeAgo"`
	TimeStep                string                  `json:"timeStep"`
	OutliersDetectionMethod string                  `json:"outliersDetectionMethod"`
	OutliersDetection       *OutliersDetectionInput `json:"outliersDetection"`
	MetricsList             []*Metric               `json:"metricsList"`
}
