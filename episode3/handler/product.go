package handler

import (
	"context"
	"episode3/data"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l: l}
}

func (p *Product) GetProducts(rw http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle Get Product")

	lp := data.GetProducts()

	err := lp.ConvertToJson(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Product) AddProduct(rw http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle POSt Product")

	prod := req.Context().Value(keyProduct{}).(data.Product)

	p.l.Println(prod)
	data.AddProducts(&prod)
}

func (p *Product) UpdateProduct(rw http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle PUT Product")
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	prod := req.Context().Value(keyProduct{}).(data.Product)

	p.l.Println(prod)
	err = data.UpdateProduct(id, &prod)

}

type keyProduct struct{}

func (p *Product) MiddlewareProductvalidation(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := data.Product{}
		err := prod.ConvertFromJson(r.Body)
		if err != nil {
			http.Error(w, "Unable to decode the json", http.StatusBadRequest)
			return
		}
		p.l.Println("prod value:", prod)
		ctx := context.WithValue(r.Context(), keyProduct{}, prod)
		req := r.WithContext(ctx)

		h.ServeHTTP(w, req)

	})
}
