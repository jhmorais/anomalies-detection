package usecases

import (
	"context"
	"math"

	"github.com/jhmorais/anomalies-detection/internal/repositories"
	"github.com/jhmorais/anomalies-detection/internal/usecases/contracts"
)

type varianceUseCase struct {
	metricRepository repositories.MetricRepository
}

func NewVarianceUseCase(metricRepository repositories.MetricRepository) contracts.VarianceUseCase {

	return &varianceUseCase{
		metricRepository: metricRepository,
	}
}

func (c *varianceUseCase) Execute(ctx context.Context, deviation []int) int {
	var variance int
	for _, val := range deviation {
		variance += int(math.Pow(float64(val), 2))
	}

	variance = int(variance / len(deviation))

	return variance
}
