package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	h "github.com/alecholmez/binary/handlers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Config - global configuration object
type Config struct {
	Port int
	Host string
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.Methods("POST").HandlerFunc(h.EncryptHandler).Path("/encrypt").Name("Encrypt")
	router.Methods("POST").HandlerFunc(h.DecryptHandler).Path("/decrypt").Name("Decrypt")
	router.Methods("GET").HandlerFunc(h.ServeDocs).Path("/").Name("Documentation")

	logger := handlers.LoggingHandler(os.Stdout, router)
	handler := cors.Default().Handler(logger)

	s := http.Server{
		Addr:         fmt.Sprintf("%s:%s", "0.0.0.0", os.Getenv("PORT")),
		Handler:      handler,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	s.ListenAndServe()
}
