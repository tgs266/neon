package services

import (
	"github.com/gin-gonic/gin"
	"github.com/tgs266/neon/neon/api"
	"github.com/tgs266/neon/neon/errors"
	"github.com/tgs266/neon/neon/store"
	"github.com/tgs266/neon/neon/store/entities"
	"github.com/tgs266/neon/neon/store/repositories"
)

func ListStoredChanges(c *gin.Context, app string, limit, offest int) *api.PaginationResponse[entities.StoredChange] {
	if res, err := store.StoredChangeRepository().Search(limit, offest, repositories.Query{Query: "target_app = ?", Arg: app}); err != nil || res == nil {
		return &api.PaginationResponse[entities.StoredChange]{
			Items: []entities.StoredChange{},
			Total: 0,
		}
	} else {
		if count, err := store.StoredChangeRepository().CountAllForApp(app); err != nil {
			errors.NewInternal("failed to count changes", err).Panic()
			return nil
		} else {
			return &api.PaginationResponse[entities.StoredChange]{
				Items: res,
				Total: count,
			}
		}
	}
}

func ListQueuedChanges(c *gin.Context, app string, limit, offest int) *api.PaginationResponse[entities.QueuedChange] {
	if res, err := store.QueuedChangeRepository().Search(limit, offest, repositories.Query{Query: "target_app = ?", Arg: app}); err != nil || res == nil {
		return &api.PaginationResponse[entities.QueuedChange]{
			Items: []entities.QueuedChange{},
			Total: 0,
		}
	} else {
		if count, err := store.QueuedChangeRepository().CountAllForApp(app); err != nil {
			errors.NewInternal("failed to count changes", err).Panic()
			return nil
		} else {
			return &api.PaginationResponse[entities.QueuedChange]{
				Items: res,
				Total: count,
			}
		}
	}
}
