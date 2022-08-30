package entities

import (
	"time"

	"github.com/uptrace/bun"
)

type ReleaseChannel struct {
	bun.BaseModel `bun:"table:release_channels,alias:rcs"`

	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"createdAt"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"updatedAt"`

	Name  string `bun:",pk" json:"name"`
	Value int    `json:"value"`
}
