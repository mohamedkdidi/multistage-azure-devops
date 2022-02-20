// Go web service.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	addr = flag.String("listen", ":8080", "http server address")
)

func main() {
	flag.Parse()

	app := App{
		Addr: *addr,
		Logf: log.Printf,
	}

	host, port := os.Getenv("HOST"), os.Getenv("PORT")
	if host != "" || port != "" {
		app.Addr = host + ":" + port
	}

	app.Run()
}

type App struct {
	Addr string
	Logf func(string, ...interface{})
}

func (app *App) Run() {
	http.Handle("/", app)
	http.HandleFunc("/shutdown", app.Shutdown)

	app.Logf("Starting on: %s", app.Addr)
	http.ListenAndServe(app.Addr, nil)
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello: %s", r.URL.Path)
}

func (app *App) Shutdown(w http.ResponseWriter, r *http.Request) {
	app.Logf("Shutting down...")
	os.Exit(0)
}
