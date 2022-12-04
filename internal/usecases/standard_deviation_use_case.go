package usecases

import (
	"context"
	"math"

	"github.com/jhmorais/anomalies-detection/internal/repositories"
	"github.com/jhmorais/anomalies-detection/internal/usecases/contracts"
)

type standardDeviationUseCase struct {
	metricRepository repositories.MetricRepository
}

func NewStandardDeviationUseCase(metricRepository repositories.MetricRepository) contracts.StandardDeviationUseCase {

	return &standardDeviationUseCase{
		metricRepository: metricRepository,
	}
}

func (c *standardDeviationUseCase) Execute(ctx context.Context, variance int) float64 {
	standardDeviation := math.Sqrt(float64(variance))

	return standardDeviation
}
