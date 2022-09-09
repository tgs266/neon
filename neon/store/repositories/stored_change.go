package repositories

import (
	"context"
	"time"

	"github.com/tgs266/neon/neon/store/entities"
	"github.com/uptrace/bun"
)

type StoredChangeRepository struct {
	DB *bun.DB
}

func (a StoredChangeRepository) CountAllForApp(app string) (int, error) {
	var item entities.StoredChange
	count, err := a.DB.NewSelect().
		Model(&item).
		Where("target_app = ?", app).
		Count(context.TODO())
	return count, err
}

func (a StoredChangeRepository) Insert(item entities.StoredChange) error {
	_, err := a.DB.NewInsert().
		Model(&item).
		Exec(context.TODO())
	return err
}

func (a StoredChangeRepository) InsertBatch(item []entities.StoredChange) error {
	_, err := a.DB.NewInsert().
		Model(&item).
		Exec(context.TODO())
	return err
}

func (a StoredChangeRepository) Update(item entities.StoredChange) error {
	call := a.DB.NewUpdate().
		Model(&item).
		Set("updated_at = ?", time.Now()).
		OmitZero()

	_, err := call.
		WherePK().
		Exec(context.TODO())
	return err
}

func (a StoredChangeRepository) Delete(id string) error {
	var item entities.StoredChange
	_, err := a.DB.NewDelete().
		Model(&item).
		Where("id = ?", id).
		Exec(context.TODO())
	return err
}

func (a StoredChangeRepository) Grab() (entities.StoredChange, error) {
	var item entities.StoredChange
	call := a.DB.NewSelect().
		Model(&item)
	err := call.Scan(context.TODO())
	return item, err
}

func (r StoredChangeRepository) Search(limit, offset int, queries ...Query) ([]entities.StoredChange, error) {
	var item []entities.StoredChange
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
