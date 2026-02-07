package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// Cores para output no terminal
const (
	green  = "\033[32m"
	red    = "\033[31m"
	yellow = "\033[33m"
	cyan   = "\033[36m"
	reset  = "\033[0m"
)

type checkResult struct {
	name     string
	required bool
	ok       bool
	version  string
	message  string
}

func main() {
	// Flag para especificar unidade
	unit := flag.Int("u", 1, "Unidade do curso (1, 2 ou 3)")
	flag.Parse()

	if *unit < 1 || *unit > 3 {
		fmt.Println("Erro: unidade deve ser 1, 2 ou 3")
		os.Exit(1)
	}

	fmt.Println()
	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║    DIM0547 - Verificação do Ambiente de Desenvolvimento    ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
	fmt.Println()

	fmt.Printf("%s📚 Verificando requisitos para: Unidade %d%s\n", cyan, *unit, reset)
	fmt.Println()

	// Informações do sistema
	fmt.Println("📋 Informações do Sistema")
	fmt.Println(strings.Repeat("-", 50))
	fmt.Printf("   Go Version:    %s\n", runtime.Version())
	fmt.Printf("   OS:            %s\n", runtime.GOOS)
	fmt.Printf("   Arquitetura:   %s\n", runtime.GOARCH)
	fmt.Printf("   GOPATH:        %s\n", getGoEnv("GOPATH"))
	fmt.Printf("   GOROOT:        %s\n", getGoEnv("GOROOT"))
	fmt.Println()

	// Lista de verificações baseada na unidade
	checks := []checkResult{}

	// === UNIDADE 1: Ferramentas básicas ===
	// Go (obrigatório sempre)
	checks = append(checks, checkCommand("go", []string{"version"}, "go version go1.", true))

	// Git (obrigatório sempre)
	checks = append(checks, checkCommand("git", []string{"--version"}, "git version", true))

	// golangci-lint (obrigatório sempre)
	checks = append(checks, checkCommand("golangci-lint", []string{"--version"}, "golangci-lint", true))

	// === UNIDADE 2+: Docker e banco de dados ===
	if *unit >= 2 {
		// Docker (obrigatório a partir da U2)
		checks = append(checks, checkCommand("docker", []string{"--version"}, "Docker version", true))

		// Docker Compose (obrigatório a partir da U2)
		checks = append(checks, checkCommand("docker", []string{"compose", "version"}, "Docker Compose version", true))

		// sqlc (obrigatório a partir da U2)
		checks = append(checks, checkCommand("sqlc", []string{"version"}, "", true))

		// Atlas - migrations (recomendado a partir da U2)
		checks = append(checks, checkCommand("atlas", []string{"version"}, "", false))
	} else {
		// Na U1, mostrar como recomendado para preparação futura
		checks = append(checks, checkCommand("docker", []string{"--version"}, "Docker version", false))
	}

	// === Ferramentas recomendadas (todas as unidades) ===
	// Air - hot reload (recomendado) - usa "air -v" mas mostra ASCII art
	checks = append(checks, checkCommandSpecial("air", []string{"-v"}, false))

	// swag - swagger generator (recomendado)
	checks = append(checks, checkCommand("swag", []string{"--version"}, "", false))

	// mockery - mock generator (recomendado)
	// mockery v2 usa "mockery --version" mas pode ter output verboso
	checks = append(checks, checkCommandSpecial("mockery", []string{"--version"}, false))

	// Exibe resultados
	fmt.Println("🔍 Verificação de Ferramentas")
	fmt.Println(strings.Repeat("-", 50))

	requiredOk := true
	recommendedMissing := []string{}

	for _, c := range checks {
		status := ""
		if c.ok {
			status = green + "✓" + reset
		} else if c.required {
			status = red + "✗" + reset
			requiredOk = false
		} else {
			status = yellow + "○" + reset
			recommendedMissing = append(recommendedMissing, c.name)
		}

		reqLabel := ""
		if c.required {
			reqLabel = "(obrigatório)"
		} else {
			reqLabel = "(recomendado)"
		}

		if c.ok {
			fmt.Printf("   %s %-18s %s %s\n", status, c.name, reqLabel, c.version)
		} else {
			fmt.Printf("   %s %-18s %s - %s\n", status, c.name, reqLabel, c.message)
		}
	}

	fmt.Println()

	// Verificação de conectividade com PostgreSQL (apenas U2+)
	if *unit >= 2 {
		fmt.Println("🐘 Verificação do PostgreSQL")
		fmt.Println(strings.Repeat("-", 50))
		pgCheck := checkPostgres()
		if pgCheck.ok {
			fmt.Printf("   %s✓%s PostgreSQL acessível via Docker\n", green, reset)
		} else {
			fmt.Printf("   %s○%s PostgreSQL não detectado (execute: docker compose up -d)\n", yellow, reset)
		}
		fmt.Println()
	}

	// Resumo final
	fmt.Println("📊 Resumo")
	fmt.Println(strings.Repeat("-", 50))

	if requiredOk {
		fmt.Printf("   %s✓ Todas as ferramentas obrigatórias para U%d estão instaladas!%s\n", green, *unit, reset)
	} else {
		fmt.Printf("   %s✗ Algumas ferramentas obrigatórias estão faltando.%s\n", red, reset)
		fmt.Println("     Consulte: docs/AMBIENTE.md para instruções de instalação.")
	}

	if len(recommendedMissing) > 0 {
		fmt.Printf("   %s○ Ferramentas recomendadas faltando: %s%s\n", yellow, strings.Join(recommendedMissing, ", "), reset)
		fmt.Println()
		fmt.Println("   Para instalar as ferramentas recomendadas:")
		fmt.Println()
		for _, tool := range recommendedMissing {
			switch tool {
			case "docker":
				fmt.Println("     curl -fsSL https://get.docker.com | sh")
			case "air":
				fmt.Println("     go install github.com/air-verse/air@latest")
			case "swag":
				fmt.Println("     go install github.com/swaggo/swag/cmd/swag@latest")
			case "atlas":
				fmt.Println("     curl -sSf https://atlasgo.sh | sh")
			case "mockery":
				fmt.Println("     go install github.com/vektra/mockery/v2@latest")
			case "sqlc":
				fmt.Println("     go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest")
			}
		}
	}

	// Dica para próxima unidade
	if *unit < 3 {
		fmt.Println()
		fmt.Printf("   %s💡 Dica:%s Para verificar requisitos da U%d, execute:\n", cyan, reset, *unit+1)
		fmt.Printf("      go run main.go -u %d\n", *unit+1)
	}

	fmt.Println()

	// Exit code baseado no resultado
	if !requiredOk {
		os.Exit(1)
	}
}

