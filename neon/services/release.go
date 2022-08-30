package services

import (
	"fmt"

	"github.com/tgs266/neon/neon/api"
	"github.com/tgs266/neon/neon/store"
	"github.com/tgs266/neon/neon/store/entities"
)

func CreateRelease(request api.CreateReleaseRequest) {
	intRc, err := store.ReleaseChannelRepository().FlipToInt(request.ReleaseChannel)
	if err != nil || intRc == -1 {
		fmt.Println(intRc)
		panic(err)
	}
	item := entities.Release{
		// ReleaseId:      uuid.New().String(),
		ProductName:    request.ProductName,
		ProductVersion: request.ProductVersion,
		ReleaseChannel: intRc,
		Dependencies:   request.Dependencies,
		HelmChart:      request.HelmChart,
	}
	if err := store.ReleaseRepository().Insert(item); err != nil {
		panic(err)
	}
}
