package contracts

import (
	"context"

	"github.com/jhmorais/anomalies-detection/internal/usecases/ports/input"
)

type AverageUseCase interface {
	Execute(ctx context.Context, metric []*input.Metric) (average int, valuesMaps map[int]int)
}
