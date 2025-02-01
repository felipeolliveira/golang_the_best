package main

import (
	"log/slog"
	"mailer_service/internal/mailer"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	var wg sync.WaitGroup
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	s := mailer.NewSender("sandbox.smtp.mailtrap.io", 2525, "ec960ee410999e", "822beebe2d70cd", "contact@xablau.com")

	r.Route("/api", func(r chi.Router) {
		r.Route("/mail", func(r chi.Router) {
			r.Route("/send", func(r chi.Router) {
				r.Post("/{email}", func(w http.ResponseWriter, r *http.Request) {
					userMail := chi.URLParam(r, "email")

					wg.Add(1)
					go s.Send(userMail, &wg)
				})
			})
		})
	})

	go http.ListenAndServe("localhost:3042", r)

	// O bloco abaixo espera o sinal de fechamento do programa atraves da syscall SIGTERM(pelo OS geralmente) ou SIGINT(pelo usuário - ctrl-C)
	// bloqueiando a execução do Wait() até que o programa ser notificado
	// Assim, se houver mais goroutines executando a rota, todas elas precisam terminar antes da main ser finalizada
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	slog.Info("main", "stopping", "Got a signal... Cleaning Up")

	wg.Wait()
}
