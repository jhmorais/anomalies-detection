package usecases

import (
	"context"
	"errors"
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

func (c *createMetricListUseCase) Execute(ctx context.Context, dataset *input.DatasetInput) (*output.AnomaliesOutput, error) {
	if dataset == nil {
		return nil, errors.New("failed, dataset is empty")
	}
	result := output.AnomaliesOutput{
		SiteId:                  dataset.SiteID,
		OutliersDetectionMethod: dataset.OutliersDetectionMethod,
		TimeAgo:                 dataset.TimeAgo,
		TimeStep:                dataset.TimeStep,
		DateStart:               time.Now(),
		Result:                  output.Result{},
	}

	err := c.metricRepository.AddMetric(ctx, dataset.MetricsList, nil)

	return &result, err
}
