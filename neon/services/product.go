package services

import (
	"github.com/gin-gonic/gin"
	"github.com/tgs266/neon/neon/api"
	"github.com/tgs266/neon/neon/errors"
	"github.com/tgs266/neon/neon/store"
	"github.com/tgs266/neon/neon/store/entities"
	"github.com/tgs266/neon/neon/store/repositories"
)

func CreateProduct(request api.CreateProductRequest) {
	internalProduct := entities.Product{
		Name: request.Name,
	}
	if err := store.ProductRepository().Insert(internalProduct); err != nil {
		panic(err)
	}
}

func FindProducts(c *gin.Context, name string, limit, offest int) []entities.Product {
	if res, err := store.ProductRepository().Search(limit, offest, repositories.Query{Query: "name LIKE ?", Arg: "%" + name + "%"}); err != nil || res == nil {
		return []entities.Product{}
	} else {
		return res
	}
}

func GetProductByName(c *gin.Context, name string) entities.Product {
	if res, err := store.ProductRepository().Query(true, true, "name = ?", name); err != nil {
		errors.NewNotFound("product not found", err).Abort(c)
		return entities.Product{}
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
