package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	// Criar um servidor de teste
	// Note: o aluno deve implementar um handler que pode ser testado
	// Este teste assume que há um handler registrado para "/"
	
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	
	// Pega o handler padrão (DefaultServeMux)
	http.DefaultServeMux.ServeHTTP(rec, req)
	
	res := rec.Result()
	defer res.Body.Close()
	
	// Verifica status code
	if res.StatusCode != http.StatusOK {
		t.Errorf("esperava status 200, recebeu %d", res.StatusCode)
	}
	
	// Verifica body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("erro ao ler body: %v", err)
	}
	
	expected := "Hello, World!"
	if string(body) != expected {
		t.Errorf("esperava %q, recebeu %q", expected, string(body))
	}
}

func TestHelloWorldContentType(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	
	http.DefaultServeMux.ServeHTTP(rec, req)
	
	res := rec.Result()
	defer res.Body.Close()
	
	contentType := res.Header.Get("Content-Type")
	if contentType != "" && contentType != "text/plain" && contentType != "text/plain; charset=utf-8" {
		t.Errorf("Content-Type deveria ser text/plain, recebeu %q", contentType)
	}
}

func init() {
	// Registra o handler antes dos testes
	// O aluno deve chamar http.HandleFunc("/", ...) no main ou init
	main()
}
