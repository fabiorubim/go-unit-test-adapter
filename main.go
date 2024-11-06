package main

import (
	"fmt"
	"io"
	"net/http"
)

// Primeiro, definimos uma interface que representará as operações que queremos fazer com o cliente HTTP:
// HTTPClient é uma interface que define o método Do
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// =============================
// ### Passo 2: Implementar o adapter para o http.Client
// Agora implementamos o adapter que implementa a interface `HTTPClient` usando `http.Client`.
// HTTPClientAdapter é um adapter que usa http.Client
type HTTPClientAdapter struct {
	client *http.Client // Referência ao cliente HTTP real
}

// Do faz a requisição HTTP e retorna a resposta ou um erro
func (hca *HTTPClientAdapter) Do(req *http.Request) (*http.Response, error) {
	return hca.client.Do(req) // Despacha a requisição usando o cliente real
}

// =============================

// ### Passo 3: Implementar o Mock
// Agora, podemos criar um mock para a interface `HTTPClient`. Este mock simulará o comportamento do `http.Client`,
// permitindo que você defina respostas predefinidas.
// MockHTTPClient é um mock do HTTPClient
type MockHTTPClient struct {
	Response *http.Response // Resposta simulada
	Err      error          // Erro simulado
}

// Do simula o método Do do HTTPClient
func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.Response, m.Err // Retorna a resposta ou erro simulado
}

// =============================
// ### Passo 4: Usar o Adapter e o Mock em seu código
// Aqui está um exemplo de como usar o adapter e o mock em seu código.
// Suponha que temos uma função que faz uma solicitação GET.
// FetchData é uma função que busca dados de uma URL usando o HTTPClient
func FetchData(client HTTPClient, url string) (string, error) {
	// Cria uma nova requisição HTTP GET
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err // Retorna o erro se a criação da requisição falhar
	}

	// Executa a requisição usando o cliente HTTP fornecido
	resp, err := client.Do(req)
	if err != nil {
		return "", err // Retorna o erro se a requisição falhar
	}
	defer resp.Body.Close() // Fecha o corpo da resposta após ler

	// Lê o corpo da resposta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err // Retorna o erro se a leitura falhar
	}

	return string(body), nil // Retorna o corpo como string
}

func main() {
	// Usar o adapter com http.Client real
	client := &http.Client{}
	adapter := &HTTPClientAdapter{client: client}
	response, err := FetchData(adapter, "https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Erro:", err) // Imprime o erro se houver
	} else {
		fmt.Println("Resposta:", response) // Imprime a resposta recebida
	}
}
