package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/joho/godotenv/autoload"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "using middleware and graceful shutting down")
}

func useMiddleware(mux *chi.Mux) {
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

}

func Run() error {
	mux := chi.NewRouter()
	useMiddleware(mux)

	mux.HandleFunc("/", index)

	addr := os.Getenv("ADDR")
	readTimeout, err := strconv.Atoi(os.Getenv("READ_TIMEOUT"))
	if err != nil {
		log.Printf("unable to load readTimeout. Got an err ->%v", err)
		return err
	}
	writeTimeout, err := strconv.Atoi(os.Getenv("WRITE_TIMEOUT"))
	if err != nil {
		log.Printf("unable to load writeTimeout. Got an err ->%v", err)
		return err
	}

	server := &http.Server{
		Addr: addr,
		Handler: mux,
		ReadTimeout: time.Duration(readTimeout * int(time.Second)),
		WriteTimeout: time.Duration(writeTimeout * int(time.Second)),
		MaxHeaderBytes: 1 << 20,

	}
	gracefulShutDown(server, addr)
	return nil
}

func gracefulShutDown(server *http.Server, addr string) {
		// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	// Run the server
	log.Printf("server running on %v...", addr)
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()
}
