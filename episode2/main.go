package main

import (
	"episode2/handler"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("server started")

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	fmt.Println("l:", l)
	hh := handler.NewHello(l)
	gg := handler.NewGoodBye(l)
	fmt.Println("hh:", hh)
	ww := handler.NewWelcome(l)
	sm := http.NewServeMux()

	fmt.Println("sm:", sm)
	sm.Handle("/", hh)
	sm.Handle("/bye", gg)
	sm.Handle("/welcome", ww)

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		log.Println("hi log")
	})

	http.ListenAndServe(":9095", sm)

}
