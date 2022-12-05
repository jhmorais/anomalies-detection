package usecases_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jhmorais/anomalies-detection/internal/domain/entities"
	"github.com/jhmorais/anomalies-detection/internal/usecases"
	"github.com/jhmorais/anomalies-detection/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestAverage(t *testing.T) {
	t.Run("when the list is empty should return an error", func(t *testing.T) {
		t.Parallel()
		metricRepositoryMock := &mocks.MetricRepository{}
		metricRepositoryMock.On("ListMetric", mock.Anything).Return(nil, errors.New("failed"))
		averageUseCase := usecases.NewAverageUseCase(metricRepositoryMock)

		output, err := averageUseCase.Execute(context.Background())
		require.Error(t, err)
		require.Equal(t, output, float64(0))
	})

	t.Run("when the list is not empty should return an average number", func(t *testing.T) {
		t.Parallel()
		metrics := []*entities.Metric{
			{
				Name:      "Visits",
				Attribute: "home",
				Value:     1000,
			},
			{
				Name:      "Visits",
				Attribute: "products",
				Value:     500,
			},
			{
				Name:      "Visits",
				Attribute: "products > shirts",
				Value:     300,
			},
			{
				Name:      "Visits",
				Attribute: "products > pants",
				Value:     200,
			},
		}
		metricRepositoryMock := &mocks.MetricRepository{}
		metricRepositoryMock.On("ListMetric", mock.Anything).Return(metrics, nil)
		averageUseCase := usecases.NewAverageUseCase(metricRepositoryMock)

		output, err := averageUseCase.Execute(context.Background())
		require.NoError(t, err)
		require.Equal(t, output, float64(500))
	})

	t.Run("when the metrics have the same value should return the same value of metrics", func(t *testing.T) {
		t.Parallel()
		metrics := []*entities.Metric{
			{
				Name:      "Visits",
				Attribute: "home",
				Value:     1000,
			},
			{
				Name:      "Visits",
				Attribute: "products",
				Value:     1000,
			},
			{
				Name:      "Visits",
				Attribute: "products > shirts",
				Value:     1000,
			},
		}
		metricRepositoryMock := &mocks.MetricRepository{}
		metricRepositoryMock.On("ListMetric", mock.Anything).Return(metrics, nil)
		averageUseCase := usecases.NewAverageUseCase(metricRepositoryMock)

		output, err := averageUseCase.Execute(context.Background())
		require.NoError(t, err)
		require.Equal(t, output, float64(1000))
	})
}
