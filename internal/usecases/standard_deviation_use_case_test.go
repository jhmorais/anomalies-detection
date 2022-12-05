package usecases_test

import (
	"context"
	"testing"

	"github.com/jhmorais/anomalies-detection/internal/usecases"
	"github.com/jhmorais/anomalies-detection/mocks"
	"github.com/stretchr/testify/require"
)

func TestStandardDeviation(t *testing.T) {
	t.Run("when the variance is sent should return the square root", func(t *testing.T) {
		t.Parallel()
		metricRepositoryMock := &mocks.MetricRepository{}
		standardDeviationUseCase := usecases.NewStandardDeviationUseCase(metricRepositoryMock)

		output := standardDeviationUseCase.Execute(context.Background(), 4)
		require.Equal(t, output, float64(2))
	})

	t.Run("when the variance is zero should return 0", func(t *testing.T) {
		t.Parallel()
		metricRepositoryMock := &mocks.MetricRepository{}
		standardDeviationUseCase := usecases.NewStandardDeviationUseCase(metricRepositoryMock)

		output := standardDeviationUseCase.Execute(context.Background(), 0)
		require.Equal(t, output, float64(0))
	})
}
