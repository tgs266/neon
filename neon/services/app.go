package services

import (
	"github.com/rs/zerolog/log"
	"github.com/tgs266/neon/neon/api"
	"github.com/tgs266/neon/neon/store"
	"github.com/tgs266/neon/neon/store/entities"
)

func ApplyApp(request api.ApplyAppRequest) {
	count, _ := store.Count[entities.App]("name = ?", request.Name)
	if count == 0 {
		createApp(request)
	} else {
		updateApp(request)
	}
}

func createApp(request api.ApplyAppRequest) {
	item := entities.App{
		Name:     request.Name,
		Products: request.Products,
	}
	if err := store.AppRepository().Insert(item); err != nil {
		panic(err)
	}
	handleAppInstalls(request.Name, false)
}

func updateApp(request api.ApplyAppRequest) {
	item := entities.App{
		Name:     request.Name,
		Products: request.Products,
	}
	if err := store.AppRepository().Update(item); err != nil {
		panic(err)
	}
	handleAppInstalls(request.Name, true)
}

func handleAppInstalls(appName string, update bool) {
	app, _ := store.AppRepository().Query(true, "name = ?", appName)

	products, err := store.PullProducts(app.Products, app.ReleaseChannel)
	if err != nil {
		panic(err)
	}
	if len(products) == 0 {
		panic("could not install requested products: no products found")
	}
	installs := []entities.Install{}
	currentInstalls := app.Installs

	// actually generate install here
	installMap := resolveDependencies(products)
	if installMap == nil {
		panic("Failed to resolve dependencies")
	}

	for _, install := range currentInstalls {
		if expectedInstall, exists := installMap[install.ProductName]; exists {
			if expectedInstall.ProductVersion == install.ReleaseVersion {
				delete(installMap, install.ProductName)
			}
		}
	}

	for k, v := range installMap {
		if stderr, err := installUpdateHelmChart(k, v); err != nil {
			log.Error().Err(err).Msg(stderr)
			continue
		}
		installs = append(installs, entities.Install{
			AppName:        app.Name,
			ProductName:    k,
			ReleaseVersion: v.ProductVersion,
		})
	}

	if err := store.InstallRepository().InsertBatch(installs); err != nil {
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

func GetAppByName(name string) entities.App {
	if res, err := store.AppRepository().Query(true, "name = ?", name); err != nil {
		panic(err)
	} else {
		return res
	}
}
