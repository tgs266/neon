package services

import (
	"fmt"
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
		Name:           request.Name,
		Products:       request.Products,
		Repository:     request.Repository,
		Credentials:    request.CredentialName,
		ReleaseChannel: 0,
	}
	if item.Repository != "" {
		git.FillRepository(c, request)
	}
	if err := store.AppRepository().Insert(item); err != nil {
		errors.NewInternal("failed to create app", err).Panic()
		return
	}
	if len(item.Products) != 0 {
		handleAppInstalls(request.Name)
	}
}

func AddProductToApp(c *gin.Context, name string, request api.AddProductRequest) {
	app, err := store.AppRepository().Query(true, "name = ?", name)
	if err != nil {
		errors.NewNotFound("app not found", err).Panic()
	}
	newProducts := append(app.Products, request.Name)
	if err := store.AppRepository().AddProduct(name, newProducts); err != nil {
		errors.NewInternal("failed to update app", err).Panic()
		return
	}
	err = git.AddProduct(c, request.Name, app)
	if err != nil {
		errors.NewInternal("failed to update git", err).Panic()
		return
	}
	handleAppInstalls(request.Name)
}

func handleAppInstalls(appName string) {
	app, _ := store.AppRepository().Query(true, "name = ?", appName)
	fmt.Println(app)
	products, err := store.PullProducts(app.Products, app.ReleaseChannel)
	fmt.Println(products)
	if err != nil {
		store.AppRepository().SetAppError(app.Name, "failed to pull products for app")
		return
	}
	if len(products) == 0 {
		store.AppRepository().SetAppError(app.Name, "no products defined in app manifest found")
		return
	}
	currentInstalls := app.Installs

	installMap := resolveDependencies(products)
	deleteList := []string{}
	updateList := []entities.Release{}
	if installMap == nil {
		store.AppRepository().SetAppError(app.Name, "failed to resolve dependencies")
		return
	}

	fmt.Println(installMap)

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
	fmt.Println(changes)
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
			errors.NewInternal("failed to count apps", err).Panic()
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
		errors.NewNotFound("app not found", err).Panic()
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
	errors.NewNotFound("install for app not found", nil).Panic()
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
