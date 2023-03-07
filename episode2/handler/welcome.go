package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Welcome struct {
	l *log.Logger
}

func NewWelcome(l *log.Logger) *Welcome {
	return &Welcome{l}
}

func (g *Welcome) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	log.Println("log in Welcome")

	b, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "Error reading the body", http.StatusBadRequest)
	}

	fmt.Fprintf(res, "Welcome: %s", b)
}
