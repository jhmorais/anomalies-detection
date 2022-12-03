package contracts

import (
	"context"

	"github.com/jhmorais/anomalies-detection/internal/usecases/ports/input"
)

type DeviationUseCase interface {
	Execute(ctx context.Context, average int, metrics []*input.Metric) []int
}
