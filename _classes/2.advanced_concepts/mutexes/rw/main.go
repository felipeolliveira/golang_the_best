package main

import "sync"

// O sync.RWMutex é uma ferramenta para separar os bloqueios de leitura ou escrita de uma ação
// - RLock e RUnlock -> bloqueia o valor para realizar leituras
// - Lock e Unlock -> bloqueio o valor para realizar alterações e leituras
//
// Supondo que tenha muitos acessos de leitura nesse valor e pouquissímos acessos de escrita,
// a variavel reservada pelo Mutex comum, obriga que todas as goroutines parem o acesso o que pode gerar
// uma fila de ações já que esse valor tem muitos acessos de leitura, tornando a modificação lenta, pois terá que
// esperar todas as leitura antes:
//   - [1:read]...[58:read]...[60:write]
//   - Nesse exemplo acima, a escrita terá que esperar todas as outras 59 rotinas lerem antes de escrever
//   - O RWMutex separa as açãoes de modo a priorizar a escrita quando ela aparecer, bloqueando qualquer acesso e garantindo
//     que o valor seja alterado com sucesso
//   - Todas os acessos de leitura precisam usar os metodos de leitura: RLock e RUnlock

type Account struct {
	mu      sync.RWMutex
	balance int
}

func (a *Account) Deposit(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.balance += amount
}

func (a *Account) AquireValue() int {
	a.mu.RLock()
	defer a.mu.RUnlock()

	return a.balance
}

func main() {}
