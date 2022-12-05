package usecases_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jhmorais/anomalies-detection/internal/usecases"
	"github.com/jhmorais/anomalies-detection/internal/usecases/ports/input"
	"github.com/jhmorais/anomalies-detection/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateMetricList(t *testing.T) {
	t.Run("when received an dataset empty should return an error", func(t *testing.T) {
		t.Parallel()
		metricRepositoryMock := &mocks.MetricRepository{}
		metricRepositoryMock.On("AddMetric", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("failed"))
		createMetricListUseCase := usecases.NewCreateMetricListUseCase(metricRepositoryMock)

		output, err := createMetricListUseCase.Execute(context.Background(), nil)
		require.Error(t, err)
		require.Nil(t, output)
	})

	t.Run("when try add an empty list should return an error", func(t *testing.T) {
		t.Parallel()
		metricRepositoryMock := &mocks.MetricRepository{}
		metricRepositoryMock.On("AddMetric", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("failed"))
		createMetricListUseCase := usecases.NewCreateMetricListUseCase(metricRepositoryMock)
		dataset := input.DatasetInput{
			SiteID:                  "",
			TimeAgo:                 "30d",
			TimeStep:                "1d",
			OutliersDetectionMethod: "3-sigma",
			OutliersDetection: &input.OutliersDetectionInput{
				OutliersMultiplier:       2,
				StrongOutliersMultiplier: 3,
			},
			MetricsList: []*input.Metric{},
		}

		output, err := createMetricListUseCase.Execute(context.Background(), &dataset)
		require.Error(t, err)
		require.NotNil(t, output)
	})

	t.Run("when add an valid list should work", func(t *testing.T) {
		t.Parallel()
		dataset := input.DatasetInput{
			SiteID:                  "",
			TimeAgo:                 "30d",
			TimeStep:                "1d",
			OutliersDetectionMethod: "3-sigma",
			OutliersDetection: &input.OutliersDetectionInput{
				OutliersMultiplier:       2,
				StrongOutliersMultiplier: 3,
			},
			MetricsList: []*input.Metric{
				{
					Name:      "Visits",
					Attribute: "home",
					Value:     1000,
				},
				{
					Name:      "Visits",
					Attribute: "products",
					Value:     500,
					Child: []*input.Metric{
						{
							Name:      "Visits",
							Attribute: "shirts",
							Value:     300,
						},
						{
							Name:      "Visits",
							Attribute: "pants",
							Value:     200,
						},
					},
				},
			},
		}
		metricRepositoryMock := &mocks.MetricRepository{}
		metricRepositoryMock.On("AddMetric", mock.Anything, mock.Anything, mock.Anything).Return(nil)
		createMetricListUseCase := usecases.NewCreateMetricListUseCase(metricRepositoryMock)

		output, err := createMetricListUseCase.Execute(context.Background(), &dataset)
		require.NoError(t, err)
		require.NotNil(t, output)
	})
}
