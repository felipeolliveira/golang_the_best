package main

import (
	"fmt"
	"time"
)

// Semaphore implementa um semáforo usando um channel bufferizado.
// O channel atua como um contador de recursos disponíveis.
type Semaphore struct {
	C chan struct{} // channel usado para controlar o acesso concorrente
}

// NewSemaphore cria um novo semáforo com um número máximo de acessos concorrentes.
// maxConcurrent define quantas goroutines podem executar simultaneamente.
func NewSemaphore(maxConcurrent int) *Semaphore {
	return &Semaphore{
		// make cria um channel bufferizado com capacidade maxConcurrent
		C: make(chan struct{}, maxConcurrent),
	}
}

// Acquire obtém uma permissão do semáforo.
// Se não houver permissões disponíveis, bloqueia até que uma esteja livre.
func (s *Semaphore) Acquire() {
	// Envia um valor ao canal, bloqueando se estiver cheio
	s.C <- struct{}{}
}

// Release libera uma permissão do semáforo.
// Usa select para evitar bloqueio se o semáforo já estiver vazio.
func (s *Semaphore) Release() {
	select {
	case <-s.C: // Tenta receber um valor do canal
		// Sucesso ao liberar uma permissão
	default:
		// Canal vazio, não há permissões para liberar
		fmt.Println("Nada para liberar")
	}
}

// worker representa uma unidade de trabalho que respeita o semáforo.
// id: identificador do worker
// sema: semáforo para controle de concorrência
// work: função que contém o trabalho a ser executado
func worker(id int, sema *Semaphore, work func()) {
	// Adquire uma permissão antes de iniciar o trabalho
	sema.Acquire()

	// Inicia uma nova goroutine para executar o trabalho
	go func() {
		// Garante que a permissão será liberada ao finalizar
		defer sema.Release()
		work()
	}()
}

func main() {
	// Cria um semáforo que permite apenas 3 execução simultânea
	semaphore := NewSemaphore(3)

	// Cria 30 workers
	for i := 1; i <= 30; i++ {
		worker(i, semaphore, func() {
			id := i // Captura o valor de i para uso dentro da closure
			fmt.Println("Go Routine iniciada", id)
			time.Sleep(2 * time.Second) // Simula trabalho que leva 2 segundos
			fmt.Println("Go Routine finalizada", id)
		})
	}

	// Aguarda tempo suficiente para todas as goroutines terminarem
	time.Sleep(60 * time.Second)
}
