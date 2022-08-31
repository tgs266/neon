package entities

import (
	"time"

	"github.com/uptrace/bun"
)

type Install struct {
	bun.BaseModel `bun:"table:installs,alias:i"`

	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"createdAt"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"updatedAt"`

	AppName        string `bun:",pk" json:"appName"`
	ProductName    string `bun:",pk" json:"productName"`
	ReleaseVersion string `json:"releaseVersion"`
	Error          string `json:"error"`
}
