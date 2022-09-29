package services

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tgs266/neon/neon/api"
	"github.com/tgs266/neon/neon/errors"
	"github.com/tgs266/neon/neon/git"
	"github.com/tgs266/neon/neon/kubernetes"
	"github.com/tgs266/neon/neon/store"
	"github.com/tgs266/neon/neon/store/entities"
	"github.com/tgs266/neon/neon/store/repositories"
)

func CreateApp(c *gin.Context, request api.CreateAppRequest) {
	item := entities.App{
		Name:        request.Name,
		Products:    request.Products,
		Repository:  request.Repository,
		Credentials: request.CredentialName,
	}
	if err := store.AppRepository().Insert(item); err != nil {
		errors.NewInternal("failed to create app", err).Abort(c)
		return
	}
	if item.Repository != "" {
		git.FillRepository(c, request)
	}
	handleAppInstalls(request.Name, true)
}

func UpdateApp(c *gin.Context, request api.ApplyAppRequest) {
	item := entities.App{
		Name:     request.Name,
		Products: request.Products,
	}
	if err := store.AppRepository().Update(item); err != nil {
		errors.NewInternal("failed to update app", err).Abort(c)
		return
	}
	handleAppInstalls(request.Name, true)
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
	currentInstalls := app.Installs

	// store.AppRepository().SetAppField(app.Name, "install_status", "IN_PROGRESS")

	// actually generate install here
	installMap := resolveDependencies(products)
	deleteList := []string{}
	updateList := []entities.Release{}
	if installMap == nil {
		store.AppRepository().SetAppError(app.Name, "failed to resolve dependencies")
		return
	}

	for _, install := range currentInstalls {
		if expectedInstall, exists := installMap[install.ProductName]; exists {
			if expectedInstall.ProductVersion == install.ReleaseVersion {
				delete(installMap, install.ProductName)
			} else {
				updateList = append(updateList, *expectedInstall)
			}
		} else {
			deleteList = append(deleteList, install.ProductName)
		}
	}

	changes := []entities.QueuedChange{}

	for _, i := range deleteList {
		changes = append(changes, entities.QueuedChange{
			Type:        "DELETE",
			TargetApp:   app.Name,
			Release:     entities.Release{ProductName: i},
			LastChecked: time.Now(),
			ID:          uuid.New().String(),
		})
	}

	for _, i := range updateList {
		changes = append(changes, entities.QueuedChange{
			Type:        "UPDATE",
			TargetApp:   app.Name,
			Release:     i,
			LastChecked: time.Now(),
			ID:          uuid.New().String(),
		})
	}

	for _, v := range installMap {
		changes = append(changes, entities.QueuedChange{
			Type:        "INSTALL",
			TargetApp:   app.Name,
			Release:     *v,
			LastChecked: time.Now(),
			ID:          uuid.New().String(),
		})
	}
	store.QueuedChangeRepository().InsertBatch(changes)
}

func ListApps(c *gin.Context, name string, limit, offest int) *api.PaginationResponse[entities.App] {
	if res, err := store.AppRepository().Search(limit, offest, repositories.Query{Query: "name LIKE ?", Arg: "%" + name + "%"}); err != nil || res == nil {
		return &api.PaginationResponse[entities.App]{
			Items: []entities.App{},
			Total: 0,
		}
	} else {
		if count, err := store.AppRepository().CountAll(); err != nil {
			errors.NewInternal("failed to count apps", err).Abort(c)
			return nil
		} else {
			return &api.PaginationResponse[entities.App]{
				Items: res,
				Total: count,
			}
		}
	}
}

func GetAppByName(c *gin.Context, name string) entities.App {
	if res, err := store.AppRepository().Query(true, "name = ?", name); err != nil {
		errors.NewNotFound("app not found", err).Abort(c)
		return entities.App{}
	} else {
		return res
	}
}

func GetAppInstall(c *gin.Context, name, productName string) *entities.Install {
	app := GetAppByName(c, name)
	for _, i := range app.Installs {
		if i.ProductName == productName {
			return i
		}
	}
	errors.NewNotFound("install for app not found", nil).Abort(c)
	return nil
}

func GetAppInstallResources(c *gin.Context, name, productName string) api.ResourceList {
	GetAppInstall(c, name, productName)
	resPods := kubernetes.Pods(c, name).ListByInstanceLabel(productName)
	resSvcs := kubernetes.Services(c, name).ListByInstanceLabel(productName)

	podNames := []string{}
	svcNames := []string{}

	for _, i := range resPods {
		podNames = append(podNames, i.Name)
	}

	for _, i := range resSvcs {
		svcNames = append(svcNames, i.Name)
	}

	return api.ResourceList{
		Pods:     podNames,
		Services: svcNames,
	}
}
