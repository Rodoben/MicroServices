package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"namee"`
	Description string  `json:"desc,omitempty"`
	Price       float32 `json:"price,omitempty"`
	Sku         string  `json:"sku,omitempty"`
	CreatedOn   string  `json:"createdon,omitempty"`
	UpdatedOn   string  `json:"updatedon,omitempty"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func (p *Products) ConvertToJson(i io.Writer) error {

	return json.NewEncoder(i).Encode(p)
}

func GetProducts() Products {
	return productList
}

var productList = Products{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "frothy milk",
		Price:       100.20,
		Sku:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},

	&Product{
		ID:          2,
		Name:        "Esspresso",
		Description: "",
		Price:       100.20,
		Sku:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
