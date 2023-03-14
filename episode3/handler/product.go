package handler

import (
	"episode3/data"
	"log"
	"net/http"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l: l}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodGet {
		p.GetProducts(rw, req)
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Product) GetProducts(rw http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle Get Product")

	lp := data.GetProducts()

	err := lp.ConvertToJson(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
