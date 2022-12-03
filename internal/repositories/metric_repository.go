package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/jhmorais/anomalies-detection/internal/domain/entities"
	"github.com/jhmorais/anomalies-detection/internal/usecases/ports/input"
)

type metricRepository struct {
	db []entities.Metric
}

func NewMetricRepository(db []entities.Metric) MetricRepository {
	return &metricRepository{db: db}
}

func (d *metricRepository) AddMetric(ctx context.Context, metrics []input.Metric) error {
	var err error

	for _, metric := range metrics {
		newMetric := entities.Metric{
			Name:      metric.Name,
			Value:     metric.Value,
			Attribute: metric.Attribute,
		}
		d.db = append(d.db, newMetric)
	}

	if len(d.db) == 0 {
		err = errors.New("failed, no metric was added")
	}
	return err
}

func (d *metricRepository) FindMetricByName(ctx context.Context, name string) (*entities.Metric, error) {
	var entity *entities.Metric
	var err error

	for _, val := range d.db {
		if val.Name == name {
			entity = &val
		}
	}
	if entity == nil {
		err = errors.New("failed, metric not found")
	}

	return entity, err
}

func (d *metricRepository) FindMetricByValue(ctx context.Context, value int) ([]*entities.Metric, error) {
	var entities []*entities.Metric
	var err error

	for _, val := range d.db {
		if val.Value == value {
			entities = append(entities, &val)
		}
	}
	if entities == nil {
		err = fmt.Errorf("failed, metrics with value: '%d' not found", value)
	}

	return entities, err
}

func (d *metricRepository) FindMetric(ctx context.Context, value int, name string) (*entities.Metric, error) {
	var entity *entities.Metric
	var err error

	for _, val := range d.db {
		if val.Name == name && val.Value == value {
			entity = &val
		}
	}
	if entity == nil {
		err = errors.New("failed, metric not found")
	}

	return entity, err
}

func (d *metricRepository) ListMetric(ctx context.Context) ([]*entities.Metric, error) {
	var entities []*entities.Metric
	var err error

	for _, val := range d.db {
		entities = append(entities, &val)
	}
	if entities == nil {
		err = fmt.Errorf("failed don't exist any metric")
	}

	return entities, err
}
