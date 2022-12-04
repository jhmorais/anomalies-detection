package contracts

import (
	"context"
)

type StandardDeviationUseCase interface {
	Execute(ctx context.Context, variance int) float64
}
