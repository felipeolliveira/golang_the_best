package lessons

import (
	"errors"
	"fmt"
	"math"
)

/*
	- Erros em Go são valores e não exceptions. Podem ser `extends` através de interfaces
		São tratados assim para evitar erros de runtime do tipo `panic` que podem quebrar a aplicação

	- Erros são retornados como um segundo valor de retorno, ou seja, na assinatura da função
*/

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func runDivider() {
	a := 10
	b := 0
	result, err := divide(a, b)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
}

/*
  - Erros previnem chamadas de ponteiros nulos
    Tratando esse tipo de erro antes de chamar métodos em ponteiros
*/
type User struct {
	Name     string
	LastName string
}

func (u User) GetFullName() string {
	return "${u.Name} ${u.LastName}"
}

func newUser(hasError bool) (*User, error) {
	if hasError {
		return nil, errors.New("cannot create a new user")
	}
	return &User{}, nil
}

func runNewUserWithNilPointer() {
	user, err := newUser(true)
	// user.GetFullName()
	// panic: runtime error: invalid memory address or nil pointer dereference
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(user.GetFullName())
	}
}

/*
  - Erros podem ser personalizados através da interface error
    Essa interface de Error tem apenas a função Error() string, então, qualquer struct que implemente essa função
    pode ser usada como um erro personalizado
*/
type SqrtError struct{}

func (e SqrtError) Error() string {
	return "cannot calculate the square root of a negative number"
}
func Sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, SqrtError{}
	}
	return math.Sqrt(n), nil
}
func runCustomErrorImplementation() {
	x := float64(-64)
	sqrt, err := Sqrt(x)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(sqrt)
	}
}

/*
  - Também é possível verificar se um erro é de um tipo específico, através dos metodos `As` e `Is` do pacote `errors`
    `Is` verifica se o erro é igual a outro erro, no caso de um erro iniciado pelo `errors.New()`
    `As` verifica se o erro é do tipo de um erro específico, no caso de um erro personalizado

- Erros em variaveis devem ter o prefixo `err<Foo>`

  - Quando usar um erro como `type struct` ou uma variável com `errors.New()`?
    Use `errors.New()` para erros simples e que não precisam de informações adicionais
    Use `type struct` para erros mais complexos e que precisam de informações adicionais
*/
var errCustomErrorWithNew = errors.New("custom error with errors.New()")

type CustomErrorWithStruct struct {
	msg  string
	code int
}

func (e CustomErrorWithStruct) Error() string {
	return fmt.Sprintf("%s - %d", e.msg, e.code)
}

func triggerErrorWithNew() error { return errCustomErrorWithNew }
func triggerErrorWithStruct() error {
	return CustomErrorWithStruct{msg: "error with struct", code: 12345}
}

func runCustomErrorIsOrAs() {
	err := triggerErrorWithNew()
	if err != nil && errors.Is(err, errCustomErrorWithNew) {
		fmt.Println("err is CustomErrorWithNew")
	}

	err = triggerErrorWithStruct()
	var customError CustomErrorWithStruct
	// É importante mandar a variável customError por referência de ponteiro para que o método As possa alterar seu tipo para o tipo correto e fazer a comparação
	if err != nil && errors.As(err, &customError) {
		// Nesse momento, a variável customError é do tipo CustomErrorWithStruct
		// e pode acessar seus atributos
		fmt.Println("msg: ", customError.msg, " code: ", customError.code)
	}
}

/*
  - Erros podem ser aninhados (wrapped) para adicionar informações adicionais
    Para isso, é necessário utilizar a função `fmt.Errorf()` para formatar os erros, com o verbo `%w`
    Caso não seja utilizado, o erro será apenas uma string simples e não será possível acessar as informações adicionais através dos métodos `Is` e `As`
*/
var errFirstError = errors.New("first error")

func triggerFistError() error {
	return errFirstError
}

func triggerSecondError() error {
	err := triggerFistError()
	if err != nil {
		// Adiciona informações adicionais ao erro, mantendo o tracking do erro anterior
		// Dessa forma, é possível acessar as informações adicionais através dos métodos `Is` e `As`
		return fmt.Errorf("second error: %w", err)
	}
	return nil
}

func runWrappedErrors() {
	err := triggerSecondError()
	// errors.Is() e errors.As() não funcionam com erros que não foram formatados com fmt.Errorf()
	if err != nil && errors.Is(err, errCustomErrorWithNew) {
		fmt.Println(err.Error())
	}
}

/*
	- Erros podem ser agrupados mantendo o tracking de todos os erros anteriores e não parando no primeiro erro, ou seja, executa tudo e retorna todos os erros caso aconteçam
		Para isso, é necessário usar o errors.Join() para agrupa-los
*/

var errFirstErrorToJoin = errors.New("first error to join")
var errSecondErrorToJoin = errors.New("second error to join")

func triggerFistErrorToJoin() error {
	return errFirstErrorToJoin
}
func triggerSecondErrorToJoin() error {
	return errSecondErrorToJoin
}

func runJoinErrors() {
	err1 := triggerFistErrorToJoin()
	err2 := triggerSecondErrorToJoin()
	errResult := errors.Join(err1, err2)
	if errResult != nil {
		fmt.Println(errResult.Error())
		// Possibilita também verificar se um erro específico está contido no erro agrupado
		fmt.Println(errors.Is(errResult, errFirstErrorToJoin))  // true
		fmt.Println(errors.Is(errResult, errSecondErrorToJoin)) // true
	}
}

func Errors() {
	runDivider()
	runNewUserWithNilPointer()
	runCustomErrorImplementation()
	runCustomErrorIsOrAs()
	runWrappedErrors()
	runJoinErrors()
}
