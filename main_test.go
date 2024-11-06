package main

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

// ### Passo 5: Configurar testes utilizando o Mock
// Agora você pode escrever testes para `FetchData` usando o `MockHTTPClient`.
// TestFetchData testa a função FetchData usando um MockHTTPClient
func TestFetchData(t *testing.T) {
	// Preparar resposta simulada
	mockResponse := &http.Response{
		StatusCode: http.StatusOK,                                              // Código de status HTTP 200
		Body:       io.NopCloser(strings.NewReader(`{"title": "Mock Title"}`)), // Corpo da resposta simulada
	}

	// Cria um cliente mock que retornará a resposta simulada
	mockClient := &MockHTTPClient{
		Response: mockResponse,
		Err:      nil, // Sem erro simulado
	}

	// Chama FetchData usando o cliente mock
	data, err := FetchData(mockClient, "http://fakeurl.com")
	if err != nil {
		t.Fatalf("Esperava não ter erro, mas obteve: %v", err) // Falha no teste se houver erro
	}

	expected := `{"title": "Mock Title"}` // Resposta esperada
	if data != expected {
		t.Errorf("Esperava '%s', mas obteve: %s", expected, data) // Falha no teste se o valor não corresponder
	}
}
