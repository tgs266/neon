package entities

import (
	"time"

	"github.com/uptrace/bun"
)

type Install struct {
	bun.BaseModel `bun:"table:installs,alias:i"`

	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"createdAt"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"updatedAt"`

	AppId          string `json:"appId"`
	InstallId      string `bun:",pk" json:"installId"`
	ProductName    string `json:"productName"`
	ReleaseVersion string `json:"releaseVersion"`
}
