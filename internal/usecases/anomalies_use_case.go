package usecases

import (
	"context"
	"time"

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

func (c *anomaliesUseCase) Execute(ctx context.Context, result *output.AnomaliesOutput, parameters *input.ParametersAnomaliesInput) (*output.AnomaliesOutput, error) {
	metrics, err := c.metricRepository.ListMetric(ctx)
	if err != nil {
		return nil, err
	}

	warningValue := parameters.OutliersDetectionInput.OutliersMultiplier * parameters.StandardDeviation
	alarmValue := parameters.OutliersDetectionInput.StrongOutliersMultiplier * parameters.StandardDeviation

	for _, metric := range metrics {
		if metric.Value >= warningValue && metric.Value < alarmValue {
			outputWarning := output.WarningOutput{
				OutlierPeriodStart: time.Now(),
				OutlierPeriodEnd:   time.Now(),
				Metric:             metric.Name,
				Attribute:          metric.Attribute,
			}
			result.Result.Warnings = append(result.Result.Warnings, outputWarning)
		} else if metric.Value >= alarmValue {
			outputAlarm := output.AlarmOutput{
				OutlierPeriodStart: time.Now(),
				OutlierPeriodEnd:   time.Now(),
				Metric:             metric.Name,
				Attribute:          metric.Attribute,
			}
			result.Result.Alarms = append(result.Result.Alarms, outputAlarm)
		}
	}

	c.metricRepository.CleanMetricList()

	return result, nil
}
