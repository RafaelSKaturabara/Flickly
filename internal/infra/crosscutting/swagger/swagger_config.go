package swagger

import (
	"crypto/sha256"
	"encoding/json"
	"flickly/docs"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// SetupSwagger configura o Swagger no aplicativo
func SetupSwagger(router *gin.Engine) {
	fmt.Println("\033[33mVerificando documentação Swagger...\033[0m")

	// SEMPRE forçar regeneração do Swagger para garantir que está atualizado
	fmt.Println("\033[33mRemovendo documentação Swagger existente...\033[0m")
	os.RemoveAll("docs")
	if err := os.MkdirAll("docs", 0755); err != nil {
		fmt.Printf("\033[31mErro ao criar diretório docs: %v\033[0m\n", err)
	}

	// Forçar regeneração do Swagger
	fmt.Println("\033[33mRegenerando documentação Swagger...\033[0m")
	regenerarSwagger()

	// Verificar se os arquivos realmente existem após regeneração
	if !fileExists("docs/swagger.json") || !fileExists("docs/docs.go") {
		fmt.Println("\033[31mERRO: Documentação Swagger não foi gerada corretamente. Tentando novamente...\033[0m")
		regenerarSwagger()

		// Verificar novamente
		if !fileExists("docs/swagger.json") {
			fmt.Println("\033[31mERRO: Falha ao gerar documentação Swagger!\033[0m")
		}
	}

	// Configurações do Swagger
	docs.SwaggerInfo.Title = "Flickly API"
	docs.SwaggerInfo.Description = "API do projeto Flickly (Atualizado em: " + time.Now().Format(time.RFC3339) + ")"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// Middleware para adicionar cabeçalhos no-cache
	noCache := func(c *gin.Context) {
		c.Header("Cache-Control", "no-cache, no-store, must-revalidate, max-age=0")
		c.Header("Pragma", "no-cache")
		c.Header("Expires", "0")
		c.Next()
	}

	// Adicionar timestamp na URL para evitar cache
	timestamp := fmt.Sprintf("%d", time.Now().Unix())

	// Rota para o swagger-ui com middleware de no-cache
	router.GET("/swagger/*any", noCache, func(c *gin.Context) {
		if c.Param("any") == "/doc.json" {
			// Servir arquivo swagger.json diretamente com cabeçalhos de cache
			c.Header("Content-Type", "application/json")
			c.File("./docs/swagger.json")
		} else {
			ginSwagger.WrapHandler(swaggerFiles.Handler)(c)
		}
	})

	// Abrir o Swagger no navegador após 2 segundos
	go abrirSwaggerNoBrowser(timestamp)
}

// regenerarSwagger executa o comando para regenerar a documentação do Swagger
func regenerarSwagger() {
	// Detectar o comando do swag
	swagCmd := ""
	if commandExists("swag") {
		swagCmd = "swag"
	} else if envGoPath := os.Getenv("GOPATH"); envGoPath != "" {
		possiblePath := fmt.Sprintf("%s/bin/swag", envGoPath)
		if fileExists(possiblePath) {
			swagCmd = possiblePath
		}
	} else if goPath, err := exec.Command("go", "env", "GOPATH").Output(); err == nil {
		possiblePath := fmt.Sprintf("%s/bin/swag", strings.TrimSpace(string(goPath)))
		if fileExists(possiblePath) {
			swagCmd = possiblePath
		}
	}

	if swagCmd == "" {
		fmt.Println("\033[31mErro: Não foi possível encontrar o comando swag. Instalando...\033[0m")
		cmd := exec.Command("go", "install", "github.com/swaggo/swag/cmd/swag@latest")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Printf("\033[31mErro ao instalar swag: %v\033[0m\n", err)
			return
		}

		// Verificar novamente após instalação
		if commandExists("swag") {
			swagCmd = "swag"
		} else if goPath, err := exec.Command("go", "env", "GOPATH").Output(); err == nil {
			possiblePath := fmt.Sprintf("%s/bin/swag", strings.TrimSpace(string(goPath)))
			if fileExists(possiblePath) {
				swagCmd = possiblePath
			} else {
				fmt.Println("\033[31mErro: Não foi possível encontrar o comando swag após instalação.\033[0m")
				return
			}
		}
	}

	// Criar diretório docs se não existir
	if err := os.MkdirAll("docs", 0755); err != nil {
		fmt.Printf("\033[31mErro ao criar diretório docs: %v\033[0m\n", err)
		return
	}

	// Executar o comando swag com opções para forçar a regeneração
	fmt.Println("\033[33mGerando nova documentação Swagger...\033[0m")
	cmd := exec.Command(swagCmd, "init", "-g", "cmd/flickly/main.go", "--parseDependency", "--parseInternal", "--overridesFile", "")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("\033[31mErro ao regenerar documentação Swagger: %v\033[0m\n", err)
	} else {
		// Adicionar timestamp nos arquivos gerados para forçar mudança de hash
		adicionarTimestampNoJSON("docs/swagger.json")
		fmt.Println("\033[32mDocumentação Swagger regenerada com sucesso!\033[0m")
	}
}

