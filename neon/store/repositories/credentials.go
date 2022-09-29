package repositories

import (
	"context"

	"github.com/tgs266/neon/neon/store/entities"
	"github.com/uptrace/bun"
)

type CredentialsRepository struct {
	DB *bun.DB
}

func (r CredentialsRepository) Insert(item entities.Credentials) error {
	_, err := r.DB.NewInsert().
		Model(&item).
		Exec(context.TODO())
	return err
}

func (r CredentialsRepository) Query(query string, args ...interface{}) (entities.Credentials, error) {
	var item entities.Credentials
	call := r.DB.NewSelect().
		Model(&item).
		Where(query, args...)

	err := call.Scan(context.TODO())
	return item, err
}

func (r CredentialsRepository) GetAll() ([]entities.Credentials, error) {
	var item []entities.Credentials
	call := r.DB.NewSelect().
		Model(&item).
		Order("created_at ASC")

	err := call.Scan(context.TODO())
	return item, err
}
