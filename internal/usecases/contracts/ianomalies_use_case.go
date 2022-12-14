package contracts

import (
	"context"

	"github.com/jhmorais/anomalies-detection/internal/usecases/ports/input"
	"github.com/jhmorais/anomalies-detection/internal/usecases/ports/output"
)

type AnomaliesUseCase interface {
	Execute(ctx context.Context, result *output.AnomaliesOutput, parameters *input.ParametersAnomaliesInput) (*output.AnomaliesOutput, error)
}
