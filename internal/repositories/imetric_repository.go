package repositories

import (
	"context"

	"github.com/jhmorais/anomalies-detection/internal/domain/entities"
	"github.com/jhmorais/anomalies-detection/internal/usecases/ports/input"
)

type MetricRepository interface {
	AddMetric(ctx context.Context, metric []*input.Metric, attributeParent *string) error
	CleanMetricList()
	FindMetricByName(ctx context.Context, id string) (*entities.Metric, error)
	FindMetricByValue(ctx context.Context, value float64) ([]*entities.Metric, error)
	FindMetric(ctx context.Context, value float64, name string) (*entities.Metric, error)
	ListMetric(ctx context.Context) ([]*entities.Metric, error)
}
