package input

type Metric struct {
	Name      string `json:"name"`
	Attribute string `json:"attribute"`
	Value     int    `json:"value"`
}
