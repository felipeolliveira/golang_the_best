package main

import "mutexes/account"

// Cuidado ao abrir o mutex de forma pública se colocar ele em um promotion de struct
// Caso aconteça uma chamada de mutex.Lock() mais de uma vez, o programa entra em deadlock!
func main() {
	a := account.AccountWithPromotion{}

	// Aqui é possível chamar um mutex que deveria ser privado ao escopo do account,
	// isso porque Account tem uma promotion de sync.Mutex, permitindo com que isso aconteça.
	// Caso esse `Lock()` aconteça antes do Deposit liberar com o `Unlock()`, acontece o deadlock.
	//
	// Ou se o Lock() for chamado em lugares diferentes ao mesmo tempo, acontece o mesmo deadlock
	a.Lock()
	a.Deposit(100) // deadlock, pois o Deposit chama novamente o Lock()
	a.Unlock()

	// Definindo apenas o mutex privado, já não corre o risco disso acontecer, pois não tem como acessar
	// os metodos de lock e unlock fora do pacote
	b := account.AccountPrivate{}
	// b.mu.Lock() // Não existe
	b.Deposit(100)
	// b.mu.Unlock() // Não existe
}
