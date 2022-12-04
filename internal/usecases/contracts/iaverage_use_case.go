package contracts

import (
	"context"
)

type AverageUseCase interface {
	Execute(ctx context.Context) (average float64, err error)
}
