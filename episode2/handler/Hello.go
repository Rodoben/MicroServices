package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	fmt.Println("creating new hello obj")
	return &Hello{l}
}

func (h *Hello) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	h.l.Println("Handle Hello requests")
	log.Println("inside hello handler")

	b, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "Error reading the body", http.StatusBadRequest)
	}

	fmt.Fprintf(res, "Hello %s", b)

}
