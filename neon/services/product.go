package services

import (
	"github.com/google/uuid"
	"github.com/tgs266/neon/neon/api"
	"github.com/tgs266/neon/neon/store"
	"github.com/tgs266/neon/neon/store/entities"
)

func CreateProduct(request api.CreateProductRequest) {
	internalProduct := entities.Product{
		ID:   uuid.New().String(),
		Name: request.Name,
	}
	if err := store.InsertProduct(internalProduct); err != nil {
		panic(err)
	}
}
