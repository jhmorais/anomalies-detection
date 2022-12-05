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

func TestDeviation(t *testing.T) {
	t.Run("when the list is empty should return an error", func(t *testing.T) {
		t.Parallel()
		metricRepositoryMock := &mocks.MetricRepository{}
		metricRepositoryMock.On("ListMetric", mock.Anything).Return(nil, errors.New("failed"))
		deviationUseCase := usecases.NewDeviationUseCase(metricRepositoryMock)

		output, err := deviationUseCase.Execute(context.Background(), float64(0))
		require.Error(t, err)
		require.Nil(t, output)
	})

	t.Run("when the list is not empty should return an deviation number", func(t *testing.T) {
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
		deviationUseCase := usecases.NewDeviationUseCase(metricRepositoryMock)

		output, err := deviationUseCase.Execute(context.Background(), float64(500))
		require.NoError(t, err)
		require.Equal(t, output[0], 500)
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
		deviationUseCase := usecases.NewDeviationUseCase(metricRepositoryMock)

		output, err := deviationUseCase.Execute(context.Background(), float64(1000))
		require.NoError(t, err)
		require.Equal(t, output[0], 0)
	})
}
