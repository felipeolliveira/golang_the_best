package main

import "fmt"

// Mapa Estado-Operação
//
// |-----------|-----------------------|---------------------|
// | operação  | estados               | resultado           |
// |-----------|-----------------------|---------------------|
// |           | nil                   | Block(deadlock)     |
// | Read/     | Aberto & Não Vazio    | Valor               |
// | Receive   | Aberto & Vazio        | Block(deadlock)     |
// |           | Fechado               | ZeroValue, false    |
// |           | Write Only            | Erro de compilação  |
// |-----------|-----------------------|---------------------|
// |           | nil                   | Block(deadlock)     |
// | Write/    | Aberto & Cheio        | Block(deadlock)     |
// | Send      | Aberto & Não Cheio    | Envia valor         |
// |           | Fechado               | panic!              |
// |           | Write Only            | Erro de compilação  |
// |-----------|-----------------------|---------------------|
// |           | nil                   | panic!              |
// |           | Aberto & Não Vazio    | Fecha, até drenar   |
// | close()   | Aberto & Vazio        | Fecha, ZeroValue    |
// |           | Fechado               | panic!              |
// |           | Read Only             | Erro de compilação  |
// |-----------|-----------------------|---------------------|
//
// ============================
//
// Ownership (o dono do channel)
//
// A definição do ownership ajuda a entender os resultados de cada operação realizada nos channels
// de acordo com o seu estado. Isso garante o funcionamento correto, evita panics e organiza os channels
// de forma a tornar o código mais legível:
//
// Responsabilidade do Dono:
// - Criar o channel
// - Escrever para os channels
// - Transferir a ownership
// - Fechar o channel (close)
// - Encapsular os itens acima, exportar para um Reader
//
// Responsibilidade Consumidor:
// - Consumir, obviamente...
// - Saber quando um channel fechou
// - Lidar com blocks
//
// Um padrão é escrever Senders/Writers e Receivers/Readers

func main() {
	ch := make(chan int, 1)

	ch <- 0
	n, ok := <-ch
	fmt.Println("Got value:", n, "| is this channel is closed?", ok)

	close(ch)

	n, ok = <-ch
	fmt.Println("Got value:", n, "| is this channel is closed?", ok)
}
