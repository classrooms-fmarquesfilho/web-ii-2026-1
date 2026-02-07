package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeRoute(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	
	http.DefaultServeMux.ServeHTTP(rec, req)
	
	res := rec.Result()
	defer res.Body.Close()
	
	if res.StatusCode != http.StatusOK {
		t.Errorf("/ esperava status 200, recebeu %d", res.StatusCode)
	}
	
	body, _ := io.ReadAll(res.Body)
	if string(body) != "Home" {
		t.Errorf("/ esperava 'Home', recebeu %q", string(body))
	}
}

func TestAboutRoute(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/about", nil)
	rec := httptest.NewRecorder()
	
	http.DefaultServeMux.ServeHTTP(rec, req)
	
	res := rec.Result()
	defer res.Body.Close()
	
	if res.StatusCode != http.StatusOK {
		t.Errorf("/about esperava status 200, recebeu %d", res.StatusCode)
	}
	
	body, _ := io.ReadAll(res.Body)
	if string(body) != "About" {
		t.Errorf("/about esperava 'About', recebeu %q", string(body))
	}
}

func TestHealthRoute(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	
	http.DefaultServeMux.ServeHTTP(rec, req)
	
	res := rec.Result()
	defer res.Body.Close()
	
	if res.StatusCode != http.StatusOK {
		t.Errorf("/health esperava status 200, recebeu %d", res.StatusCode)
	}
	
	// Verifica Content-Type
	contentType := res.Header.Get("Content-Type")
	if contentType != "application/json" && contentType != "application/json; charset=utf-8" {
		t.Errorf("/health Content-Type deveria ser application/json, recebeu %q", contentType)
	}
	
	// Verifica JSON
	var data map[string]string
	body, _ := io.ReadAll(res.Body)
	if err := json.Unmarshal(body, &data); err != nil {
		t.Errorf("/health deveria retornar JSON válido: %v", err)
		return
	}
	
	if data["status"] != "ok" {
		t.Errorf("/health esperava status='ok', recebeu %q", data["status"])
	}
}

func TestNotFoundRoute(t *testing.T) {
	routes := []string{"/xyz", "/foo", "/bar/baz", "/notexist"}
	
	for _, route := range routes {
		t.Run(route, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, route, nil)
			rec := httptest.NewRecorder()
			
			http.DefaultServeMux.ServeHTTP(rec, req)
			
			res := rec.Result()
			defer res.Body.Close()
			
			if res.StatusCode != http.StatusNotFound {
				t.Errorf("%s esperava status 404, recebeu %d", route, res.StatusCode)
			}
			
			body, _ := io.ReadAll(res.Body)
			if string(body) != "Not Found" && string(body) != "Not Found\n" {
				t.Errorf("%s esperava 'Not Found', recebeu %q", route, string(body))
			}
		})
	}
}

func init() {
	main()
}
