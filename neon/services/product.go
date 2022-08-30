package services

import (
	"github.com/tgs266/neon/neon/api"
	"github.com/tgs266/neon/neon/store"
	"github.com/tgs266/neon/neon/store/entities"
)

func CreateProduct(request api.CreateProductRequest) {
	internalProduct := entities.Product{
		Name: request.Name,
	}
	if err := store.ProductRepository().Insert(internalProduct); err != nil {
		panic(err)
	}
}

func GetProductByName(name string) entities.Product {
	if res, err := store.ProductRepository().Query(true, "name = ?", name); err != nil {
		panic(err)
	} else {
		return res
	}
}

func ListProducts() []entities.Product {
	products, err := store.List[entities.Product]()
	if err != nil {
		panic(err)
	}
	return products
}
