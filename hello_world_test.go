package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gabrielluizsf/dll_call/handlers"
)

func TestHelloWorld(test *testing.T) {
	// Cria uma requisição HTTP falsa
	req, err := http.NewRequest("GET", "/v1/hello", nil)
	if err != nil {
		test.Fatal(err)
	}

	// Cria um ResponseRecorder (implementa http.ResponseWriter) para gravar a resposta
	rr := httptest.NewRecorder()

	// Chama a função HelloWorld passando a ResponseRecorder e a requisição falsa
	handlers.HelloWorld(rr, req)

	assertEquals(test, rr.Code, http.StatusOK)

	var response handlers.Response
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		test.Fatal(err)
	}

	expectedMessage := "Olha o console"
	assertEquals(test, response.Message, expectedMessage)
}

func assertEquals(t *testing.T, result, expected interface{}) {
	if result != expected {
		t.Errorf("Result: %v, Expected: %v", result, expected)
	}
}
