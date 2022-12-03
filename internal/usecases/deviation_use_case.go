package usecases

import (
	"context"
	"math"

	"github.com/jhmorais/anomalies-detection/internal/repositories"
	"github.com/jhmorais/anomalies-detection/internal/usecases/contracts"
	"github.com/jhmorais/anomalies-detection/internal/usecases/ports/input"
)

type deviationUseCase struct {
	metricRepository repositories.MetricRepository
}

func NewDeviationUseCase(metricRepository repositories.MetricRepository) contracts.DeviationUseCase {

	return &deviationUseCase{
		metricRepository: metricRepository,
	}
}

func (c *deviationUseCase) Execute(ctx context.Context, average int, metrics []*input.Metric) []int {
	deviationList := make([]int, len(metrics))

	for _, metric := range metrics {
		deviation := int(math.Abs(float64(average - metric.Value)))
		deviationList = append(deviationList, deviation)
	}

	return deviationList
}
