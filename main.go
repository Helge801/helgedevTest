package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/updateviagit", HandleUpdate)

	srv := &http.Server{
		Handler: r,
		Addr:    ":8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

// HomeHandler handles home page requests
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<H1>Hello from the raspberry pi</H1>"))
}

func HandleUpdate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("success"))
	go update()
}

func update() {
	time.Sleep(time.Second * 3)
	exec.Command("go", "run", "~/go/src/github.com/updater/main.go")
	os.Exit(0)
}
