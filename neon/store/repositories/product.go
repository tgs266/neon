package repositories

import (
	"context"
	"time"

	"github.com/tgs266/neon/neon/store/entities"
	"github.com/uptrace/bun"
)

type ProductRepository struct {
	DB *bun.DB
}

func (r ProductRepository) Insert(item entities.Product) error {
	_, err := r.DB.NewInsert().
		Model(&item).
		Exec(context.TODO())
	return err
}

func (r ProductRepository) Update(item entities.Product) error {
	_, err := r.DB.NewUpdate().
		Model(&item).
		Set("updated_at = ?", time.Now()).
		OmitZero().
		WherePK().
		Exec(context.TODO())
	return err
}

func (r ProductRepository) Query(includeReleases bool, includeInstalls bool, query string, args ...interface{}) (entities.Product, error) {
	var item entities.Product
	call := r.DB.NewSelect().
		Model(&item).
		Where(query, args...)

	if includeReleases {
		call = call.Relation("Releases")
	}

	if includeInstalls {
		call = call.Relation("Installs")
	}

	err := call.Scan(context.TODO())
	return item, err
}

func (r ProductRepository) Search(limit, offset int, queries ...Query) ([]entities.Product, error) {
	var item []entities.Product
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
