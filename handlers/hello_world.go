package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"syscall"
)

type Response struct {
	Message string `json:"message"`
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	// Define qual DLL será chamada
	dll := syscall.NewLazyDLL("./dll/hello_Dll.dll")

	// Procura a função na DLL pelo nome
	hello := dll.NewProc("helloWorld")

	var response Response

	// Chama a função da DLL
	_, _, err := hello.Call()
	if err == nil{
		log.Fatalln(err)
	}

	response.Message = "Olha o console"

	// Escreve a resposta no corpo da resposta HTTP
	w.Header().Set("Content-Type", "application/json")
 	defer json.NewEncoder(w).Encode(response)
}
