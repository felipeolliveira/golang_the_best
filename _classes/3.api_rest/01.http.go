package api_rest

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

/*
Utilizando o pacote http do Go, os requests são mais simples e não tem muitas opções
*/
func getSimple() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// data, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("getSimple", resp.Status)
}
func postSimple() {
	resp, err := http.Post("https://minhaapi.com", "application/json", nil)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// data, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("postSimple", resp.Status)
}

/*
  - Para fazer requisições mais complexas, é necessário criar uma requisição personalizada e ativamente enviar ela através de um client HTTP
    Há duas formas de fazer isso:
    1. http.NewRequest
    2. http.NewRequestWithContext
*/
func requestWithNewRequest() {
	req, err := http.NewRequest(http.MethodGet, "https://www.google.com", nil)
	if err != nil {
		panic(err)
	}

	// Aqui é possível adicionar headers, query params, etc
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// data, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("requestWithNewRequest", resp.Status)
}

func requestWithNewRequestWithContext() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://www.google.com", nil)
	if err != nil {
		panic(err)
	}

	// Aqui é possível adicionar headers, query params, etc
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// data, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("requestWithNewRequestWithContext", resp.Status)
}

func HTTP() {
	getSimple()
	postSimple()
	requestWithNewRequest()
	requestWithNewRequestWithContext()
}
