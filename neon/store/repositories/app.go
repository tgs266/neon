package repositories

import (
	"context"
	"time"

	"github.com/tgs266/neon/neon/store/entities"
	"github.com/uptrace/bun"
)

type AppRepository struct {
	DB *bun.DB
}

func (a AppRepository) Insert(item entities.App) error {
	_, err := a.DB.NewInsert().
		Model(&item).
		Exec(context.TODO())
	return err
}

func (a AppRepository) Update(item entities.App) error {
	call := a.DB.NewUpdate().
		Model(&item).
		Set("updated_at = ?", time.Now()).
		OmitZero()

	_, err := call.
		WherePK().
		Exec(context.TODO())
	return err
}

func (a AppRepository) Query(includeInstalls bool, query string, args ...interface{}) (entities.App, error) {
	var item entities.App
	call := a.DB.NewSelect().
		Model(&item).
		Where(query, args...)

	if includeInstalls {
		call = call.Relation("Installs")
	}

	err := call.Scan(context.TODO())
	return item, err
}

func (a AppRepository) SetAppError(appName string, errString string) error {
	var item entities.App
	_, err := a.DB.NewUpdate().
		Model(&item).
		Where("name = ?", appName).
		Set("error = ?", errString).
		Exec(context.TODO())
	return err
}

func (a AppRepository) SetAppField(appName string, fieldName, str string) error {
	var item entities.App
	_, err := a.DB.NewUpdate().
		Model(&item).
		Where("name = ?", appName).
		Set(fieldName+" = ?", str).
		Exec(context.TODO())
	return err
}

func (r AppRepository) Search(limit, offset int, queries ...Query) ([]entities.App, error) {
	var item []entities.App
	call := r.DB.NewSelect().
		Model(&item)

	for _, q := range queries {
		call = q.Apply(call)
	}

	if offset >= 0 {
		call = call.Offset(offset)
	}

	if limit >= 0 {
		call = call.Limit(limit)
	}

	err := call.Scan(context.TODO())
	return item, err
}
