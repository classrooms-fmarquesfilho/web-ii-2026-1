package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGreetWithName(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/greet?name=Maria", nil)
	rec := httptest.NewRecorder()
	
	http.DefaultServeMux.ServeHTTP(rec, req)
	
	res := rec.Result()
	defer res.Body.Close()
	
	if res.StatusCode != http.StatusOK {
		t.Errorf("esperava status 200, recebeu %d", res.StatusCode)
	}
	
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("erro ao ler body: %v", err)
	}
	
	expected := "Hello, Maria!"
	if string(body) != expected {
		t.Errorf("esperava %q, recebeu %q", expected, string(body))
	}
}

func TestGreetWithoutName(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/greet", nil)
	rec := httptest.NewRecorder()
	
	http.DefaultServeMux.ServeHTTP(rec, req)
	
	res := rec.Result()
	defer res.Body.Close()
	
	if res.StatusCode != http.StatusOK {
		t.Errorf("esperava status 200, recebeu %d", res.StatusCode)
	}
	
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("erro ao ler body: %v", err)
	}
	
	expected := "Hello, World!"
	if string(body) != expected {
		t.Errorf("esperava %q, recebeu %q", expected, string(body))
	}
}

func TestGreetWithEmptyName(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/greet?name=", nil)
	rec := httptest.NewRecorder()
	
	http.DefaultServeMux.ServeHTTP(rec, req)
	
	res := rec.Result()
	defer res.Body.Close()
	
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("erro ao ler body: %v", err)
	}
	
	expected := "Hello, World!"
	if string(body) != expected {
		t.Errorf("esperava %q para name vazio, recebeu %q", expected, string(body))
	}
}

func TestGreetWithDifferentNames(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{"João", "Hello, João!"},
		{"Ana", "Hello, Ana!"},
		{"Carlos", "Hello, Carlos!"},
	}
	
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/greet?name="+tc.name, nil)
			rec := httptest.NewRecorder()
			
			http.DefaultServeMux.ServeHTTP(rec, req)
			
			res := rec.Result()
			defer res.Body.Close()
			
			body, _ := io.ReadAll(res.Body)
			if string(body) != tc.expected {
				t.Errorf("esperava %q, recebeu %q", tc.expected, string(body))
			}
		})
	}
}

func init() {
	main()
}
