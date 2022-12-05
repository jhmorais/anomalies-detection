package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jhmorais/anomalies-detection/internal/usecases/contracts"
	"github.com/jhmorais/anomalies-detection/internal/usecases/ports/input"
	"github.com/jhmorais/anomalies-detection/utils"
)

type Handler struct {
	WorkerPort               string
	AnomaliesUseCase         contracts.AnomaliesUseCase
	CreateMetricListUseCase  contracts.CreateMetricListUseCase
	AverageUseCase           contracts.AverageUseCase
	DeviationUseCase         contracts.DeviationUseCase
	StandardDeviationUseCase contracts.StandardDeviationUseCase
	VarianceUseCase          contracts.VarianceUseCase
}

func NewHTTPRouter(anomaliesUseCase contracts.AnomaliesUseCase,
	createMetricListUseCase contracts.CreateMetricListUseCase,
	averageUseCase contracts.AverageUseCase,
	deviationUseCase contracts.DeviationUseCase,
	standardDeviationUseCase contracts.StandardDeviationUseCase,
	varianceUseCase contracts.VarianceUseCase) *mux.Router {
	router := mux.NewRouter()
	handler := Handler{
		AnomaliesUseCase:         anomaliesUseCase,
		CreateMetricListUseCase:  createMetricListUseCase,
		AverageUseCase:           averageUseCase,
		DeviationUseCase:         deviationUseCase,
		StandardDeviationUseCase: standardDeviationUseCase,
		VarianceUseCase:          varianceUseCase,
	}
	router.UseEncodedPath()
	router.Use(utils.CommonMiddleware)

	router.HandleFunc("/anomalies", handler.CreateAnomalies).Methods(http.MethodPost)

	return router
}

func (h *Handler) CreateAnomalies(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteErrModel(w, http.StatusBadRequest, utils.NewErrorResponse("error reading request body"))
		return
	}

	datasetInput := input.DatasetInput{}
	err = json.Unmarshal(body, &datasetInput)
	if err != nil {
		utils.WriteErrModel(w, http.StatusBadRequest, utils.NewErrorResponse("failed to parse request body"))
		return
	}

	response, err := h.CreateMetricListUseCase.Execute(ctx, &datasetInput)
	if err != nil {
		utils.WriteErrModel(w, http.StatusNotFound,
			utils.NewErrorResponse(fmt.Sprintf("failed to create metric list, error: '%s'", err.Error())))
		return
	}

	average, err := h.AverageUseCase.Execute(ctx)
	if err != nil {
		utils.WriteErrModel(w, http.StatusNotFound,
			utils.NewErrorResponse(fmt.Sprintf("failed to calculate average, error: '%s'", err.Error())))
		return
	}

	deviation, err := h.DeviationUseCase.Execute(ctx, average)
	if err != nil {
		utils.WriteErrModel(w, http.StatusNotFound,
			utils.NewErrorResponse(fmt.Sprintf("failed to calculate deviation, error: '%s'", err.Error())))
		return
	}

	variance, err := h.VarianceUseCase.Execute(ctx, deviation)
	if err != nil {
		utils.WriteErrModel(w, http.StatusNotFound,
			utils.NewErrorResponse(fmt.Sprintf("failed to calculate variance, error: '%s'", err.Error())))
		return
	}
	standardDeviation := h.StandardDeviationUseCase.Execute(ctx, variance)
	parametersAnomalies := input.ParametersAnomaliesInput{
		StandardDeviation: standardDeviation,
		OutliersDetectionInput: &input.OutliersDetectionInput{
			OutliersMultiplier:       datasetInput.OutliersDetection.OutliersMultiplier,
			StrongOutliersMultiplier: datasetInput.OutliersDetection.StrongOutliersMultiplier,
		},
	}

	result, err := h.AnomaliesUseCase.Execute(ctx, response, &parametersAnomalies)
	if err != nil {
		utils.WriteErrModel(w, http.StatusNotFound,
			utils.NewErrorResponse(fmt.Sprintf("failed to create metric list, error: '%s'", err.Error())))
		return
	}

	jsonResponse, err := json.Marshal(result)
	if err != nil {
		utils.WriteErrModel(w, http.StatusInternalServerError,
			utils.NewErrorResponse("Failed to marshal anomalies output response"))
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonResponse))
}
