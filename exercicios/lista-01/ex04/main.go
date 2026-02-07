package main

// Exercício 4: JSON Echo
//
// Implemente um handler POST /echo que recebe JSON e retorna
// o mesmo JSON com um campo adicional "received": true.
//
// Requisitos:
// - Porta: 8080
// - Rota: POST /echo
// - Recebe JSON com qualquer estrutura
// - Adiciona campo "received": true ao JSON
// - Retorna o JSON modificado
// - Content-Type da resposta: application/json
// - Se o body não for JSON válido, retorna:
//   - Status: 400
//   - Body: {"error":"invalid JSON"}
//
// Exemplo de sucesso:
//   curl -X POST http://localhost:8080/echo \
//     -H "Content-Type: application/json" \
//     -d '{"message":"hello","count":42}'
//   # {"count":42,"message":"hello","received":true}
//
// Exemplo de erro:
//   curl -X POST http://localhost:8080/echo -d "not json"
//   # {"error":"invalid JSON"}

func main() {
	// TODO: Implemente aqui
}
