package services

import (
	"github.com/gin-gonic/gin"
	"github.com/tgs266/neon/neon/api"
	"github.com/tgs266/neon/neon/errors"
	"github.com/tgs266/neon/neon/store"
	"github.com/tgs266/neon/neon/store/entities"
)

func ApplyApp(c *gin.Context, request api.ApplyAppRequest) {
	count, _ := store.Count[entities.App]("name = ?", request.Name)
	if count == 0 {
		createApp(c, request)
	} else {
		updateApp(c, request)
	}
}

func createApp(c *gin.Context, request api.ApplyAppRequest) {
	item := entities.App{
		Name:     request.Name,
		Products: request.Products,
	}
	if err := store.AppRepository().Insert(item); err != nil {
		errors.NewInternal("failed to create app", err).Abort(c)
		return
	}
	go handleAppInstalls(request.Name, true)
}

func updateApp(c *gin.Context, request api.ApplyAppRequest) {
	item := entities.App{
		Name:     request.Name,
		Products: request.Products,
	}
	if err := store.AppRepository().Update(item); err != nil {
		errors.NewInternal("failed to update app", err).Abort(c)
		return
	}
	go handleAppInstalls(request.Name, true)
}

func handleAppInstalls(appName string, update bool) {
	app, _ := store.AppRepository().Query(true, "name = ?", appName)

	products, err := store.PullProducts(app.Products, app.ReleaseChannel)
	if err != nil {
		store.AppRepository().SetAppError(app.Name, "failed to pull products for app")
		return
	}
	if len(products) == 0 {
		store.AppRepository().SetAppError(app.Name, "no products defined in app manifest found")
		return
	}
	installs := []entities.Install{}
	currentInstalls := app.Installs

	// actually generate install here
	installMap := resolveDependencies(products)
	if installMap == nil {
		store.AppRepository().SetAppError(app.Name, "failed to resolve dependencies")
		return
	}

	for _, install := range currentInstalls {
		if expectedInstall, exists := installMap[install.ProductName]; exists {
			if expectedInstall.ProductVersion == install.ReleaseVersion {
				delete(installMap, install.ProductName)
			}
		}
	}

	for k, v := range installMap {
		stderr, _ := installUpdateHelmChart(app.Name, k, v)
		installs = append(installs, entities.Install{
			AppName:        app.Name,
			ProductName:    k,
			ReleaseVersion: v.ProductVersion,
			Error:          stderr,
		})
	}

	if err := store.InstallRepository().InsertBatch(installs); err != nil {
		store.AppRepository().SetAppError(app.Name, "failed to store installs in database")
		return
	}
	store.AppRepository().SetAppError(app.Name, "")
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
