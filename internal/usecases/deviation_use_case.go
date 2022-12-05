package usecases

import (
	"context"
	"math"

	"github.com/jhmorais/anomalies-detection/internal/repositories"
	"github.com/jhmorais/anomalies-detection/internal/usecases/contracts"
)

type deviationUseCase struct {
	metricRepository repositories.MetricRepository
}

func NewDeviationUseCase(metricRepository repositories.MetricRepository) contracts.DeviationUseCase {

	return &deviationUseCase{
		metricRepository: metricRepository,
	}
}

func (c *deviationUseCase) Execute(ctx context.Context, average float64) ([]int, error) {
	metrics, err := c.metricRepository.ListMetric(ctx)
	if err != nil {
		return nil, err
	}
	deviationList := make([]int, 0)

	for _, metric := range metrics {
		deviation := int(math.Abs(average - metric.Value))
		deviationList = append(deviationList, deviation)
	}

	return deviationList, nil
}
