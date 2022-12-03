package repositories

import (
	"context"

	"github.com/jhmorais/anomalies-detection/internal/domain/entities"
)

type MetricRepository interface {
	FindMetricByName(ctx context.Context, id string) (*entities.Device, error)
	FindMetricByValue(ctx context.Context, brand string) ([]*entities.Device, error)
	FindMetric(ctx context.Context, brand, name string) (*entities.Device, error)
	ListMetric(ctx context.Context) ([]*entities.Device, error)
}
