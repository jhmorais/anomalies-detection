package contracts

import (
	"context"
)

type VarianceUseCase interface {
	Execute(ctx context.Context, deviation []int) (int, error)
}
