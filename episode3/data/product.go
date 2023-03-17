package data

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"regexp"

	"github.com/go-playground/validator"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"namee" validate:"required"`
	Description string  `json:"desc,omitempty"`
	Price       float32 `json:"price,omitempty" validate:"gt=0"`
	Sku         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"createdon,omitempty"`
	UpdatedOn   string  `json:"updatedon,omitempty"`
	DeletedOn   string  `json:"-"`
}

func (p *Product) ConvertFromJson(i io.Reader) error {
	p.ID = getNextId()
	return json.NewDecoder(i).Decode(p)
}

func (p *Product) Validate() error {
	log.Println("entered validation of sku ")
	validate := validator.New()
	validate.RegisterValidation("sku", validateSku)
	return validate.Struct(p)
}

func validateSku(f validator.FieldLevel) bool {

	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(f.Field().String(), -1)
	if len(matches) != 1 {
		return false
	}
	return true
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
