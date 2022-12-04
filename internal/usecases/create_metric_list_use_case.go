package usecases

import (
	"context"
	"time"

	"github.com/jhmorais/anomalies-detection/internal/repositories"
	"github.com/jhmorais/anomalies-detection/internal/usecases/contracts"
	"github.com/jhmorais/anomalies-detection/internal/usecases/ports/input"
	"github.com/jhmorais/anomalies-detection/internal/usecases/ports/output"
)

type createMetricListUseCase struct {
	metricRepository repositories.MetricRepository
}

func NewCreateMetricListUseCase(metricRepository repositories.MetricRepository) contracts.CreateMetricListUseCase {

	return &createMetricListUseCase{
		metricRepository: metricRepository,
	}
}

func (c *createMetricListUseCase) Execute(ctx context.Context, dataset *input.DatasetInput) (output.AnomaliesOutput, error) {
	result := output.AnomaliesOutput{
		SiteId:                  dataset.SiteID,
		OutliersDetectionMethod: dataset.OutliersDetectionMethod,
		TimeAgo:                 dataset.TimeAgo,
		TimeStep:                dataset.TimeStep,
		DateStart:               time.Now(),
		Result:                  output.Result{},
	}

	err := c.metricRepository.AddMetric(ctx, dataset.MetricesList, nil)

	return result, err
}
