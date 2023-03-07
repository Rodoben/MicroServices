package handler

import (
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}


func NewHello(l *log.Logger) *Hello{
	return &Hello{}
}


func (h *Hello) ServeHTTP(res http.ResponseWriter, req *http.Request){
	log.Println("inside hello handler", req, res)
}