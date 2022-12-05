package usecases_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jhmorais/anomalies-detection/internal/domain/entities"
	"github.com/jhmorais/anomalies-detection/internal/usecases"
	"github.com/jhmorais/anomalies-detection/internal/usecases/ports/input"
	"github.com/jhmorais/anomalies-detection/internal/usecases/ports/output"
	"github.com/jhmorais/anomalies-detection/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCheckAnomalies(t *testing.T) {
	t.Run("when the list is empty should return an error", func(t *testing.T) {
		t.Parallel()
		metricRepositoryMock := &mocks.MetricRepository{}
		metricRepositoryMock.On("ListMetric", mock.Anything).Return(nil, errors.New("failed"))
		anomaliesUseCase := usecases.NewAnomaliesUseCase(metricRepositoryMock)

		result := &output.AnomaliesOutput{}
		parameters := &input.ParametersAnomaliesInput{}

		output, err := anomaliesUseCase.Execute(context.Background(), result, parameters)
		require.Error(t, err)
		require.Nil(t, output)
	})

	t.Run("when the list is not empty should return an result", func(t *testing.T) {
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
				Value:     960,
			},
			{
				Name:      "Visits",
				Attribute: "products > shirts",
				Value:     380,
			},
			{
				Name:      "Visits",
				Attribute: "products > pants",
				Value:     200,
			},
		}
		metricRepositoryMock := &mocks.MetricRepository{}
		metricRepositoryMock.On("ListMetric", mock.Anything).Return(metrics, nil)
		metricRepositoryMock.On("CleanMetricList", mock.Anything)
		anomaliesUseCase := usecases.NewAnomaliesUseCase(metricRepositoryMock)

		result := &output.AnomaliesOutput{}
		parameters := &input.ParametersAnomaliesInput{
			StandardDeviation: 351,
			OutliersDetectionInput: &input.OutliersDetectionInput{
				OutliersMultiplier:       2,
				StrongOutliersMultiplier: 3,
			},
		}

		output, err := anomaliesUseCase.Execute(context.Background(), result, parameters)
		require.NoError(t, err)
		require.Greater(t, len(output.Result.Warnings), 0)
	})

	t.Run("when the metrics have the same value should return all as alarms", func(t *testing.T) {
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
		metricRepositoryMock.On("CleanMetricList", mock.Anything)
		anomaliesUseCase := usecases.NewAnomaliesUseCase(metricRepositoryMock)

		result := &output.AnomaliesOutput{}
		parameters := &input.ParametersAnomaliesInput{
			StandardDeviation: 0,
			OutliersDetectionInput: &input.OutliersDetectionInput{
				OutliersMultiplier:       2,
				StrongOutliersMultiplier: 3,
			},
		}

		output, err := anomaliesUseCase.Execute(context.Background(), result, parameters)
		require.NoError(t, err)
		require.Equal(t, len(output.Result.Warnings), 0)
		require.Equal(t, len(output.Result.Alarms), 3)
	})
}
