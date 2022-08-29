package services

import (
	"github.com/google/uuid"
	"github.com/tgs266/neon/neon/api"
	"github.com/tgs266/neon/neon/store"
	"github.com/tgs266/neon/neon/store/entities"
)

func CreateRelease(request api.CreateReleaseRequest) {
	item := entities.Release{
		ReleaseId:      uuid.New().String(),
		ProductName:    request.ProductName,
		ProductVersion: request.ProductVersion,
		ReleaseChannel: request.ReleaseChannel,
		Dependencies:   request.Dependencies,
	}
	if err := store.InsertRelease(item); err != nil {
		panic(err)
	}
}
