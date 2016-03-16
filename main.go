package main

import (
	"flag"
	"github.com/agent-pink/essocony-go/app"
	"log"
	"net/http"
)

func main() {
	on := flag.String("on", ":8080", "on")
	flag.Parse()
	http.Handle("/static/", http.FileServer(http.Dir("public")))
	http.Handle("/", app.App)
	log.Println("Listening on:", *on)
	log.Fatal(http.ListenAndServe(*on, nil))
}
