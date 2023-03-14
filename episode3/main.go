package main

import (
	"context"
	"episode3/handler"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	c := make(chan os.Signal, log.LstdFlags)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)
	l := log.New(os.Stdout, "Product-api", 1)

	pp := handler.NewProduct(l)

	sm := http.NewServeMux()
	sm.Handle("/product", pp)

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
	sig := <-c
	log.Println("Got signal:", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
