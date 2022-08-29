package entities

import (
	"time"

	"github.com/uptrace/bun"
)

type Release struct {
	bun.BaseModel `bun:"table:releases,alias:r"`

	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"createdAt"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"updatedAt"`

	ReleaseId string `bun:",pk" json:"releaseId"`

	ProductName    string `bun:",unique" json:"productName"`
	ProductVersion string `bun:",unique" json:"productVersion"`
	ReleaseChannel string `json:"releaseChannel"`

	Recalled bool `json:"recalled"`

	Dependencies []Dependency `json:"dependencies"`
}

type Dependency struct {
	ProductName string `json:"productName"`
	MinVersion  string `json:"minVersion"`
	MaxVersion  string `json:"maxVersion"`
}
