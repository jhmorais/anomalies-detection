package contracts

import (
	"context"

	"github.com/jhmorais/anomalies-detection/internal/usecases/ports/input"
	"github.com/jhmorais/anomalies-detection/internal/usecases/ports/output"
)

type CreateMetricListUseCase interface {
	Execute(ctx context.Context, dataset *input.DatasetInput) (output.AnomaliesOutput, error)
}
