package main

import (
	"fmt"
	"github.com/RamiroBalbo/go-plantilla-server/utils/logger"
	"log"
	"net/http"
)

type BodyRequest struct {
	Message string `json:"message"`
	Origin  string `json:"origin"`
}

func init() {
	loggerLevel := logger.LevelInfo
	err := logger.ConfigureLogger("server.log", loggerLevel)
	if err != nil {
		fmt.Println("No se pudo crear el logger - ", err)
	}
}

func main() {
	serverPort := "8080"
	http.HandleFunc("/server/doSomething", doSomething)
	http.HandleFunc("/", notFound)

	logger.Info("Corriendo server en el puerto %v", serverPort)
	log.Fatal(http.ListenAndServe("localhost:"+serverPort, nil))

}

func doSomething(w http.ResponseWriter, r *http.Request) {
	logger.Info("Request recibida: %v, desde %v", r.RequestURI, r.RemoteAddr)
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Hola recibi la request"))
	if err != nil {
		logger.Error("No se pudo escribir la respuesta - %v", err.Error())
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	logger.Debug("Request inválida %v, desde %v", r.RequestURI, r.RemoteAddr)
	w.WriteHeader(http.StatusBadRequest)
	_, err := w.Write([]byte("Request inválida"))
	if err != nil {
		logger.Error("No se pudo escribir la respuesta - %v", err.Error())
	}
}
