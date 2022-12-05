package usecases_test

import (
	"context"
	"testing"

	"github.com/jhmorais/anomalies-detection/internal/usecases"
	"github.com/jhmorais/anomalies-detection/mocks"
	"github.com/stretchr/testify/require"
)

func TestVariance(t *testing.T) {
	t.Run("when the list is empty should return an error", func(t *testing.T) {
		t.Parallel()
		metricRepositoryMock := &mocks.MetricRepository{}
		varianceUseCase := usecases.NewVarianceUseCase(metricRepositoryMock)

		deviation := []int{}

		output, err := varianceUseCase.Execute(context.Background(), deviation)
		require.Error(t, err)
		require.Equal(t, 0, output)
	})

	t.Run("when the list is empty should return an error", func(t *testing.T) {
		t.Parallel()
		metricRepositoryMock := &mocks.MetricRepository{}
		varianceUseCase := usecases.NewVarianceUseCase(metricRepositoryMock)

		deviation := []int{5, 4, 6, 8, 2}

		output, err := varianceUseCase.Execute(context.Background(), deviation)
		require.NoError(t, err)
		require.Equal(t, 29, output)
	})
}
