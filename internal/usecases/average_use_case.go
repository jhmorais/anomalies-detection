package usecases

import (
	"context"

	"github.com/jhmorais/anomalies-detection/internal/repositories"
	"github.com/jhmorais/anomalies-detection/internal/usecases/contracts"
)

type averageUseCase struct {
	metricRepository repositories.MetricRepository
}

func NewAverageUseCase(metricRepository repositories.MetricRepository) contracts.AverageUseCase {

	return &averageUseCase{
		metricRepository: metricRepository,
	}
}

func (c *averageUseCase) Execute(ctx context.Context) (average float64, err error) {
	total := float64(0)
	metrics, err := c.metricRepository.ListMetric(ctx)
	if err != nil {
		return average, err
	}

	for _, metric := range metrics {
		total += metric.Value
	}

	average = total / float64(len(metrics))

	return average, err
}
