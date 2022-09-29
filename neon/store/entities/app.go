package entities

import (
	"time"

	"github.com/uptrace/bun"
)

type App struct {
	bun.BaseModel `bun:"table:apps,alias:a"`

	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"createdAt"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"updatedAt"`

	// ID       string     `bun:",pk" json:"appId"`
	Name     string     `bun:",pk" json:"name"`
	Products []string   `bun:",array" json:"products"`
	Installs []*Install `bun:"rel:has-many,join:name=app_name" json:"installs,omitempty"`

	ReleaseChannel int `json:"releaseChannel"`

	Error         string `json:"error"`
	InstallStatus string `json:"installStatus"`

	Repository  string `json:"repository"`
	Credentials string `json:"credentials"`
}
