package entities

import (
	"time"

	"github.com/uptrace/bun"
)

type App struct {
	bun.BaseModel `bun:"table:apps,alias:a"`

	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"createdAt"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"updatedAt"`

	AppId string `bun:",pk" json:"appId"`
}
