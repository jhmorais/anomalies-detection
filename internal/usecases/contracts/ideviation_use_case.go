package contracts

import (
	"context"
)

type DeviationUseCase interface {
	Execute(ctx context.Context, average float64) ([]int, error)
}
