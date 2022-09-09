package repositories

import (
	"context"

	"github.com/tgs266/neon/neon/store/entities"
	"github.com/uptrace/bun"
)

type QueuedChangeRepository struct {
	DB *bun.DB
}

func (a QueuedChangeRepository) CountAllForApp(app string) (int, error) {
	var item entities.QueuedChange
	count, err := a.DB.NewSelect().
		Model(&item).
		Where("target_app = ?", app).
		Count(context.TODO())
	return count, err
}

func (a QueuedChangeRepository) Insert(item entities.QueuedChange) error {
	_, err := a.DB.NewInsert().
		Model(&item).
		Exec(context.TODO())
	return err
}

func (a QueuedChangeRepository) InsertBatch(item []entities.QueuedChange) error {
	_, err := a.DB.NewInsert().
		Model(&item).
		Exec(context.TODO())
	return err
}

func (a QueuedChangeRepository) Update(item entities.QueuedChange) error {
	call := a.DB.NewUpdate().
		Model(&item).
		OmitZero()

	_, err := call.
		WherePK().
		Exec(context.TODO())
	return err
}

func (a QueuedChangeRepository) Delete(id string) error {
	var item entities.QueuedChange
	_, err := a.DB.NewDelete().
		Model(&item).
		Where("id = ?", id).
		Exec(context.TODO())
	return err
}

func (a QueuedChangeRepository) Grab() (entities.QueuedChange, error) {
	var item entities.QueuedChange
	call := a.DB.NewSelect().
		Model(&item).
		Order("last_checked ASC")
	err := call.Scan(context.TODO())
	return item, err
}

func (r QueuedChangeRepository) Search(limit, offset int, queries ...Query) ([]entities.QueuedChange, error) {
	var item []entities.QueuedChange
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
