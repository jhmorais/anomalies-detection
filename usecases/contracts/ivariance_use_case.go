package contracts

import (
	"context"
)

type VarianceDeviationUseCase interface {
	Execute(ctx context.Context, deviation []int) int
}
