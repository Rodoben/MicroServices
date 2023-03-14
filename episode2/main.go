package main

import (
	"context"
	"episode2/handler"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	fmt.Println("server started")

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	hh := handler.NewHello(l)
	gg := handler.NewGoodBye(l)
	ww := handler.NewWelcome(l)

	sm := http.NewServeMux()

	sm.Handle("/", hh)
	sm.Handle("/bye", gg)
	sm.Handle("/welcome", ww)

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		log.Println("hi log")
	})

	s := &http.Server{
		Addr:         ":9095",
		Handler:      sm,
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}

	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)

	//http.ListenAndServe(":9095", sm)

}