func checkCommand(name string, args []string, expectedOutput string, required bool) checkResult {
	result := checkResult{
		name:     name,
		required: required,
		ok:       false,
	}

	cmd := exec.Command(name, args...)
	output, err := cmd.CombinedOutput()

	if err != nil {
		result.message = "não encontrado"
		return result
	}

	outputStr := strings.TrimSpace(string(output))

	// Se não há output esperado específico, qualquer output é válido
	if expectedOutput == "" || strings.Contains(outputStr, expectedOutput) {
		result.ok = true
		// Extrai primeira linha como versão
		lines := strings.Split(outputStr, "\n")
		if len(lines) > 0 {
			result.version = strings.TrimSpace(lines[0])
			// Limita tamanho da versão para exibição
			if len(result.version) > 40 {
				result.version = result.version[:40] + "..."
			}
		}
	} else {
		result.message = "versão incompatível"
	}

	return result
}

func checkPostgres() checkResult {
	result := checkResult{
		name:     "postgres",
		required: false,
		ok:       false,
	}

	// Verifica se o container está ativo
	cmd := exec.Command("docker", "ps", "--filter", "name=postgres-dev", "--format", "{{.Names}}")
	output, err := cmd.Output()

	if err != nil || strings.TrimSpace(string(output)) == "" {
		result.message = "container não encontrado"
		return result
	}

	// Testa conexão REAL com o banco devdb
	testCmd := exec.Command("docker", "exec", "postgres-dev",
		"pg_isready", "-U", "dev", "-d", "devdb", "-q")

	if testCmd.Run() == nil {
		result.ok = true
		result.version = "conectado (devdb)"
	} else {
		result.message = "container ativo mas sem conexão ao devdb"
	}

	return result
}

func getGoEnv(key string) string {
	cmd := exec.Command("go", "env", key)
	output, err := cmd.Output()
	if err != nil {
		return "(não disponível)"
	}
	return strings.TrimSpace(string(output))
}

// checkCommandSpecial lida com ferramentas que têm output não-padrão
func checkCommandSpecial(name string, args []string, required bool) checkResult {
	result := checkResult{
		name:     name,
		required: required,
		ok:       false,
	}

	cmd := exec.Command(name, args...)
	output, err := cmd.CombinedOutput()

	if err != nil {
		// Para mockery, verificar se o comando existe mesmo com erro
		// pois ele pode retornar erro mas ainda estar instalado
		if name == "mockery" {
			// Tenta apenas verificar se o binário existe
			_, pathErr := exec.LookPath(name)
			if pathErr == nil {
				result.ok = true
				result.version = "instalado"
				return result
			}
		}
		result.message = "não encontrado"
		return result
	}

	result.ok = true
	outputStr := strings.TrimSpace(string(output))

	// Extrair versão
	lines := strings.Split(outputStr, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		// Procura por padrões de versão
		if strings.Contains(strings.ToLower(line), "version") ||
			strings.HasPrefix(line, "v") ||
			(len(line) > 0 && (line[0] >= '0' && line[0] <= '9')) {
			result.version = line
			if len(result.version) > 40 {
				result.version = result.version[:40] + "..."
			}
			return result
		}
	}

	// Se não encontrou padrão de versão, usa "instalado"
	result.version = "instalado"
	return result
}
