package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/RamiroBalbo/go-plantilla-server/utils/logger"
	"net/http"
)

type BodyRequest struct {
	Message string `json:"message"`
	Origin  string `json:"origin"`
}

func init() {
	loggerLevel := logger.LevelInfo
	err := logger.ConfigureLogger("client.log", loggerLevel)
	if err != nil {
		fmt.Println("No se pudo crear el logger - ", err)
	}
}

func main() {
	logger.Info("--- Comienzo ejecución del client ---")

	generateRequest("server", "8080")
}

func generateRequest(receiver string, port string) {
	// Defino la structura que acepta server
	receiverStruct := BodyRequest{
		Message: "Hola " + receiver,
		Origin:  "client",
	}

	// Serializo la estructura de memoria a JSON
	receiverjson, err := json.Marshal(receiverStruct)
	if err != nil {
		logger.Error("Error al serializar json - %v", err.Error())
	}
	// Convertir los bytes JSON a un io.Reader
	receiverRequest := bytes.NewBuffer(receiverjson)
	// POST
	receiverResponse, err := http.Post("http://localhost:"+port+"/server/accion", "application/json", receiverRequest)
	if err != nil {
		logger.Error("Error al conectar con server - %v", err.Error())
	} else {
		logger.Info("Conección con server sastifactoria - %v", receiverResponse.StatusCode)
	}
}
