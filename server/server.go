package server

import (
	"context"
	"fmt"
	"go-crud/routes/user"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	routes := mux.NewRouter()
	user.Routes(routes)
	return routes
}

func SetupServer(routes *mux.Router) *http.Server {
	srv := &http.Server{
		Addr:         ":8000",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      routes, // Pass our instance of gorilla/mux in.
	}
	return srv
}
func StartServer(srv *http.Server) {
	go func() {
		log.Printf(fmt.Sprintf("server starting on %s", srv.Addr))
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
}
func StopServer(wait time.Duration, srv *http.Server) {
	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
