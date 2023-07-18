package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gabrielluizsf/dll_call/dll"
)

type Response struct {
	Message string `json:"message"`
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	var response Response;
	response.Message = "Olha o console"
	dll.Dll_Call()
	// Escreve a resposta no corpo da resposta HTTP
	w.Header().Set("Content-Type", "application/json")
 	defer json.NewEncoder(w).Encode(response)
}
