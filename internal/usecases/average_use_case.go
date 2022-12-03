package usecases

import (
	"context"

	"github.com/jhmorais/anomalies-detection/internal/repositories"
	"github.com/jhmorais/anomalies-detection/internal/usecases/contracts"
	"github.com/jhmorais/anomalies-detection/internal/usecases/ports/input"
)

type averageUseCase struct {
	metricRepository repositories.MetricRepository
}

func NewAverageUseCase(metricRepository repositories.MetricRepository) contracts.AverageUseCase {

	return &averageUseCase{
		metricRepository: metricRepository,
	}
}

func (c *averageUseCase) Execute(ctx context.Context, metrics []*input.Metric) (average int, valuesMaps map[int]int) {
	valuesMaps = make(map[int]int)
	total := 0
	for _, metric := range metrics {
		valuesMaps[metric.Value] = metric.Value
		total += metric.Value
	}

	average = total / len(metrics)

	return average, valuesMaps
}
