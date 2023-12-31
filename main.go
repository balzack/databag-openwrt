package main

import (
	app "databag/internal"
	"databag/internal/store"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
  "os"
)

func main() {

  args := os.Args
  if len(args) == 3 {
    port := ":" + args[2]
    store.SetPath("/var/lib/databag/databag.db")
    router := app.NewRouter(args[3])
    origins := handlers.AllowedOrigins([]string{"*"})
    methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
    log.Fatal(http.ListenAndServe(port, handlers.CORS(origins, methods)(router)))
  }
  else {
    log.Printf("usage: databag <port> <web path>");
  }
}
