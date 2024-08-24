package lessons

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"
)

/*
- Desde a versão 1.18 foi implementado um novo pacote de log, o slog
- O slog é um pacote de log que é compatível com o pacote log padrão, porém ganha o adicional de suportar campos estruturados
*/
func simpleLogs() {
	slog.Info("Serviço iniciado", "version", "0.1.0")
	slog.Warn("Serviço pausado", "reason", "falta de memória")
	slog.Error("Serviço parado", "reason", "erro no banco de dados")
}

/*
  - É possível criar um logger customizado com o slog, passando um handler e uma saída para o log
    Quando criado, é possível definir o logger como padrão do slog, ou então utilizar o logger criado pela variavel
*/
func createLogger() {
	l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	l.Info("Serviço iniciado com o novo logger", "version", "0.1.0")
	// slog.SetDefault(l)
}

/*
  - Para garantir a estruturação dos logs, é possível passar os campos como argumentos para o slog
    Porém o padrão não faz checagem de tipos, o que pode gerar um campo que esperava receber um int, receber uma string, por exemplo
    Nesses casos, pode ser interessante usar duas soluções:
    1. Usar os métodos de criação de campos do slog: slog.Int("status", 10); slog.String("version", "1.0.0")
    2. Usar o LogAttrs (Log Attributes) que cria um mapa de campos que aceita tipagem por interface

    É possível também passar uma gama de opções para o handler, como adicionar o source do log, definir um level padrão, ou substituir atributos
*/
func logsWithFields() {
	options := &slog.HandlerOptions{
		AddSource:   true, // Adiciona o source do log: arquivo e linha, qual função chamou o log, etc
		Level:       nil,  // Define o level padrão mínimo do log, se não for passado, o padrão é Info
		ReplaceAttr: nil,
	}

	l := slog.New(slog.NewJSONHandler(os.Stdout, options))
	l.Info("request on",
		"method", http.MethodDelete,
		"status", 10,
		"boolean", "1.0.0", // Não conferencia de tipos, pode se tornar um problema
		slog.Int64("id", 1), // Usando o método de criação de campos do slog
	)

	// Com o metodo With, é possível adicionar campos comuns a todos os logs depois da declaração
	// Isso é útil para adicionar campos como versão da aplicação, ambiente, etc
	// O With retorna um novo logger, então é necessário atribuir o retorno a uma variável
	l = l.With(slog.Group("app_info", "version", "1.0.0"))

	// Usando o LogAttrs, que aceita tipagem por interface e garante todos os tipos dentro do log
	l.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		"request on",
		slog.Group("request", // Agrupando os campos em um grupo
			slog.String("method", http.MethodDelete),
			slog.Int64("status", http.StatusOK),
		),
		slog.Duration("duration", 10*time.Second),
		slog.String("user_agent", "Mozilla/5.0"),
	)
}

func Logs() {
	simpleLogs()
	createLogger()
	logsWithFields()
}
