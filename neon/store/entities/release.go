package entities

import (
	"time"

	"github.com/uptrace/bun"
)

type Release struct {
	bun.BaseModel `bun:"table:releases,alias:r"`

	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"createdAt"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"updatedAt"`

	ProductName    string `bun:",pk" json:"productName"`
	ProductVersion string `bun:",pk" json:"productVersion"`
	ReleaseChannel int    `json:"releaseChannel"`

	Recalled bool `json:"recalled"`

	Dependencies []Dependency `json:"dependencies"`
}

type Dependency struct {
	ProductName string `json:"productName"`
	MinVersion  string `json:"minVersion"`
	MaxVersion  string `json:"maxVersion"`
}
