package repositories

import (
	"context"

	"github.com/tgs266/neon/neon/store/entities"
	"github.com/uptrace/bun"
)

type ReleaseChannelRepository struct {
	DB *bun.DB
}

func (r ReleaseChannelRepository) GetByName(name string) (entities.ReleaseChannel, error) {
	var item entities.ReleaseChannel
	err := r.DB.NewSelect().
		Model(&item).
		Where("name = ?", name).
		Scan(context.TODO())
	return item, err
}

func (r ReleaseChannelRepository) FlipToInt(name string) (int, error) {
	var item entities.ReleaseChannel
	err := r.DB.NewSelect().
		Model(&item).
		Where("name = ?", name).
		Scan(context.TODO())
	if err == nil {
		return item.Value, nil
	}
	return -1, err
}
