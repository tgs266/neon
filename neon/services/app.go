package services

import (
	"github.com/google/uuid"
	"github.com/tgs266/neon/neon/api"
	"github.com/tgs266/neon/neon/store"
	"github.com/tgs266/neon/neon/store/entities"
)

func CreateApp(request api.CreateAppRequest) {
	item := entities.App{
		AppId: uuid.New().String(),
		Name:  request.Name,
	}
	if err := store.Insert(item); err != nil {
		panic(err)
	}
}

func ListApps() []entities.App {
	products, err := store.List[entities.App]()
	if err != nil {
		panic(err)
	}
	return products
}
