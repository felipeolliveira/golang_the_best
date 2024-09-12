# Palindrome Checker (Spacelaxy Challenge)

Este projeto é um verificador de palíndromos que permite ao usuário inserir uma palavra, frase ou sequência de números e verificar se é um palíndromo.

## Estrutura do Projeto

O projeto está organizado da seguinte forma:


### Arquivos

- `main.go`: Contém a função principal que gerencia o fluxo do programa.
- `check_palindrome.go`: Contém a função `checkPalindrome` que verifica se a string é um palíndromo.
- `input.go`: Contém a função `scanUserInput` que lê a entrada do usuário.
- `sanitize_string.go`: Contém a função `sanitizeString` que normaliza e sanitiza a string de entrada.

## Como Executar o binário
Compilei o programa para **Windows** e **Linux**
```sh
# Linux
./is_palindrome_amd64

#Windows 64 bits
./is_palindrome_win64.exe
```

## Como Executar o Projeto

1. Certifique-se de ter o Go instalado.
2. Navegue até o diretório challenges/06.palindrome/.
3. Execute o comando:
```sh
# instalar as dependencias caso não tenha
go mod tidy

# executar o programa
go run .
```
4. Siga as instruções no terminal para inserir uma string e verificar se é um palíndromo.

## Exemplo de uso
```txt
======================
== É um palindromo? ==
======================
Digite uma palavra, frase ou uma sequencia números:
A man a plan a canal Panama
Sim, é um palíndromo!
======================
Você gostaria de tentar outra palavra? (Y/n):
```
### Contribuição
Sinta-se à vontade para contribuir com melhorias ou correções. Para isso, faça um fork do repositório, crie uma branch para suas alterações e envie um pull request.

## Licença
Este projeto está licenciado sob a licença MIT
