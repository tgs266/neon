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
