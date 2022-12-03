package di

import (
	"github.com/jhmorais/anomalies-detection/internal/domain/entities"
	"github.com/jhmorais/anomalies-detection/internal/repositories"
	"github.com/jhmorais/anomalies-detection/internal/usecases"
	"github.com/jhmorais/anomalies-detection/internal/usecases/contracts"
)

type DenpencyBuild struct {
	DB           []entities.Metric
	Repositories Repositories
	Usecases     Usecases
}

type Repositories struct {
	MetricRepository repositories.MetricRepository
}

type Usecases struct {
	AnomaliesUseCase         contracts.AnomaliesUseCase
	AverageUseCase           contracts.AverageUseCase
	DeviationUseCase         contracts.DeviationUseCase
	StandardDeviationUseCase contracts.StandardDeviationUseCase
	VarianceUseCase          contracts.VarianceUseCase
}

func NewBuild() *DenpencyBuild {

	builder := &DenpencyBuild{}

	builder = builder.buildDB().
		buildRepositories().
		buildUseCases()

	return builder
}

func (d *DenpencyBuild) buildDB() *DenpencyBuild {
	metricList := make([]entities.Metric, 0)
	d.DB = metricList
	return d
}

func (d *DenpencyBuild) buildRepositories() *DenpencyBuild {
	d.Repositories.MetricRepository = repositories.NewMetricRepository(d.DB)
	return d
}

func (d *DenpencyBuild) buildUseCases() *DenpencyBuild {
	d.Usecases.AnomaliesUseCase = usecases.NewAnomaliesUseCase(d.Repositories.MetricRepository)
	d.Usecases.AverageUseCase = usecases.NewAverageUseCase(d.Repositories.MetricRepository)
	d.Usecases.DeviationUseCase = usecases.NewDeviationUseCase(d.Repositories.MetricRepository)
	d.Usecases.StandardDeviationUseCase = usecases.NewStandardDeviationUseCase(d.Repositories.MetricRepository)
	d.Usecases.VarianceUseCase = usecases.NewVarianceUseCase(d.Repositories.MetricRepository)

	return d
}
