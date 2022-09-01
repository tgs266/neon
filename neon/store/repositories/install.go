package repositories

import (
	"context"
	"time"

	"github.com/tgs266/neon/neon/store/entities"
	"github.com/uptrace/bun"
)

type InstallRepository struct {
	DB *bun.DB
}

func (a InstallRepository) Insert(item entities.Install) error {
	_, err := a.DB.NewInsert().
		Model(&item).
		Exec(context.TODO())
	return err
}

func (a InstallRepository) InsertBatch(items []entities.Install) error {
	_, err := a.DB.NewInsert().
		Model(&items).
		Exec(context.TODO())
	return err
}

func (a InstallRepository) Update(item entities.Install) error {
	_, err := a.DB.NewUpdate().
		Model(&item).
		Set("updated_at = ?", time.Now()).
		OmitZero().
		WherePK().
		Exec(context.TODO())
	return err
}

func (a InstallRepository) UpdateBatch(item []entities.Install) error {
	_, err := a.DB.NewInsert().
		Model(&item).
		Set("updated_at = ?", time.Now()).
		Ignore().
		Exec(context.TODO())
	return err
}

func (a InstallRepository) DeleteByPk(appName string, ProductName string) error {
	var item entities.Install
	_, err := a.DB.NewDelete().
		Model(&item).
		Where("app_name = ?", appName).
		Where("product_name = ?", ProductName).
		Exec(context.TODO())
	return err
}
