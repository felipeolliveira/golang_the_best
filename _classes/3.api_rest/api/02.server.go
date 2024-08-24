package api

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"time"
)

/*
  - Web Server com Go está disponível no pacote http.
    Basta criar pelo pacote `http.Server` e definir o endereço e o handler (Mux).
    Para iniciar o servidor, basta chamar o método `ListenAndServe` do servidor.
*/
func newServer() *http.Server {
	mux := http.NewServeMux()
	createRoutes(mux)

	return &http.Server{
		Addr:                         ":8080", // Addr: ":8080", comum para desenvolvimento
		Handler:                      logMiddleware(mux),
		DisableGeneralOptionsHandler: false,
		TLSConfig:                    &tls.Config{},
		ReadTimeout:                  10 * time.Second,
		// ReadHeaderTimeout:            0,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  1 * time.Minute,
		// MaxHeaderBytes:               0,
		// TLSNextProto:                 map[string]func(*http.Server, *tls.Conn, http.Handler){},
		// ConnState: func(net.Conn, http.ConnState) {
		// },
		// ErrorLog: &log.Logger{},
		// BaseContext: func(net.Listener) context.Context {
		// },
		// ConnContext: func(ctx context.Context, c net.Conn) context.Context {
		// },
	}
}

/*
- Mux (Multiplexer) é um roteador de solicitações HTTP padrão do Go.
- Ele compara a URL da solicitação com uma lista de caminhos registrados e chama o manipulador para o caminho correspondente.
*/
func createRoutes(mux *http.ServeMux) {
	// Pode ser usado o método HTTP diretamente na definição da rota, por padrão é GET
	mux.HandleFunc("GET /healthcheck", func(w http.ResponseWriter, r *http.Request) {
		// Essa é uma forma de responder, passando o writer para o método Fprintln
		fmt.Fprintln(w, "healthcheck")
	})

	// Pode receber parâmetros na URL, para isso é necessário definir o parâmetro entre chaves
	// Exemplo: /api/users/{id} => /api/users/1
	mux.HandleFunc("/api/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintln(w, "healthcheck", id)
	})
}

/*
- Middlewares podem ser usados através de uma HOC (High Order Component) para interceptar as requisições.
*/
func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		begin := time.Now()
		next.ServeHTTP(w, r)
		fmt.Println(r.Method, r.URL.String(), time.Since(begin))
	})
}

func Server() {
	server := newServer()

	if err := server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}
}
