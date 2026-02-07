package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEchoValidJSON(t *testing.T) {
	input := `{"message":"hello","count":42}`
	req := httptest.NewRequest(http.MethodPost, "/echo", bytes.NewBufferString(input))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	
	http.DefaultServeMux.ServeHTTP(rec, req)
	
	res := rec.Result()
	defer res.Body.Close()
	
	if res.StatusCode != http.StatusOK {
		t.Errorf("esperava status 200, recebeu %d", res.StatusCode)
	}
	
	// Verifica Content-Type
	contentType := res.Header.Get("Content-Type")
	if contentType != "application/json" && contentType != "application/json; charset=utf-8" {
		t.Errorf("Content-Type deveria ser application/json, recebeu %q", contentType)
	}
	
	// Verifica resposta
	var data map[string]interface{}
	body, _ := io.ReadAll(res.Body)
	if err := json.Unmarshal(body, &data); err != nil {
		t.Errorf("resposta deveria ser JSON válido: %v", err)
		return
	}
	
	// Verifica campos originais
	if data["message"] != "hello" {
		t.Errorf("esperava message='hello', recebeu %v", data["message"])
	}
	
	if data["count"] != float64(42) {
		t.Errorf("esperava count=42, recebeu %v", data["count"])
	}
	
	// Verifica campo adicionado
	if data["received"] != true {
		t.Errorf("esperava received=true, recebeu %v", data["received"])
	}
}

func TestEchoEmptyObject(t *testing.T) {
	input := `{}`
	req := httptest.NewRequest(http.MethodPost, "/echo", bytes.NewBufferString(input))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	
	http.DefaultServeMux.ServeHTTP(rec, req)
	
	res := rec.Result()
	defer res.Body.Close()
	
	if res.StatusCode != http.StatusOK {
		t.Errorf("esperava status 200, recebeu %d", res.StatusCode)
	}
	
	var data map[string]interface{}
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &data)
	
	if data["received"] != true {
		t.Errorf("esperava received=true, recebeu %v", data["received"])
	}
}

func TestEchoInvalidJSON(t *testing.T) {
	inputs := []string{
		"not json",
		"{invalid}",
		"",
		"[",
	}
	
	for _, input := range inputs {
		t.Run(input, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/echo", bytes.NewBufferString(input))
			rec := httptest.NewRecorder()
			
			http.DefaultServeMux.ServeHTTP(rec, req)
			
			res := rec.Result()
			defer res.Body.Close()
			
			if res.StatusCode != http.StatusBadRequest {
				t.Errorf("esperava status 400 para input inválido %q, recebeu %d", input, res.StatusCode)
			}
			
			var data map[string]string
			body, _ := io.ReadAll(res.Body)
			if err := json.Unmarshal(body, &data); err != nil {
				t.Errorf("resposta de erro deveria ser JSON válido: %v", err)
				return
			}
			
			if data["error"] != "invalid JSON" {
				t.Errorf("esperava error='invalid JSON', recebeu %q", data["error"])
			}
		})
	}
}

func TestEchoNestedJSON(t *testing.T) {
	input := `{"user":{"name":"Ana","age":25},"active":true}`
	req := httptest.NewRequest(http.MethodPost, "/echo", bytes.NewBufferString(input))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	
	http.DefaultServeMux.ServeHTTP(rec, req)
	
	res := rec.Result()
	defer res.Body.Close()
	
	if res.StatusCode != http.StatusOK {
		t.Errorf("esperava status 200, recebeu %d", res.StatusCode)
	}
	
	var data map[string]interface{}
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &data)
	
	// Verifica que estrutura aninhada foi preservada
	user, ok := data["user"].(map[string]interface{})
	if !ok {
		t.Error("campo 'user' deveria ser objeto")
		return
	}
	
	if user["name"] != "Ana" {
		t.Errorf("esperava user.name='Ana', recebeu %v", user["name"])
	}
	
	if data["received"] != true {
		t.Errorf("esperava received=true, recebeu %v", data["received"])
	}
}

func init() {
	main()
}
