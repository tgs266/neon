package services

import (
	"time"

	"github.com/gammazero/workerpool"
	"github.com/tgs266/neon/neon/store"
	"github.com/tgs266/neon/neon/store/entities"
)

var wp *workerpool.WorkerPool

func InitPool(size int) {
	wp = workerpool.New(size)
	go StartWorkers()
}

func GrabQueuedChange() (entities.QueuedChange, error) {
	if change, err := store.QueuedChangeRepository().Grab(); err != nil {
		return entities.QueuedChange{}, nil
	} else {
		return change, err
	}
}

func RunQueuedChange(change *entities.QueuedChange) (*entities.QueuedChange, error) {

	appName := change.TargetApp

	app, err := store.AppRepository().Query(true, "name = ?", appName)
	if err != nil {
		return change, err
	}

	if change.Type == "delete" {
		return RunDeleteChange(change, app)
	} else if change.Type == "update" {
		return RunUpdateChange(change, app)
	} else {
		return RunInstallChange(change, app)
	}
}

func RunDeleteChange(change *entities.QueuedChange, app entities.App) (*entities.QueuedChange, error) {
	release := change.Release
	if stderr, err := deleteHelmRelease(app.Name, release.ProductName); err == nil {
		err = store.InstallRepository().DeleteByPk(app.Name, release.ProductName)
		return change, err
	} else {
		change.Details = stderr
		return change, err
	}
}

func RunUpdateChange(change *entities.QueuedChange, app entities.App) (*entities.QueuedChange, error) {
	release := change.Release

	stderr, err := installUpdateHelmChart(app.Name, release.ProductName, &release)
	if err != nil {
		change.Details = stderr
		return change, err
	}

	install := entities.Install{
		AppName:        app.Name,
		ProductName:    release.ProductName,
		ReleaseVersion: release.ProductVersion,
	}

	err = store.InstallRepository().Update(install)
	return change, err
}

func RunInstallChange(change *entities.QueuedChange, app entities.App) (*entities.QueuedChange, error) {
	release := change.Release

	stderr, err := installUpdateHelmChart(app.Name, release.ProductName, &release)
	if err != nil {
		change.Details = stderr
		return change, err
	}

	install := entities.Install{
		AppName:        app.Name,
		ProductName:    release.ProductName,
		ReleaseVersion: release.ProductVersion,
	}

	err = store.InstallRepository().Insert(install)
	return change, err
}

func RunJob() {
	qc, err := GrabQueuedChange()
	if err != nil {
		return
	}
	newQc, err := RunQueuedChange(&qc)
	newQc.LastChecked = time.Now()
	if err != nil {
		store.QueuedChangeRepository().Update(*newQc)
	} else {
		store.QueuedChangeRepository().Delete(newQc.ID)
		store.StoredChangeRepository().Insert(newQc.ToSC())
	}
}

func StartWorkers() {
	for {
		wp.Submit(RunJob)
		time.Sleep(5 * time.Second)
	}
}
