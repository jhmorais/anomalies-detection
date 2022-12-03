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
	WorkerPort       string
	AnomaliesUseCase contracts.AnomaliesUseCase
}

func NewHTTPRouter(anomaliesUseCase contracts.AnomaliesUseCase) *mux.Router {
	router := mux.NewRouter()
	handler := Handler{
		AnomaliesUseCase: anomaliesUseCase,
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

	device := input.DatasetInput{}
	err = json.Unmarshal(body, &device)
	if err != nil {
		utils.WriteErrModel(w, http.StatusBadRequest, utils.NewErrorResponse("failed to parse request body"))
		return
	}

	response, err := h.AnomaliesUseCase.Execute(ctx, &device)
	if err != nil {
		utils.WriteErrModel(w, http.StatusNotFound,
			utils.NewErrorResponse(fmt.Sprintf("failed to create device, error: '%s'", err.Error())))
		return
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		utils.WriteErrModel(w, http.StatusInternalServerError,
			utils.NewErrorResponse("Failed to marshal device response"))
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonResponse))
}
