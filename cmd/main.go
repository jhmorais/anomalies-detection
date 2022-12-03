package main

import (
	"fmt"
	"net/http"

	"github.com/jhmorais/anomalies-detection/config"
	"github.com/jhmorais/anomalies-detection/internal/infra/di"
	"github.com/jhmorais/anomalies-detection/services"
)

func main() {
	config.LoadServerEnvironmentVars()

	dependencies := di.NewBuild()

	router := services.NewHTTPRouter(dependencies.Usecases.AnomaliesUseCase)

	fmt.Println("Starting SERVER, LISTEN PORT: " + config.GetServerPort())
	deviceErr := http.ListenAndServe(fmt.Sprintf(":%s", config.GetServerPort()), router)
	if deviceErr != nil && deviceErr != http.ErrServerClosed {
		fmt.Println("failed to create server rest on port: " + config.GetServerPort())
		fmt.Println(deviceErr.Error())
	}
}