// adicionarTimestampNoJSON adiciona um timestamp nos arquivos JSON para forçar mudança de hash
func adicionarTimestampNoJSON(filePath string) {
	// Ler o arquivo JSON
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("\033[31mErro ao ler arquivo %s: %v\033[0m\n", filePath, err)
		return
	}

	// Converter para mapa
	var jsonData map[string]interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		fmt.Printf("\033[31mErro ao fazer parse do JSON: %v\033[0m\n", err)
		return
	}

	// Adicionar timestamp no campo info.description
	info, ok := jsonData["info"].(map[string]interface{})
	if ok {
		ts := time.Now().Format(time.RFC3339)
		if desc, ok := info["description"].(string); ok {
			// Remove timestamp anterior, se existir
			if idx := strings.Index(desc, " (Atualizado em:"); idx > 0 {
				desc = desc[:idx]
			}
			info["description"] = fmt.Sprintf("%s (Atualizado em: %s)", desc, ts)
		}
	}

	// Converter de volta para JSON com indentação
	newData, err := json.MarshalIndent(jsonData, "", "    ")
	if err != nil {
		fmt.Printf("\033[31mErro ao converter para JSON: %v\033[0m\n", err)
		return
	}

	// Escrever de volta no arquivo
	if err := os.WriteFile(filePath, newData, 0644); err != nil {
		fmt.Printf("\033[31mErro ao escrever no arquivo %s: %v\033[0m\n", filePath, err)
		return
	}
}

// commandExists verifica se um comando existe no PATH
func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

// fileExists verifica se um arquivo existe
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// isSwaggerDesatualizado verifica se o Swagger precisa ser atualizado
func isSwaggerDesatualizado() bool {
	// Verificar se o arquivo da documentação existe
	_, err := os.Stat("docs/swagger.json")
	if os.IsNotExist(err) {
		return true
	}

	// Verificar as datas de modificação dos arquivos importantes
	swaggerInfo, err := os.Stat("docs/swagger.json")
	if err != nil {
		return true
	}

	swaggerTime := swaggerInfo.ModTime()

	// Lista de arquivos e diretórios a verificar
	arquivosPrincipais := []string{
		"cmd/flickly/main.go",
		"internal/api/flickly/router.go",
		"internal/api/users/controllers/user_controller.go",
		"internal/infra/crosscutting/swagger/swagger_config.go",
	}

	// Verificar arquivos específicos
	for _, arquivo := range arquivosPrincipais {
		if fileInfo, err := os.Stat(arquivo); err == nil {
			if fileInfo.ModTime().After(swaggerTime) {
				fmt.Printf("\033[33mArquivo modificado: %s\033[0m\n", arquivo)
				return true
			}
		}
	}

	// Diretórios principais para verificar
	diretorios := []string{
		"internal/domain/users/commands",
		"internal/domain/users/entities",
		"internal/api/users/viewmodels",
		"internal/domain/core",
		"internal/api",
	}

	for _, dir := range diretorios {
		if diretorioModificadoDepois(dir, swaggerTime) {
			return true
		}
	}

	return false
}

