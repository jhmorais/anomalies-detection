package usecases

import (
	"context"

	"github.com/jhmorais/anomalies-detection/internal/repositories"
	"github.com/jhmorais/anomalies-detection/internal/usecases/contracts"
	"github.com/jhmorais/anomalies-detection/internal/usecases/ports/input"
	"github.com/jhmorais/anomalies-detection/internal/usecases/ports/output"
)

type anomaliesUseCase struct {
	metricRepository repositories.MetricRepository
}

func NewAnomaliesUseCase(metricRepository repositories.MetricRepository) contracts.AnomaliesUseCase {

	return &anomaliesUseCase{
		metricRepository: metricRepository,
	}
}

func (c *anomaliesUseCase) Execute(ctx context.Context, metric *input.DatasetInput) (*output.AnomaliesOutput, error) {
	//Method to create anomalies
	return nil, nil
}
