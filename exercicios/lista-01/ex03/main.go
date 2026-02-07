package main

// Exercício 3: Múltiplas Rotas
//
// Implemente um servidor com múltiplas rotas usando http.ServeMux.
//
// Requisitos:
// - Porta: 8080
// - Rotas:
//   - GET /       → retorna "Home" (text/plain)
//   - GET /about  → retorna "About" (text/plain)
//   - GET /health → retorna {"status":"ok"} (application/json)
//   - Outras rotas → retorna "Not Found" com status 404
//
// Exemplos:
//   curl http://localhost:8080/
//   # Home
//
//   curl http://localhost:8080/about
//   # About
//
//   curl http://localhost:8080/health
//   # {"status":"ok"}
//
//   curl http://localhost:8080/xyz
//   # Not Found (status 404)

func main() {
	// TODO: Implemente aqui
}
