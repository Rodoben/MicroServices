package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type GoodBye struct {
	l *log.Logger
}

func NewGoodBye(l *log.Logger) *GoodBye {
	return &GoodBye{l}
}

func (g *GoodBye) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	log.Println("log in goodbye")

	b, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "Error reading the body", http.StatusBadRequest)
	}

	fmt.Fprintf(res, "GoodBye %s", b)
}
