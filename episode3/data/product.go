package data

import (
	"encoding/json"
	"errors"
	"io"
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

func (p *Product) ConvertFromJson(i io.Reader) error {
	p.ID = getNextId()
	return json.NewDecoder(i).Decode(p)
}

func getNextId() int {
	lp := len(productList) - 1
	return lp + 1
}

type Products []*Product

func (p *Products) ConvertToJson(i io.Writer) error {

	return json.NewEncoder(i).Encode(p)
}

func GetProducts() Products {
	return productList
}

func AddProducts(p *Product) {
	p.ID = getNextId()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[pos] = p
	return nil

}

func findProduct(id int) (*Product, int, error) {
	for i, v := range productList {
		if v.ID == id {
			return v, i, nil
		}
	}
	return nil, -1, errors.New("Product not found")
}

var productList = Products{}

// &Product{
// 	ID:          1,
// 	Name:        "Latte",
// 	Description: "frothy milk",
// 	Price:       100.20,
// 	Sku:         "abc123",
// 	CreatedOn:   time.Now().UTC().String(),
// 	UpdatedOn:   time.Now().UTC().String(),
// },

// &Product{
// 	ID:          2,
// 	Name:        "Esspresso",
// 	Description: "",
// 	Price:       100.20,
// 	Sku:         "abc123",
// 	CreatedOn:   time.Now().UTC().String(),
// 	UpdatedOn:   time.Now().UTC().String(),
// },
//}
