package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup" // Pacote que permite sincronização de goroutines com tratamento de erros
)

// O errgroup é um pacote do Go que facilita a execução de múltiplas goroutines
// enquanto permite um controle de erros coordenado.
//
// Os principais benefícios do errgroup neste código são:
// 1.Execução concorrente: Permite executar múltiplas tarefas simultaneamente usando goroutines
// 2.Propagação de cancelamento: Se uma goroutine falhar, o contexto é cancelado automaticamente, sinalizando para outras goroutines pararem
// 3.Coleta de erros: O g.Wait() retorna o primeiro erro que ocorrer em qualquer das goroutines
// 4.Sincronização: Espera todas as goroutines terminarem antes de continuar

// taskWithContext simula uma tarefa que pode falhar ou ser cancelada
func taskWithContext(ctx context.Context, id int) error {
	select {
	case <-ctx.Done():
		// Se o contexto foi cancelado, retorna o erro do contexto
		return ctx.Err()
	default:
		// Se o ID for diferente de 2, a tarefa é bem sucedida após 500ms
		if id != 2 {
			time.Sleep(time.Millisecond * 500)
			fmt.Printf("Task %d completed successfully\n", id)
			return nil
		}
		// Se o ID for 2, retorna um erro
		return errors.New("task 2 failed")
	}
}

func main() {
	// Cria um novo grupo de erro com um contexto
	// Se qualquer goroutine retornar erro, o contexto será cancelado automaticamente
	g, ctx := errgroup.WithContext(context.Background())

	// Inicia 5 goroutines (0 a 4)
	for i := range 5 {
		i := i // Cria uma nova variável para cada iteração para evitar closure problems
		// g.Go adiciona uma nova goroutine ao grupo
		g.Go(func() error {
			return taskWithContext(ctx, i)
		})
	}

	// g.Wait() espera todas as goroutines terminarem e retorna o primeiro erro encontrado
	if err := g.Wait(); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	fmt.Println("Everything executed with success")
}