// diretorioModificadoDepois verifica se algum arquivo .go em um diretório foi modificado após um determinado tempo
func diretorioModificadoDepois(dir string, referenceTime time.Time) bool {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".go") && info.ModTime().After(referenceTime) {
			return fmt.Errorf("found newer file")
		}
		return nil
	})

	// Se encontrou um arquivo mais recente, retorna true
	return err != nil
}

// verificarSwaggerExistente apenas verifica se os arquivos Swagger existem
func verificarSwaggerExistente() {
	_, err := os.Stat("docs/swagger.json")
	if os.IsNotExist(err) {
		fmt.Println("\033[33mAVISO: Documentação Swagger não encontrada.\033[0m")
		fmt.Println("\033[33mExecute 'make swagger' ou './update_swagger.sh' para gerar a documentação.\033[0m")
	}
}

// abrirSwaggerNoBrowser abre o Swagger UI no navegador padrão
func abrirSwaggerNoBrowser(timestamp string) {
	time.Sleep(2 * time.Second)
	swaggerURL := fmt.Sprintf("http://localhost:8080/swagger/index.html?v=%s", timestamp)
	fmt.Printf("Abrindo Swagger UI em: %s\n", swaggerURL)
	err := abrirNavegador(swaggerURL)
	if err != nil {
		fmt.Printf("Erro ao abrir o navegador: %v\n", err)
	}
}

// abrirNavegador abre uma URL no navegador padrão do sistema
func abrirNavegador(url string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin": // macOS
		cmd = exec.Command("open", url)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	default:
		return fmt.Errorf("sistema operacional não suportado para abrir navegador automaticamente")
	}

	return cmd.Start()
}

// gerarHashTipos gera um hash dos tipos atuais na aplicação
func gerarHashTipos() string {
	// Lista de diretórios que contêm modelos ou comandos relevantes para o Swagger
	diretoriosChave := []string{
		"internal/domain/users/commands",
		"internal/domain/users/entities",
		"internal/api/users/viewmodels",
		"internal/domain/core",
		"internal/api",
	}

	hasher := sha256.New()

	// Percorrer todos os diretórios relevantes
	for _, dir := range diretoriosChave {
		if err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil // Ignorar erros
			}

			// Apenas arquivos Go
			if !info.IsDir() && strings.HasSuffix(path, ".go") {
				conteudo, err := os.ReadFile(path)
				if err != nil {
					return nil
				}

				// Adicionar o conteúdo do arquivo ao hasher
				hasher.Write(conteudo)
			}
			return nil
		}); err != nil {
			fmt.Printf("\033[33mAviso ao processar diretório %s: %v\033[0m\n", dir, err)
		}
	}

	// Gerar hash em formato hexadecimal
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

// verificarTiposAlterados verifica se os tipos da aplicação foram alterados desde a última geração
func verificarTiposAlterados() bool {
	// Gerar hash do estado atual
	hashAtual := gerarHashTipos()

	// Verificar se existe um arquivo com o hash anterior
	hashAnterior := ""
	hashFile := "docs/types_hash.txt"

	if fileExists(hashFile) {
		data, err := os.ReadFile(hashFile)
		if err == nil {
			hashAnterior = strings.TrimSpace(string(data))
		}
	}

	// Se o hash mudou, atualizar o arquivo e retornar true
	if hashAnterior != hashAtual {
		if err := os.WriteFile(hashFile, []byte(hashAtual), 0644); err != nil {
			fmt.Printf("\033[31mErro ao salvar hash de tipos: %v\033[0m\n", err)
		}
		return true
	}

	return false
}
