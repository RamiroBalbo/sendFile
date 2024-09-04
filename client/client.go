package main

import (
	"encoding/json"
	"fmt"
	"github.com/RamiroBalbo/sendFile/utils/logger"
	"github.com/jlaffaye/ftp"
	"log"
	"os"
)

type Config struct {
	Ip         string `json:"ip"`
	Port       string `json:"port"`
	Usuario    string `json:"usuario"`
	Contraseña string `json:"contraseña"`
}

var ClientConfig Config

func init() {
	// Load config
	configData, err := os.ReadFile("config.json")
	if err != nil {
		logger.Fatal("No se pudo leer el archivo de configuración - %v", err.Error())
	}

	err = json.Unmarshal(configData, &ClientConfig)
	if err != nil {
		logger.Fatal("No se pudo parsear el archivo de configuración - %v", err.Error())
	}

	loggerLevel := logger.LevelInfo
	err = logger.ConfigureLogger("client.log", loggerLevel)
	if err != nil {
		fmt.Println("No se pudo crear el logger - ", err)
	}
}

func main() {
	logger.Info("Iniciando proceso sendFile")

	//Leer archivo json
	data, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatal("Error al leer el archivo de configuración - ", err)
	}

	//Deserializar json a []byte
	err = json.Unmarshal(data, &ClientConfig)
	if err != nil {
		log.Fatal("Error al parsear el archivo de configuración - ", err)
	}

	server, err := ftp.Dial(ClientConfig.Ip + ":" + ClientConfig.Port)
	if err != nil {
		logger.Fatal("Error al conectarse al servidor - %v", err.Error())
		println("Error al conectarse al servidor - %v", err.Error())
	}

	err = server.Login(ClientConfig.Usuario, ClientConfig.Contraseña)
	if err != nil {
		logger.Error("Contraseña y/o Usuario inválidos - %v", err.Error())
		server.Quit()
	}

	// Abrir un dirección local para cargar
	var address string
	print("Elije la dirección que quieres exportar: ")
	fmt.Scanln(&address)
	archivo, err := os.Open("/home/ramabalbo/" + address)
	if err != nil {
		fmt.Println("Error al abrir el archivo: "+address, err)
		return
	}
	defer archivo.Close()

	println("Subiendo /home/ramabalbo/" + address + " al servidor")
	err = server.Stor("/home/rama/rami/"+address, archivo)
	if err != nil {
		logger.Fatal("Error al subir el archivo - %v", err.Error())
	}
	logger.Info("Carga completada")

	server.Quit()
}
