package repositories

import (
	"context"

	"github.com/jhmorais/anomalies-detection/internal/domain/entities"
)

type MetricRepository interface {
	FindMetricByName(ctx context.Context, id string) (*entities.Metric, error)
	FindMetricByValue(ctx context.Context, value int) ([]*entities.Metric, error)
	FindMetric(ctx context.Context, value int, name string) (*entities.Metric, error)
	ListMetric(ctx context.Context) ([]*entities.Metric, error)
}
