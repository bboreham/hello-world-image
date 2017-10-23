package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

const (
	port = 80
)

var (
	indexTmpl *template.Template
)

func init() {
	indexTmpl = template.Must(template.ParseFiles("index.template"))
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/logo.png", logoHandler)
	log.Printf("Listening on port %d", port)
	server := &http.Server{Addr: fmt.Sprintf(":%d", port)}
	server.SetKeepAlivesEnabled(false)
	log.Fatal(server.ListenAndServe())
}

func logoHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "logo.png")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		http.Error(w, "Can't get hostname", 500)
	}
	str := fmt.Sprintf("%s at %v", hostname, time.Now().Format("15:04:05"))
	indexTmpl.Execute(w, str)
}
