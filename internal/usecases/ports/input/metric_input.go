package input

type Metric struct {
	Name      string    `json:"name"`
	Attribute string    `json:"attribute"`
	Value     float64   `json:"value"`
	Child     []*Metric `json:"child"`
}
