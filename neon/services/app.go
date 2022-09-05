package services

import (
	"github.com/gin-gonic/gin"
	"github.com/tgs266/neon/neon/api"
	"github.com/tgs266/neon/neon/errors"
	"github.com/tgs266/neon/neon/store"
	"github.com/tgs266/neon/neon/store/entities"
	"github.com/tgs266/neon/neon/store/repositories"
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

	store.AppRepository().SetAppField(app.Name, "install_status", "IN_PROGRESS")

	// actually generate install here
	installMap := resolveDependencies(products)
	deleteList := []string{}
	if installMap == nil {
		store.AppRepository().SetAppError(app.Name, "failed to resolve dependencies")
		return
	}

	for _, install := range currentInstalls {
		if expectedInstall, exists := installMap[install.ProductName]; exists {
			if expectedInstall.ProductVersion == install.ReleaseVersion {
				delete(installMap, install.ProductName)
			}
		} else {
			deleteList = append(deleteList, install.ProductName)
		}
	}

	hasInstallError := false

	for _, i := range deleteList {
		deleteHelmRelease(app.Name, i)
		store.InstallRepository().DeleteByPk(app.Name, i)
	}

	for k, v := range installMap {
		stderr, err := installUpdateHelmChart(app.Name, k, v)
		if err != nil {
			hasInstallError = true
		}
		installs = append(installs, entities.Install{
			AppName:        app.Name,
			ProductName:    k,
			ReleaseVersion: v.ProductVersion,
			Error:          stderr,
		})
	}

	if hasInstallError {
		store.AppRepository().SetAppField(app.Name, "install_status", "FAILED")
	} else {
		store.AppRepository().SetAppField(app.Name, "install_status", "COMPLETE")
	}

	if len(installs) != 0 {
		if err := store.InstallRepository().InsertBatch(installs); err != nil {
			store.AppRepository().SetAppError(app.Name, "failed to store installs in database")
			return
		}
	}
	store.AppRepository().SetAppError(app.Name, "")
}

func ListApps(c *gin.Context, name string, limit, offest int) *api.PaginationResponse[entities.App] {
	if res, err := store.AppRepository().Search(limit, offest, repositories.Query{Query: "name LIKE ?", Arg: "%" + name + "%"}); err != nil || res == nil {
		return &api.PaginationResponse[entities.App]{
			Items: []entities.App{},
			Total: 0,
		}
	} else {
		if count, err := store.ProductRepository().CountAll(); err != nil {
			errors.NewInternal("failed to count products", err).Abort(c)
			return nil
		} else {
			return &api.PaginationResponse[entities.App]{
				Items: res,
				Total: count,
			}
		}
	}
}

func GetAppByName(name string) entities.App {
	if res, err := store.AppRepository().Query(true, "name = ?", name); err != nil {
		panic(err)
	} else {
		return res
	}
}
