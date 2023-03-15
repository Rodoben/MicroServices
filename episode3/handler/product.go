package handler

import (
	"episode3/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
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
		return
	}

	if req.Method == http.MethodPost {
		p.AddProduct(rw, req)
		return
	}

	if req.Method == http.MethodPut {
		p.l.Println("PUT:", req.URL.Path)

		reg := regexp.MustCompile(`/([0-9]+)`)
		p.l.Println(reg)
		g := reg.FindAllStringSubmatch(req.URL.Path, -1)
		p.l.Println(g)
		p.l.Println(len(g))

		if len(g) != 1 {
			p.l.Println("Invalid URI more than one id")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		p.l.Println(len(g[0]))
		if len(g[0]) != 2 {
			p.l.Println("Invalid URI more than one capture group")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		p.l.Println("id:", idString)
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.l.Println("Invalid URI unable to convert to numer", idString)
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		p.UpdateProduct(id, rw, req)
		return
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

func (p *Product) AddProduct(rw http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle POSt Product")

	prod := data.Product{}
	err := prod.ConvertFromJson(req.Body)
	if err != nil {
		http.Error(rw, "Unable to decode the json", http.StatusBadRequest)
	}

	p.l.Println(prod)
	data.AddProducts(&prod)
}

func (p *Product) UpdateProduct(id int, rw http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle PUT Product")

	prod := data.Product{}
	err := prod.ConvertFromJson(req.Body)
	if err != nil {
		http.Error(rw, "Unable to decode the json", http.StatusBadRequest)
	}
	p.l.Println(prod)
	data.UpdateProduct(id, &prod)

}
