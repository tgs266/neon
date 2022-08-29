package entities

import (
	"time"

	"github.com/uptrace/bun"
)

type Product struct {
	bun.BaseModel `bun:"table:products,alias:p"`

	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"createdAt"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"updatedAt"`

	ID   string `bun:",pk" json:"id"`
	Name string `bun:",unique" json:"name"`

	Releases []*Release `bun:"rel:has-many,join:name=product_name" json:"releases,omitempty"`
	Installs []*Install `bun:"rel:has-many,join:name=product_name" json:"installs,omitempty"`
}
