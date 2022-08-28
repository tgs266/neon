package entities

import "github.com/uptrace/bun"

type Product struct {
	bun.BaseModel `bun:"table:products,alias:p"`

	ID   string `bun:",pk"`
	Name string `bun:",unique"`
}
