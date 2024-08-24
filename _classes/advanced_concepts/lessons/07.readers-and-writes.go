package lessons

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
  - Readers e Writers são interfaces que permitem a leitura e escrita de dados como Streams
    Para iniciar um reader ou um writer, é necessário implementar os métodos Read e Write
*/

type myWriter struct{}

func (myWriter) Write(b []byte) (n int, err error) {
	fmt.Print(string(b))
	return len(b), nil
}

func runWriter() {
	str := "Hello, World!"
	reader := strings.NewReader(str)
	buffer := make([]byte, 2)
	writer := myWriter{}
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		writer.Write(buffer[:n])
	}
	// Output: Hello, World!
}

/*
  - Os Writers e Readers também podem ser sequenciados para criar um pipeline de escrita ou leitura, ou ambos
    Basta que o próximo writer seja passado como um atributo do writer atual
*/
type myWriterWithStdout struct {
	nextWriter io.Writer
}

func (w myWriterWithStdout) Write(bytes []byte) (n int, err error) {
	for i, b := range bytes {
		bytes[i] = b + 10
	}
	return w.nextWriter.Write(bytes)
}

func runWriterNesting() {
	str := "Hello, World!"
	reader := strings.NewReader(str)
	buffer := make([]byte, 2)
	nestingWriters := myWriterWithStdout{nextWriter: os.Stdout}
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		nestingWriters.Write(buffer[:n])
	}
	// Output: Rovvy6*ay|vn+
	// Porque cada byte foi incrementado em 10
}

/*
	- O pacote io também possui funções auxiliares para criar readers e writers
			io.ReadAll(io.Reader) ([]byte, error)
			io.WriteString(io.Writer, string) (int, error)
			etc...
*/

func ReadersAndWriters() {
	runWriter()
	fmt.Println()
	runWriterNesting()
}
