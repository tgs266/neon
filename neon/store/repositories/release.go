package repositories

import (
	"context"
	"time"

	"github.com/tgs266/neon/neon/store/entities"
	"github.com/uptrace/bun"
)

type ReleaseRepository struct {
	DB *bun.DB
}

func (r ReleaseRepository) Insert(item entities.Release) error {
	_, err := r.DB.NewInsert().
		Model(&item).
		Exec(context.TODO())
	return err
}

func (r ReleaseRepository) Update(item entities.Release) error {
	_, err := r.DB.NewUpdate().
		Model(&item).
		Set("updated_at = ?", time.Now()).
		OmitZero().
		WherePK().
		Exec(context.TODO())
	return err
}

func (r ReleaseRepository) Query(query string, args ...interface{}) (entities.Release, error) {
	var item entities.Release
	call := r.DB.NewSelect().
		Model(&item).
		Where(query, args...)

	err := call.Scan(context.TODO())
	return item, err
}
